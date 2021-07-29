package bod_test

import (
	"log"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/jadahbakar/asastarealty-backend/app"
	"github.com/jadahbakar/asastarealty-backend/internal/config"
	"github.com/stretchr/testify/assert"
)

// func TestGetAll(t *testing.T) {
// 	// var mockBod bod.Bod
// 	// err := faker.FakeData(&mockBod)
// 	// assert.NoError(t, err)
// 	// mockService := new(mocks.BodService)
// 	// mockListBod := make([]bod.Bod, 0)
// 	// mockListBod = append(mockListBod, mockBod)
// 	// mockService.On("Fetch").Return(mockListBod, nil)

// 	// _, b, _, _ := runtime.Caller(0)
// 	// root := filepath.Join(filepath.Dir(b), "../../../../")
// 	// // Setup Config
// 	// config, err := config.New(root)
// 	// if err != nil {
// 	// 	log.Printf("error Config New -> %v\n", err)
// 	// }
// 	// assert.NotNil(t, config)
// 	// apps := app.New(config, nil)
// 	// assert.NotNil(t, apps)
// 	// engine := apps.GetEngine()
// 	// assert.NotNil(t, engine)
// 	// logger := apps.GetLogger()
// 	// assert.NotNil(t, logger)

// 	// // req, err := http.NewRequest(fiber.MethodGet, "/api/v1/bod", strings.NewReader(""))
// 	// req, err := http.NewRequest(fiber.MethodGet, "/api/v1/bod", nil)
// 	// assert.NoError(t, err)

// 	// rec := httptest.NewRecorder()

// 	// // res, err := engine.Test(req, -1)

// 	// // c := fiber.Ctx(req, rec)
// 	// // handler := bod.BodHandler{bod.BodService: mockService}

// 	// // err = handler.GetAll(ctx)
// 	// require.NoError(t, err)

// 	// Define Fiber config.
// 	config := fiber.Config{}

// 	// Define a new Fiber app with config.
// 	app := fiber.New(config)
// 	// Create route with GET method for test:
// 	app.Get("/api/v1", func(c *fiber.Ctx) error {
// 		return c.SendString("hello, customVerb!")
// 	})
// 	// http.Request
// 	req := httptest.NewRequest("GET", "/api/v1", nil)

// 	// http.Response
// 	resp, err := app.Test(req)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	// Do something with results:
// 	if resp.StatusCode == 200 {
// 		body, _ := ioutil.ReadAll(resp.Body)
// 		fmt.Println(string(body)) // => hello, customVerb!
// 	}
// 	if resp.StatusCode != 200 {
// 		t.Error(resp.StatusCode)
// 	}

// }

func TestGetAll(t *testing.T) {
	// Define Fiber config.
	// config := fiber.Config{}
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

}
