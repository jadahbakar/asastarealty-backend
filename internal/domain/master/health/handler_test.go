package health_test

import (
	"io/ioutil"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/jadahbakar/asastarealty-backend/app"
	"github.com/jadahbakar/asastarealty-backend/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestGetHealth(t *testing.T) {
	// setup path to env
	envPath, err := filepath.Abs("../../../../")
	assert.NoError(t, err)

	// setup config
	config, err := config.New(envPath)
	assert.NoError(t, err)
	config.App.LogFolder = filepath.Join(envPath, "/log")

	// setup app
	apps := app.New(config, nil)
	assert.NotNil(t, apps)

	// get engine
	eng := apps.GetEngine()
	assert.NotNil(t, eng)

	// define test route or handler
	req := httptest.NewRequest("GET", "/api/v1/health", nil)

	//  run integration test
	resp, err := eng.Test(req)
	assert.NoError(t, err)

	// transform response body to string
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	// assert expectation
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "{\"error\":false,\"message\":\"healthty\",\"data\":null}", string(body))
}
