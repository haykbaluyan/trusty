package appcontainer

import (
	"io"
	"os"
	"time"

	"github.com/ekspand/trusty/authority"
	"github.com/ekspand/trusty/client"
	"github.com/ekspand/trusty/internal/config"
	"github.com/ekspand/trusty/internal/db"
	"github.com/ekspand/trusty/pkg/awskmscrypto"
	"github.com/ekspand/trusty/pkg/jwt"
	"github.com/ekspand/trusty/pkg/oauth2client"
	"github.com/go-phorce/dolly/audit"
	fauditor "github.com/go-phorce/dolly/audit/log"
	"github.com/go-phorce/dolly/tasks"
	"github.com/go-phorce/dolly/xpki/cryptoprov"
	"github.com/juju/errors"
	"github.com/sony/sonyflake"
	"go.uber.org/dig"
)

// ContainerFactoryFn defines an app container factory interface
type ContainerFactoryFn func() (*dig.Container, error)

// ProvideConfigurationFn defines Configuration provider
type ProvideConfigurationFn func() (*config.Configuration, error)

// ProvideDiscoveryFn defines Discovery provider
type ProvideDiscoveryFn func() (Discovery, error)

// ProvideAuditorFn defines Auditor provider
type ProvideAuditorFn func(cfg *config.Configuration, r CloseRegistrator) (audit.Auditor, error)

// ProvideSchedulerFn defines Scheduler provider
type ProvideSchedulerFn func() (tasks.Scheduler, error)

// ProvideJwtFn defines JWT provider
type ProvideJwtFn func(cfg *config.Configuration) (jwt.Parser, jwt.Provider, error)

// ProvideOAuthClientsFn defines OAuth clients provider
type ProvideOAuthClientsFn func(cfg *config.Configuration) (*oauth2client.Provider, error)

// ProvideCryptoFn defines Crypto provider
type ProvideCryptoFn func(cfg *config.Configuration) (*cryptoprov.Crypto, error)

// ProvideAuthorityFn defines Crypto provider
type ProvideAuthorityFn func(cfg *config.Configuration, crypto *cryptoprov.Crypto) (*authority.Authority, error)

// ProvideDbFn defines DB provider
type ProvideDbFn func(cfg *config.Configuration) (db.OrgsDb, db.CertsDb, error)

// ProvideClientFactoryFn defines client.Facroty provider
type ProvideClientFactoryFn func(cfg *config.Configuration) (client.Factory, error)

// CloseRegistrator provides interface to release resources on close
type CloseRegistrator interface {
	OnClose(closer io.Closer)
}

// ContainerFactory is default implementation
type ContainerFactory struct {
	closer CloseRegistrator

	configProvider        ProvideConfigurationFn
	discoveryProvider     ProvideDiscoveryFn
	auditorProvider       ProvideAuditorFn
	schedulerProvider     ProvideSchedulerFn
	cryptoProvider        ProvideCryptoFn
	authorityProvider     ProvideAuthorityFn
	dbProvider            ProvideDbFn
	oauthProvider         ProvideOAuthClientsFn
	jwtProvider           ProvideJwtFn
	clientFactoryProvider ProvideClientFactoryFn
}

// NewContainerFactory returns an instance of ContainerFactory
func NewContainerFactory(closer CloseRegistrator) *ContainerFactory {
	f := &ContainerFactory{
		closer: closer,
	}

	defaultSchedulerProv := func() (tasks.Scheduler, error) {
		return tasks.NewScheduler(), nil
	}

	// configure with default providers
	return f.
		WithDiscoveryProvider(provideDiscovery).
		WithAuditorProvider(provideAuditor).
		WithSchedulerProvider(defaultSchedulerProv).
		WithCryptoProvider(provideCrypto).
		WithAuthorityProvider(provideAuthority).
		WithDbProvider(provideDB).
		WithOAuthClientsProvider(provideOAuth).
		WithJwtProvider(provideJwt).
		WithClientFactoryProvider(provideClientFactory)
}

// WithConfigurationProvider allows to specify configuration
func (f *ContainerFactory) WithConfigurationProvider(p ProvideConfigurationFn) *ContainerFactory {
	f.configProvider = p
	return f
}

// WithDiscoveryProvider allows to specify Discovery
func (f *ContainerFactory) WithDiscoveryProvider(p ProvideDiscoveryFn) *ContainerFactory {
	f.discoveryProvider = p
	return f
}

// WithClientFactoryProvider allows to specify custom client.Factory provider
func (f *ContainerFactory) WithClientFactoryProvider(p ProvideClientFactoryFn) *ContainerFactory {
	f.clientFactoryProvider = p
	return f
}

// WithJwtProvider allows to specify custom JWT provider
func (f *ContainerFactory) WithJwtProvider(p ProvideJwtFn) *ContainerFactory {
	f.jwtProvider = p
	return f
}

// WithOAuthClientsProvider allows to specify custom OAuth clients provider
func (f *ContainerFactory) WithOAuthClientsProvider(p ProvideOAuthClientsFn) *ContainerFactory {
	f.oauthProvider = p
	return f
}

// WithAuditorProvider allows to specify custom Auditor
func (f *ContainerFactory) WithAuditorProvider(p ProvideAuditorFn) *ContainerFactory {
	f.auditorProvider = p
	return f
}

// WithDbProvider allows to specify custom DB provider
func (f *ContainerFactory) WithDbProvider(p ProvideDbFn) *ContainerFactory {
	f.dbProvider = p
	return f
}

// WithSchedulerProvider allows to specify custom Scheduler
func (f *ContainerFactory) WithSchedulerProvider(p ProvideSchedulerFn) *ContainerFactory {
	f.schedulerProvider = p
	return f
}

// WithCryptoProvider allows to specify custom Crypto loader
func (f *ContainerFactory) WithCryptoProvider(p ProvideCryptoFn) *ContainerFactory {
	f.cryptoProvider = p
	return f
}

// WithAuthorityProvider allows to specify custom Authority
func (f *ContainerFactory) WithAuthorityProvider(p ProvideAuthorityFn) *ContainerFactory {
	f.authorityProvider = p
	return f
}

// CreateContainerWithDependencies returns an instance of Container
func (f *ContainerFactory) CreateContainerWithDependencies() (*dig.Container, error) {
	container := dig.New()

	err := container.Provide(f.configProvider)
	if err != nil {
		return nil, errors.Trace(err)
	}

	container.Provide(func() CloseRegistrator {
		return f.closer
	})

	err = container.Provide(f.discoveryProvider)
	if err != nil {
		return nil, errors.Trace(err)
	}

	err = container.Provide(f.schedulerProvider)
	if err != nil {
		return nil, errors.Trace(err)
	}

	err = container.Provide(f.auditorProvider)
	if err != nil {
		return nil, errors.Trace(err)
	}

	err = container.Provide(f.jwtProvider)
	if err != nil {
		return nil, errors.Trace(err)
	}

	err = container.Provide(f.oauthProvider)
	if err != nil {
		return nil, errors.Trace(err)
	}

	err = container.Provide(f.cryptoProvider)
	if err != nil {
		return nil, errors.Trace(err)
	}

	err = container.Provide(f.authorityProvider)
	if err != nil {
		return nil, errors.Trace(err)
	}

	err = container.Provide(f.dbProvider)
	if err != nil {
		return nil, errors.Trace(err)
	}
	err = container.Provide(f.clientFactoryProvider)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return container, nil
}

const (
	nullDevName = "/dev/null"
)

func provideDiscovery() (Discovery, error) {
	return NewDiscovery(), nil
}

func provideAuditor(cfg *config.Configuration, r CloseRegistrator) (audit.Auditor, error) {
	var auditor audit.Auditor
	if cfg.Audit.Directory != "" && cfg.Audit.Directory != nullDevName {
		os.MkdirAll(cfg.Audit.Directory, 0644)

		// create auditor
		var err error
		auditor, err = fauditor.New(cfg.ServiceName+".log", cfg.Audit.Directory, cfg.Audit.MaxAgeDays, cfg.Audit.MaxSizeMb)
		if err != nil {
			logger.Errorf("reason=auditor, err=[%v]", errors.ErrorStack(err))
			return nil, errors.Annotate(err, "failed to create Auditor")
		}
	} else {
		auditor = auditornoop{}
	}
	if r != nil {
		r.OnClose(auditor)
	}
	return auditor, nil
}

func provideJwt(cfg *config.Configuration) (jwt.Parser, jwt.Provider, error) {
	var provider jwt.Provider
	var err error
	if cfg.JWT != "" {
		provider, err = jwt.Load(cfg.JWT)
		if err != nil {
			return nil, nil, errors.Trace(err)
		}
	}

	return provider, provider, nil
}

func provideOAuth(cfg *config.Configuration) (*oauth2client.Provider, error) {
	p, err := oauth2client.NewProvider(cfg.OAuthClients)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return p, nil
}

func provideCrypto(cfg *config.Configuration) (*cryptoprov.Crypto, error) {
	cryptoprov.Register("AWSKMS", awskmscrypto.KmsLoader)
	crypto, err := cryptoprov.Load(cfg.CryptoProv.Default, cfg.CryptoProv.Providers)
	if err != nil {
		logger.Errorf("default=%s, providers=%v, err=[%v]",
			cfg.CryptoProv.Default, cfg.CryptoProv.Providers,
			errors.ErrorStack(err))
		return nil, errors.Trace(err)
	}
	return crypto, nil
}

func provideAuthority(cfg *config.Configuration, crypto *cryptoprov.Crypto) (*authority.Authority, error) {
	caCfg, err := authority.LoadConfig(cfg.Authority)
	if err != nil {
		return nil, errors.Trace(err)
	}
	ca, err := authority.NewAuthority(caCfg, crypto)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return ca, nil
}

func provideDB(cfg *config.Configuration) (db.OrgsDb, db.CertsDb, error) {
	var idGenerator = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		/* TODO: machine ID from config
		MachineID: func() (uint16, error) {
			return uint16(os.Getpid()), nil
		},
		*/
	})

	d, err := db.New(cfg.SQL.Driver, cfg.SQL.DataSource, cfg.SQL.MigrationsDir, idGenerator.NextID)
	if err != nil {
		return nil, nil, errors.Trace(err)
	}
	return d, d, nil
}

func provideClientFactory(cfg *config.Configuration) (client.Factory, error) {
	return client.NewFactory(&cfg.TrustyClient), nil
}
