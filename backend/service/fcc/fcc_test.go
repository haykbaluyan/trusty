package fcc_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sync"
	"syscall"
	"testing"
	"time"

	v1 "github.com/ekspand/trusty/api/v1"
	"github.com/ekspand/trusty/backend/service/fcc"
	"github.com/ekspand/trusty/backend/trustymain"
	"github.com/ekspand/trusty/internal/config"
	"github.com/ekspand/trusty/pkg/gserver"
	"github.com/ekspand/trusty/tests/testutils"
	"github.com/go-phorce/dolly/testify/servefiles"
	"github.com/go-phorce/dolly/xhttp/marshal"
	"github.com/juju/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	trustyServer *gserver.Server
)

const (
	projFolder = "../../../"
)

func TestMain(m *testing.M) {
	var err error
	cfg, err := testutils.LoadConfig(projFolder, "UNIT_TEST")
	if err != nil {
		panic(errors.Trace(err))
	}

	httpAddr := testutils.CreateURLs("http", "")
	for name, httpCfg := range cfg.HTTPServers {
		switch name {
		case config.WFEServerName:
			httpCfg.Services = []string{fcc.ServiceName}
			httpCfg.ListenURLs = []string{httpAddr}
			httpCfg.Disabled = false
		default:
			// disable other servers
			httpCfg.Disabled = true
		}
	}

	sigs := make(chan os.Signal, 2)

	app := trustymain.NewApp([]string{}).
		WithConfiguration(cfg).
		WithSignal(sigs)

	var wg sync.WaitGroup
	startedCh := make(chan bool)

	var rc int
	var expError error

	go func() {
		defer wg.Done()
		wg.Add(1)

		expError = app.Run(startedCh)
		if expError != nil {
			startedCh <- false
		}
	}()

	// wait for start
	select {
	case ret := <-startedCh:
		if ret {
			trustyServer = app.Server(config.WFEServerName)
			if trustyServer == nil {
				panic("wfe not found!")
			}

			// Run the tests
			rc = m.Run()

			// trigger stop
			sigs <- syscall.SIGTERM
		}

	case <-time.After(20 * time.Second):
		break
	}

	// wait for stop
	wg.Wait()

	os.Exit(rc)
}

func Test_FccFrnHandler(t *testing.T) {
	service := trustyServer.Service(fcc.ServiceName).(*fcc.Service)
	require.NotNil(t, service)

	h := service.FccFrnHandler()

	server := servefiles.New(t)
	server.SetBaseDirs("testdata")

	u, err := url.Parse(server.URL() + "/")
	require.NoError(t, err)

	service.FccBaseURL = u.Scheme + "://" + u.Host

	t.Run("no_filer_id", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodGet, v1.PathForFccFrn, nil)
		require.NoError(t, err)

		h(w, r, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "{\"code\":\"invalid_request\",\"message\":\"missing filer_id parameter\"}", w.Body.String())
	})

	t.Run("wrong_filer_id", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodGet, v1.PathForFccFrn+"?filer_id=wrong", nil)
		require.NoError(t, err)

		h(w, r, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "{\"code\":\"invalid_request\",\"message\":\"unknown failure. Please check server logs.\"}", w.Body.String())
	})

	t.Run("url", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodGet, v1.PathForFccFrn+"?filer_id=831188", nil)
		require.NoError(t, err)

		h(w, r, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		var res v1.FccFrnResponse
		require.NoError(t, marshal.Decode(w.Body, &res))
		require.NotNil(t, res)
		assert.Equal(t, "0024926677", res.FRN)
	})
}

func Test_FccSearchDetailHandler(t *testing.T) {
	service := trustyServer.Service(fcc.ServiceName).(*fcc.Service)
	require.NotNil(t, service)

	h := service.FccSearchDetailHandler()

	server := servefiles.New(t)
	server.SetBaseDirs("testdata")

	u, err := url.Parse(server.URL() + "/")
	require.NoError(t, err)

	service.FccBaseURL = u.Scheme + "://" + u.Host

	t.Run("no_frn", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodGet, v1.PathForFccSearchDetail, nil)
		require.NoError(t, err)

		h(w, r, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "{\"code\":\"invalid_request\",\"message\":\"missing frn parameter\"}", w.Body.String())
	})

	t.Run("wrong_frn", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodGet, v1.PathForFccSearchDetail+"?frn=wrong", nil)
		require.NoError(t, err)

		h(w, r, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "{\"code\":\"invalid_request\",\"message\":\"unknown failure. Please check server logs.\"}", w.Body.String())
	})

	t.Run("url", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodGet, v1.PathForFccSearchDetail+"?frn=0024926677", nil)
		require.NoError(t, err)

		h(w, r, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		var res v1.FccSearchDetailResponse
		require.NoError(t, marshal.Decode(w.Body, &res))
		require.NotNil(t, res)
		assert.Equal(t, "tara.lyle@veracitynetworks.com", res.Email)
	})
}
