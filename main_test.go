package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "", nil)

	/// Incase of an error we fail and stop the test

	if err != nil {
		t.Fatal(err)
	}

	// create a mini recorder from the http library that will consume our request

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(handler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got  %v  want %v ", status, http.StatusOK)
	}
	// check if the response body is what we expected
	expected := "Hello  Mahinya"

	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("Handler returned wrong body  : got %v  expected %v", actual, expected)
	}
}

func TestRouter(t *testing.T) {
	r := newRouter()

	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/hello")

	// handle any unsuspected error

	if err != nil {
		t.Fatal(err)
	}

	// check if status code is 200 ok
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code should be ok , got %d", resp.StatusCode)
	}

	// read the response body convert it into string then compare

	defer resp.Body.Close()

	//// read the body into bytes
	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	/// if its fine

	respString := string(b)

	expected := "Hello  Mahinya"

	if respString != expected {
		t.Errorf("Response should be %s , got %s", expected, respString)
	}
}

func TestRouterForNonExistentRoutes(t *testing.T) {
	r := newRouter()

	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/hello")

	// handle any unsuspected error

	if err != nil {
		t.Fatal(err)
	}

	// check if status code is 405 ok
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code should be 405 , got %d", resp.StatusCode)
	}

	// read the response body convert it into string then compare

	defer resp.Body.Close()

	//// read the body into bytes
	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	/// if its fine

	respString := string(b)

	expected := "Hello  Mahinya"

	if respString != expected {
		t.Errorf("Response should be %s , got %s", expected, respString)
	}
}


func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	// We want to hit the `GET /assets/` route to get the index.html file response
	resp, err := http.Get(mockServer.URL + "/assets/")
	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 200 (ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}

	// It isn't wise to test the entire content of the HTML file.
	// Instead, we test that the content-type header is "text/html; charset=utf-8"
	// so that we know that an html file has been served
	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}

}