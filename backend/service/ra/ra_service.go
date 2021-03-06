package ra

import (
	"context"
	"sync"

	pb "github.com/ekspand/trusty/api/v1/pb"
	"github.com/ekspand/trusty/client"
	"github.com/ekspand/trusty/client/embed/proxy"
	"github.com/ekspand/trusty/internal/config"
	"github.com/ekspand/trusty/internal/db"
	"github.com/ekspand/trusty/internal/db/model"
	"github.com/ekspand/trusty/pkg/gserver"
	"github.com/ekspand/trusty/pkg/poller"
	"github.com/go-phorce/dolly/rest"
	"github.com/go-phorce/dolly/xlog"
	"github.com/go-phorce/dolly/xpki/certutil"
	"github.com/juju/errors"
	"google.golang.org/grpc"
)

// ServiceName provides the Service Name for this package
const ServiceName = "ra"

var logger = xlog.NewPackageLogger("github.com/ekspand/trusty/backend/service", "ra")

// Service defines the Status service
type Service struct {
	server        *gserver.Server
	db            db.CertsDb
	clientFactory client.Factory
	grpClient     *client.Client
	ca            client.CAClient
	registered    bool
	cfg           *config.Configuration
	lock          sync.RWMutex
	ctx           context.Context
	cancel        context.CancelFunc
}

// Factory returns a factory of the service
func Factory(server *gserver.Server) interface{} {
	if server == nil {
		logger.Panic("status.Factory: invalid parameter")
	}

	return func(cfg *config.Configuration, db db.CertsDb, clientFactory client.Factory) {
		svc := &Service{
			server:        server,
			cfg:           cfg,
			db:            db,
			clientFactory: clientFactory,
		}

		svc.ctx, svc.cancel = context.WithCancel(context.Background())

		server.AddService(svc)
	}
}

// Name returns the service name
func (s *Service) Name() string {
	return ServiceName
}

// IsReady indicates that the service is ready to serve its end-points
func (s *Service) IsReady() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.registered && s.ca != nil
}

// Close the subservices and it's resources
func (s *Service) Close() {
	s.cancel()

	if s.grpClient != nil {
		s.grpClient.Close()
	}
	logger.KV(xlog.INFO, "closed", ServiceName)
}

// RegisterRoute adds the Status API endpoints to the overall URL router
func (s *Service) RegisterRoute(r rest.Router) {
}

// RegisterGRPC registers gRPC handler
func (s *Service) RegisterGRPC(r *grpc.Server) {
	pb.RegisterRAServiceServer(r, s)
}

// OnStarted is called when the server started and
// is ready to serve requests
func (s *Service) OnStarted() error {
	err := s.registerRoots(context.Background())
	if err != nil {
		return errors.Trace(err)
	}

	p := poller.New(nil,
		func(ctx context.Context) (interface{}, error) {
			c, err := s.getCAClient()
			if err != nil {
				return nil, errors.Trace(err)
			}
			return c, nil
		},
		func(err error) {})
	p.Start(s.ctx, s.cfg.TrustyClient.DialKeepAliveTimeout)
	go s.getCAClient()

	return nil
}

func (s *Service) registerCert(ctx context.Context, trust pb.Trust, location string) error {
	crt, err := certutil.LoadFromPEM(location)
	if err != nil {
		return err
	}
	pem, err := certutil.EncodeToPEMString(true, crt)
	if err != nil {
		return err
	}
	c := model.NewRootCertificate(crt, int(trust), pem)
	_, err = s.db.RegisterRootCertificate(ctx, c)
	if err != nil {
		return err
	}
	logger.Infof("trust=%v, subject=%q", trust, c.Subject)
	return nil
}

func (s *Service) registerRoots(ctx context.Context) error {
	for _, r := range s.cfg.RegistrationAuthority.PrivateRoots {
		err := s.registerCert(ctx, pb.Trust_Private, r)
		if err != nil {
			logger.Errorf("err=[%v]", errors.ErrorStack(err))
			return err
		}
	}
	for _, r := range s.cfg.RegistrationAuthority.PublicRoots {
		err := s.registerCert(ctx, pb.Trust_Public, r)
		if err != nil {
			logger.Errorf("err=[%v]", errors.ErrorStack(err))
			return err
		}
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	s.registered = true

	return nil
}

func (s *Service) getCAClient() (client.CAClient, error) {
	var ca client.CAClient
	s.lock.RLock()
	ca = s.ca
	s.lock.RUnlock()
	if ca != nil {
		return ca, nil
	}

	var pb pb.CAServiceServer
	err := s.server.Discovery().Find(&pb)
	if err == nil {
		s.lock.Lock()
		defer s.lock.Unlock()
		s.ca = client.NewCAClientFromProxy(proxy.CAServerToClient(pb))
		return s.ca, nil
	}

	grpClient, err := s.clientFactory.NewClient("ca")
	if err != nil {
		logger.KV(xlog.ERROR,
			"status", "failed to get CA client",
			"err", errors.Details(err))
		return nil, errors.Trace(err)
	}
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.grpClient != nil {
		s.grpClient.Close()
	}
	s.grpClient = grpClient
	s.ca = grpClient.CAClient()

	logger.KV(xlog.INFO, "status", "created CA client")

	return s.ca, nil
}
