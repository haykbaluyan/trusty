package testutils

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sync/atomic"

	"github.com/go-phorce/dolly/algorithms/guid"
)

var (
	nextPort    = int32(0)
	testDirPath = filepath.Join(os.TempDir(), "tests", "trusty", guid.MustCreate())
)

// CreateURLs returns URL with a random port
func CreateURLs(scheme, host string) string {
	if nextPort == 0 {
		nextPort = 17891 + int32(rand.Intn(5000))
	}
	next := atomic.AddInt32(&nextPort, 1)
	return fmt.Sprintf("%s://%s:%d", scheme, host, next)
}
