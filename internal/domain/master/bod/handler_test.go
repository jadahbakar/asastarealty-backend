package bod_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/asastarealty-backend/app"
	"github.com/jadahbakar/asastarealty-backend/internal/config"
	"github.com/jadahbakar/asastarealty-backend/internal/domain/master/bod"
	"github.com/jadahbakar/asastarealty-backend/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAll(t *testing.T) {
	var mockBod bod.Bod
	err := faker.FakeData(&mockBod)
	assert.NoError(t, err)
	mockService := new(mocks.BodService)
	mockListBod := make([]bod.Bod, 0)
	mockListBod = append(mockListBod, mockBod)
	mockService.On("Fetch").Return(mockListBod, nil)

	// f := fiber.New()
	_, b, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(b), "../../../../")
	// Setup Config
	config, err := config.New(root)
	if err != nil {
		log.Printf("error Config New -> %v\n", err)
	}
	assert.NotNil(t, config)
	apps := app.New(config, nil)
	assert.NotNil(t, apps)
	engine := apps.GetEngine()
	assert.NotNil(t, engine)
	logger := apps.GetLogger()
	assert.NotNil(t, logger)

	req, err := http.NewRequest(fiber.MethodGet, "/api/v1/bod", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	res, err := engine.Test(req, -1)
	// c := fiber.Ctx(req, rec)
	handler := bod.BodHandler{bod.BodService: mockService}
	// err = handler.GetAll(ctx)
	require.NoError(t, err)

}
