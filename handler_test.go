package bod_test

import (
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/jadahbakar/asastarealty-backend/app"
	"github.com/jadahbakar/asastarealty-backend/internal/config"
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
			route:         "/api/v1/bod",
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

	// noted - (Bapak Shandy) di mock, tapi kalau memang tidak, tambahi parameter di Config untuk relative path
	// config := &config.Config{}
	// Get the Directory of Config File
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
	// assert.NotNil(t, engine)
	// logger := apps.GetLogger()
	// assert.NotNil(t, logger)

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
		log.Printf("%v", res)

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
		// body, err := ioutil.ReadAll(res.Body)

		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)

		// Verify, that the reponse body equals the expected body
		// assert.Equalf(t, test.expectedBody, string(body), test.description)
	}
}
