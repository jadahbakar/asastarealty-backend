package bod_test

import (
	"io/ioutil"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/jadahbakar/asastarealty-backend/app"
	"github.com/jadahbakar/asastarealty-backend/internal/config"
	"github.com/jadahbakar/asastarealty-backend/internal/infrastructure/postgresql"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	// setup path to env
	envPath, err := filepath.Abs("../../../../")
	assert.NoError(t, err)

	// setup config
	config, err := config.New(envPath)
	assert.NoError(t, err)
	config.App.LogFolder = filepath.Join(envPath, "/log")

	// setup db
	db, err := postgresql.NewPgClient(config)
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// setup app
	apps := app.New(config, db)
	assert.NotNil(t, apps)

	// get engine
	eng := apps.GetEngine()
	assert.NotNil(t, eng)

	// define test route or handler
	req := httptest.NewRequest("GET", "/api/v1/bod", nil)

	//  run integration test
	resp, err := eng.Test(req)
	assert.NoError(t, err)

	// transform response body to string
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	// assert expectation
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, string(body))
}

func TestGetById(t *testing.T) {
	// setup path to env
	envPath, err := filepath.Abs("../../../../")
	assert.NoError(t, err)

	// setup config
	config, err := config.New(envPath)
	assert.NoError(t, err)
	config.App.LogFolder = filepath.Join(envPath, "/log")

	// setup db
	db, err := postgresql.NewPgClient(config)
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// setup app
	apps := app.New(config, db)
	assert.NotNil(t, apps)

	// get engine
	eng := apps.GetEngine()
	assert.NotNil(t, eng)

	// define test route or handler
	req := httptest.NewRequest("GET", "/api/v1/bod/1", nil)

	//  run integration test
	resp, err := eng.Test(req)
	assert.NoError(t, err)

	// transform response body to string
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	// assert expectation
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, string(body))
}
