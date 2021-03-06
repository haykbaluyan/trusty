package cis

import (
	"context"
	"sync"

	pb "github.com/ekspand/trusty/api/v1/pb"
	"github.com/ekspand/trusty/client"
	"github.com/ekspand/trusty/client/embed/proxy"
	"github.com/ekspand/trusty/internal/config"
	"github.com/ekspand/trusty/internal/db"
	"github.com/ekspand/trusty/pkg/gserver"
	"github.com/ekspand/trusty/pkg/poller"
	"github.com/go-phorce/dolly/rest"
	"github.com/go-phorce/dolly/xlog"
	"github.com/juju/errors"
	"google.golang.org/grpc"
)

// ServiceName provides the Service Name for this package
const ServiceName = "cis"

var logger = xlog.NewPackageLogger("github.com/ekspand/trusty/backend/service", "cis")

// Service defines the Status service
type Service struct {
	server        *gserver.Server
	db            db.CertsDb
	cfg           *config.Configuration
	clientFactory client.Factory
	grpClient     *client.Client
	ra            client.RAClient
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
	return s.ra != nil
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
	pb.RegisterCIServiceServer(r, s)
}

// OnStarted is called when the server started and
// is ready to serve requests
func (s *Service) OnStarted() error {
	p := poller.New(nil,
		func(ctx context.Context) (interface{}, error) {
			c, err := s.getRAClient()
			if err != nil {
				return nil, errors.Trace(err)
			}
			return c, nil
		},
		func(err error) {})
	p.Start(s.ctx, s.cfg.TrustyClient.DialKeepAliveTimeout)
	go s.getRAClient()
	return nil
}

func (s *Service) getRAClient() (client.RAClient, error) {
	var ra client.RAClient
	s.lock.RLock()
	ra = s.ra
	s.lock.RUnlock()
	if ra != nil {
		return ra, nil
	}

	var pb pb.RAServiceServer
	err := s.server.Discovery().Find(&pb)
	if err == nil {
		s.lock.Lock()
		defer s.lock.Unlock()
		s.ra = client.NewRAClientFromProxy(proxy.RAServiceServerToClient(pb))
		return s.ra, nil
	}

	grpClient, err := s.clientFactory.NewClient("ra")
	if err != nil {
		logger.KV(xlog.ERROR,
			"status", "failed to get RA client",
			"err", errors.Details(err))
		return nil, errors.Trace(err)
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	if s.grpClient != nil {
		s.grpClient.Close()
	}
	s.grpClient = grpClient
	s.ra = grpClient.RAClient()

	logger.KV(xlog.INFO, "status", "created RA client")

	return s.ra, nil
}
