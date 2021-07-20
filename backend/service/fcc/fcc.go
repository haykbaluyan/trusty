package fcc

import (
	"context"
	"fmt"
	"net/http"

	v1 "github.com/ekspand/trusty/api/v1"
	"github.com/ekspand/trusty/pkg/gserver"
	"github.com/go-phorce/dolly/rest"
	"github.com/go-phorce/dolly/xhttp/httperror"
	"github.com/go-phorce/dolly/xhttp/marshal"
	"github.com/go-phorce/dolly/xhttp/retriable"
	"github.com/go-phorce/dolly/xlog"
)

// ServiceName provides the Service Name for this package
const ServiceName = "fcc"

var logger = xlog.NewPackageLogger("github.com/ekspand/trusty/backend/service", "fcc")

const (
	fccDefaultBaseURL = "https://apps.fcc.gov"
)

// Service defines the Status service
type Service struct {
	server *gserver.Server

	FccBaseURL string
}

// Factory returns a factory of the service
func Factory(server *gserver.Server) interface{} {
	if server == nil {
		logger.Panic("status.Factory: invalid parameter")
	}

	return func() {
		svc := &Service{
			server: server,
		}

		svc.FccBaseURL = fccDefaultBaseURL

		server.AddService(svc)
	}
}

// Name returns the service name
func (s *Service) Name() string {
	return ServiceName
}

// IsReady indicates that the service is ready to serve its end-points
func (s *Service) IsReady() bool {
	return true
}

// Close the subservices and it's resources
func (s *Service) Close() {
	logger.KV(xlog.INFO, "closed", ServiceName)
}

// RegisterRoute adds the Status API endpoints to the overall URL router
func (s *Service) RegisterRoute(r rest.Router) {
	r.GET(v1.PathForFccFrn, s.FccFrnHandler())
	r.GET(v1.PathForFccSearchDetail, s.FccSearchDetailHandler())
}

type fccResponseWriter struct {
	data []byte
}

func (w *fccResponseWriter) Write(data []byte) (int, error) {
	w.data = append(w.data, data...)
	return len(data), nil
}

// FccFrnHandler handles v1.PathForFccFrn
func (s *Service) FccFrnHandler() rest.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ rest.Params) {
		filerID, ok := r.URL.Query()["filer_id"]
		if !ok || len(filerID) != 1 || filerID[0] == "" {
			marshal.WriteJSON(w, r, httperror.WithInvalidRequest("missing filer_id parameter"))
			return
		}

		httpClient := retriable.New()
		resFromFcc := new(fccResponseWriter)
		path := fmt.Sprintf("/cgb/form499/499results.cfm?FilerID=%s&XML=TRUE", filerID[0])
		_, statusCode, err := httpClient.Request(context.Background(), "GET", []string{s.FccBaseURL}, path, nil, resFromFcc)
		if err != nil {
			logger.Errorf("fccBaseUrl=%q, path=%q, err=%q", s.FccBaseURL, path, err.Error())
			marshal.WriteJSON(w, r, httperror.WithInvalidRequest("unknown failure. Please check server logs."))
			return
		}

		if statusCode >= 400 {
			logger.Errorf("fccBaseUrl=%q, path=%q, statusCode=%d", s.FccBaseURL, path, statusCode)
			marshal.WriteJSON(w, r, httperror.WithInvalidRequest("unknown failure. Please check server logs."))
			return
		}

		fq := NewFiler499QueryResultsFromXML(string(resFromFcc.data))
		frn, err := fq.GetFRN()
		if err != nil {
			logger.Errorf("fccBaseUrl=%q, path=%q, respFromFCC=%q, err=%q", s.FccBaseURL, path, string(resFromFcc.data), err.Error())
			marshal.WriteJSON(w, r, httperror.WithInvalidRequest("unknown failure. Please check server logs."))
			return
		}

		res := v1.FccFrnResponse{
			FRN: frn,
		}

		logger.Tracef("filerID=%q, frn=%q, path=%q", filerID[0], frn, path)
		marshal.WriteJSON(w, r, res)
	}
}

// FccSearchDetailHandler handles v1.PathForFccSearchDetail
func (s *Service) FccSearchDetailHandler() rest.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ rest.Params) {
		frn, ok := r.URL.Query()["frn"]
		if !ok || len(frn) != 1 || frn[0] == "" {
			marshal.WriteJSON(w, r, httperror.WithInvalidRequest("missing frn parameter"))
			return
		}

		httpClient := retriable.New()

		resFromFcc := new(fccResponseWriter)
		path := fmt.Sprintf("/coresWeb/searchDetail.do?frn=%s", frn[0])
		_, _, err := httpClient.Request(context.Background(), "GET", []string{s.FccBaseURL}, path, nil, resFromFcc)
		if err != nil {
			logger.Errorf("fccBaseUrl=%q, path=%q, err=%q", s.FccBaseURL, path, err.Error())
			marshal.WriteJSON(w, r, httperror.WithInvalidRequest("unknown failure. Please check server logs."))
			return
		}

		sd := NewSearchDetailFromHTML(string(resFromFcc.data))
		email, err := sd.GetEmail()
		if err != nil {
			logger.Errorf("fccBaseUrl=%q, path=%q, respFromFCC=%q, err=%q", s.FccBaseURL, path, string(resFromFcc.data), err.Error())
			marshal.WriteJSON(w, r, httperror.WithInvalidRequest("unknown failure. Please check server logs."))
			return
		}

		res := v1.FccSearchDetailResponse{
			Email: email,
		}

		logger.Tracef("frn=%q, email=%q, path=%q", frn[0], email, path)

		marshal.WriteJSON(w, r, res)
	}
}
