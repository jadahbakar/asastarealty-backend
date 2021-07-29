package app_test

import (
	"path/filepath"
	"testing"

	"github.com/jadahbakar/asastarealty-backend/app"
	"github.com/jadahbakar/asastarealty-backend/internal/config"
	"github.com/stretchr/testify/assert"
)

func Test_GetEngine(t *testing.T) {
	// setup path to env
	envPath, err := filepath.Abs("../")
	assert.NoError(t, err)

	// setup config
	config, err := config.New(envPath)
	assert.NoError(t, err)
	config.App.LogFolder = filepath.Join(envPath, "/log")

	// setup database
	// if you want to

	// testing
	apps := app.New(config, nil)
	assert.NotNil(t, apps)

	// testing
	engine := apps.GetEngine()
	assert.NotNil(t, engine)

	// testing
	logger := apps.GetLogger()
	assert.NotNil(t, logger)

	// testing
	configure := apps.GetConfig()
	assert.NotNil(t, configure)

	// testing
	database := apps.GetDB()
	assert.Nil(t, database)
}
