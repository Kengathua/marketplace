package retail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Kengathua/marketplace/pkg/common"
	"github.com/Kengathua/marketplace/pkg/models"
	"github.com/Kengathua/marketplace/tests"
	"github.com/gofiber/fiber/v2"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func randomPort() int {
	rand.Seed(time.Now().Unix())
	min := 32768
	max := 60999
	port := rand.Intn(max-min+1) + min
	return port
}

type TestingT interface {
	Errorf(format string, args ...interface{})
}

func CreateTestUser(db *gorm.DB) models.User {
	var t TestingT
	assert := assert.New(t)

	user := models.User{
		BioData: common.BioData{
			FirstName:   "Test",
			LastName:    "User",
			Email:       "testuser@email.com",
			PhoneNumber: "+254712345678",
		},
		Password:    "TestPass123!",
		IsStaff:     true,
		IsAdmin:     true,
		IsSuperUser: true,
	}
	user.GeneratePasswordHarsh()
	err := db.Create(&user).Error

	if err != nil {
		assert.Nil(err)
	}

	err = db.First(&user, user.ID).Error
	if err != nil {
		assert.Nil(err)
	}

	return user
}

func InitializeTestServer(db *gorm.DB) *fiber.App {
	testUser := CreateTestUser(db)
	port := randomPort()
	baseURL := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", baseURL)

	if err != nil {
		panic(err)
	}

	testServer := httptest.NewUnstartedServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		_, err := res.Write([]byte(`{"message": "success two"`))
		if err != nil {
			panic("cannot return http response")
		}
	}))
	defer testServer.Close()

	testServer.Listener.Close()
	testServer.Listener = listener

	app := fiber.New()

	api := app.Group("/api")

	v1 := api.Group("/v1", func(c *fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v1")
		c.Locals("user", testUser)
		return c.Next()
	})
	retailURL := v1.Group("/retail", func(c *fiber.Ctx) error { // middleware for /api/v1/retail
		c.Set("Version", "v1")
		return c.Next()
	})

	RegisterRetailRoutes(retailURL, db)

	listeningPort := fmt.Sprintf(":%d", port)
	app.Listen(listeningPort)

	// Start the server.
	testServer.Start()
	// Stop the server on return from the function.
	defer testServer.Close()

	return app
}

type TestCases []struct {
	description  string                 // description of the test case
	route        string                 // route path to test
	expectedCode int                    // expected HTTP status code
	httpMethod   string                 // http method to be tested
	payload      map[string]interface{} // payload expected for the test case
}

func APITests(t *testing.T, testCases TestCases) {
	db := tests.GetTestDB()
	app := InitializeTestServer(db)

	for _, test := range testCases {
		req, err := http.NewRequest(test.httpMethod, test.route, nil)
		assert.Nil(t, err)

		req.Header.Set("Content-Type", "application/json")

		if test.httpMethod == "POST" || test.httpMethod == "PUT" {
			body := new(bytes.Buffer)
			json.NewEncoder(body).Encode(test.payload)
			req, err = http.NewRequest(test.httpMethod, test.route, body)
			assert.Nil(t, err)
			req.Header.Set("Content-Type", "application/json")
		}

		resp, err := app.Test(req)
		assert.Nil(t, err)

		defer resp.Body.Close()

		if resp.StatusCode != test.expectedCode {
			assert.Equal(t, test.expectedCode, resp.StatusCode)
		}

		_, err = ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)
	}
}
