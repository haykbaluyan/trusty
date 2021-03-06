package testutils

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sync/atomic"

	"github.com/ekspand/trusty/internal/config"
)

var (
	nextPort = int32(os.Getpid()%10000) + int32(17891) + rand.Int31n(1000)
)

// CreateURLs returns URL with a random port
func CreateURLs(scheme, host string) string {
	next := atomic.AddInt32(&nextPort, 1)
	return fmt.Sprintf("%s://%s:%d", scheme, host, next)
}

// LoadConfig returns Configuration
func LoadConfig(projFolder, hostname string) (*config.Configuration, error) {
	if hostname != "" {
		os.Setenv(config.EnvHostnameKey, hostname)
	}
	cfgPath, err := filepath.Abs(projFolder + "etc/dev/" + config.ConfigFileName)
	if err != nil {
		return nil, err
	}

	cfg, err := config.LoadConfigForHostName(cfgPath, hostname)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
