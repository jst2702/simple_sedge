package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	config "simplesedge.com/feed/pkg/config"
	mockdb "simplesedge.com/feed/pkg/db/mock"
	"simplesedge.com/gokit/util"
)

func newTestServer(t *testing.T, store *mockdb.MockStore) *Server {
	config := &config.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenduration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	// gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
