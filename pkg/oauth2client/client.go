package oauth2client

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/json"
	"encoding/pem"
	"io/ioutil"
	"strings"

	"github.com/go-phorce/dolly/fileutil"
	"github.com/go-phorce/dolly/xlog"
	"github.com/juju/errors"
	yaml "gopkg.in/yaml.v2"
)

var logger = xlog.NewPackageLogger("github.com/ekspand/trusty/pkg", "oauth2client")

// Client of OAuth2
type Client struct {
	cfg       *Config
	verifyKey *rsa.PublicKey
}

// LoadConfig returns configuration loaded from a file
func LoadConfig(file string) (*Config, error) {
	if file == "" {
		return &Config{}, nil
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.Trace(err)
	}

	var config Config
	if strings.HasSuffix(file, ".json") {
		err = json.Unmarshal(b, &config)
	} else {
		err = yaml.Unmarshal(b, &config)
	}
	if err != nil {
		return nil, errors.Annotatef(err, "unable to unmarshal %q", file)
	}

	return &config, nil
}

// Load returns new Provider
func Load(cfgfile string) (*Client, error) {
	logger.Infof("file=%q", cfgfile)

	cfg, err := LoadConfig(cfgfile)
	if err != nil {
		return nil, errors.Trace(err)
	}

	cfg.ClientID, err = fileutil.LoadConfigWithSchema(cfg.ClientID)
	if err != nil {
		return nil, errors.Trace(err)
	}
	cfg.ClientSecret, err = fileutil.LoadConfigWithSchema(cfg.ClientSecret)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return New(cfg)
}

// New returns new Provider
func New(cfg *Config) (*Client, error) {
	p := &Client{
		cfg: cfg,
	}

	if cfg.PubKey != "" {
		key := strings.TrimSpace(cfg.PubKey)
		verifyKey, err := ParseRSAPublicKeyFromPEM([]byte(key))
		if err != nil {
			return nil, errors.Annotatef(err, "unable to parse Public Key: %q", key)
		}
		p.verifyKey = verifyKey
	}

	logger.Infof("provider=%q, audience=%q, issuer=%q", cfg.ProviderID, cfg.Audience, cfg.Issuer)

	return p, nil
}

// Config returns OAuth2 configuration
func (p *Client) Config() *Config {
	return p.cfg
}

// SetPubKey replaces the OAuth public signing key loaded from configuration
// During normal operation, identity provider's public key is read from config on start-up.
func (p *Client) SetPubKey(newPubKey *rsa.PublicKey) {
	p.verifyKey = newPubKey
}

// SetClientSecret sets Client Secret
func (p *Client) SetClientSecret(s string) *Client {
	p.cfg.ClientSecret = s
	return p
}

// ParseRSAPublicKeyFromPEM parses PEM encoded RSA public key
// TODO: move to dolly
func ParseRSAPublicKeyFromPEM(key []byte) (*rsa.PublicKey, error) {
	var err error

	// Parse PEM block
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("key must be PEM encoded")
	}

	// Parse the key
	parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		if _, err = asn1.Unmarshal(block.Bytes, &parsedKey); err != nil {
			return nil, errors.New("unable to parse RSA Public Key")
		}
	}

	pkey, ok := parsedKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not RSA Public Key")
	}

	return pkey, nil
}
