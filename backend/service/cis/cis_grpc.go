package cis

import (
	"context"

	pb "github.com/ekspand/trusty/api/v1/trustypb"
	"github.com/juju/errors"
)

// Roots returns the root CAs
func (s *Service) Roots(ctx context.Context, _ *pb.EmptyRequest) (*pb.RootsResponse, error) {
	roots, err := s.db.GetRootCertificates(ctx)
	if err != nil {
		logger.Errorf("src=Roots, err=[%v]", errors.ErrorStack(err))
		return nil, errors.Annotatef(err, "unable to query root certificates")
	}

	res := &pb.RootsResponse{
		Roots: roots.ToDTO(),
	}

	return res, nil
}