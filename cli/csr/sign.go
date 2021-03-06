package csr

import (
	"strings"

	"github.com/ekspand/trusty/authority"
	"github.com/ekspand/trusty/cli"
	"github.com/ekspand/trusty/pkg/csr"
	"github.com/ekspand/trusty/pkg/print"
	"github.com/go-phorce/dolly/ctl"
	"github.com/juju/errors"
)

// SignFlags specifies flags for Sign command
type SignFlags struct {
	// CACert specifies file name with CA cert
	CACert *string
	// CAKey specifies file name with CA key
	CAKey *string
	// CAConfig specifies file name with ca-config
	CAConfig *string
	// Csr specifies file name with pem-encoded CSR
	Csr *string
	// SAN specifies coma separated alt names for generated cert
	SAN *string
	// Profile specifies the profile name from ca-config
	Profile *string
	// Output specifies the optional prefix for output files,
	// if not set, the output will be printed to STDOUT only
	Output *string
}

// Sign a certificate
func Sign(c ctl.Control, p interface{}) error {
	flags := p.(*SignFlags)

	cryptoprov, _ := c.(*cli.Cli).CryptoProv()
	if cryptoprov == nil {
		return errors.Errorf("unsupported command for this crypto provider")
	}

	// Load CSR
	csrPEM, err := c.(*cli.Cli).ReadFileOrStdin(*flags.Csr)
	if err != nil {
		return errors.Annotate(err, "read CSR")
	}

	// Load ca-config
	cacfg, err := authority.LoadConfig(*flags.CAConfig)
	if err != nil {
		return errors.Annotate(err, "ca-config")
	}
	err = cacfg.Validate()
	if err != nil {
		return errors.Annotate(err, "invalid ca-config")
	}

	isscfg := &authority.IssuerConfig{
		CertFile: *flags.CACert,
		KeyFile:  *flags.CAKey,
		Profiles: cacfg.Profiles,
	}

	issuer, err := authority.NewIssuer(isscfg, cryptoprov)
	if err != nil {
		return errors.Annotate(err, "create issuer")
	}

	var san []string
	if flags.SAN != nil && len(*flags.SAN) > 0 {
		san = strings.Split(*flags.SAN, ",")
	}
	signReq := csr.SignRequest{
		SAN:     san,
		Request: string(csrPEM),
		Profile: *flags.Profile,
	}

	_, certPEM, err := issuer.Sign(signReq)
	if err != nil {
		return errors.Annotate(err, "sign request")
	}

	if *flags.Output == "" {
		print.CSRandCert(c.Writer(), nil, nil, certPEM)
	} else {
		err = SaveCert(*flags.Output, nil, nil, certPEM)
		if err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}
