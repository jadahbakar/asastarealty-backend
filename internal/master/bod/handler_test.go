package bod_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jadahbakar/asastarealty-backend/app"
	"github.com/stretchr/testify/assert"
)

// https://github.com/gofiber/recipes/blob/master/unit-test/main_test.go

func TestIndexRoute(t *testing.T) {
	// Define a structure for specifying input and output
	// data of a single test case. This structure is then used
	// to create a so called test map, which contains all test
	// cases, that should be run for testing this function
	tests := []struct {
		description string

		// Test input
		route string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "bod all",
			route:         "/bod",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "OK",
		},
		// {
		// 	description:   "bod by id",
		// 	route:         "/bod/:id",
		// 	expectedError: false,
		// 	expectedCode:  200,
		// 	expectedBody:  "OK",
		// },
	}

	// Setup the app as it is done in the main function
	apps := app.SetupApp()
	engine := apps.GetEngine()
	fmt.Printf("Value for engine -> %v", engine)
	// engine := fiber.New(fiber.Config{})

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route
		// from the test case
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res, err := engine.Test(req, -1)

		// verify that no error occured, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		// Read the response body
		body, err := ioutil.ReadAll(res.Body)

		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)

		// Verify, that the reponse body equals the expected body
		assert.Equalf(t, test.expectedBody, string(body), test.description)
	}
}

// func TestGetAll(t *testing.T) {
// 	var mockBod bod.Bod
// 	err := faker.FakeData(&mockBod)
// 	assert.NoError(t, err)
// 	mockService := new(mocks.BodService)
// 	mockListBod := make([]bod.Bod, 0)
// 	mockListBod = append(mockListBod, mockBod)
// 	mockService.On("Fetch").Return(mockListBod, nil)

// 	f := fiber.New()
// 	req, err := http.NewRequest(fiber.MethodGet, "/bod/all", strings.NewReader(""))
// 	assert.NoError(t, err)

// 	rec := httptest.NewRecorder()
// 	// c := fiber.Ctx(req, rec)
// 	handler := bod.BodHandler{
// 		bodService: mockService,
// 	}
// 	err = handler.GetAll(ctx)
// 	require.NoError(t, err)

// }
