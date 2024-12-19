package tests

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
	"catapp/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchCatImage(t *testing.T) {
	// Create a new HTTP request with breed ID as query parameter
	req, err := http.NewRequest("GET", "/cat/image?breed_id=beng", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Create a new Beego context and bind it to the request and response
	ctx := web.NewContext()
	ctx.Reset(recorder, req)

	// Initialize the controller
	controller := &controllers.CatController{}
	controller.Ctx = ctx
	controller.FetchCatImage()

	// Check that the response code is 200 OK
	assert.Equal(t, 200, recorder.Code)

	// Check that the body contains the expected image URL
	assert.Contains(t, recorder.Body.String(), "image_url")
}
