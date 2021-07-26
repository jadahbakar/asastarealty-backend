package app_test

import (
	"testing"

	"github.com/jadahbakar/asastarealty-backend/app"
	"github.com/stretchr/testify/assert"
)

func Test_GetEngine(t *testing.T) {
	apps := app.New(nil, nil)
	assert.NotNil(t, apps)
	// engine := apps.GetEngine()
	// assert.NotNil(t, engine)
}
