package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)



func TestHandler(t *testing.T)  {
	
	req, err := http.NewRequest("GET","",nil)



/// Incase of an error we fail and stop the test 

if err != nil{
	t.Fatal(err)
}

// create a mini recorder from the http library that will consume our request 

recorder := httptest.NewRecorder()


hf:= http.HandlerFunc(handler)


hf.ServeHTTP(recorder,req)

if status := recorder.Code; status != http.StatusOK{
	t.Errorf("Handler returned wrong status code : got  %v  want %v ", status ,http.StatusOK)
}
// check if the response body is what we expected 
expected := "Hello  Mahinya"

actual:= recorder.Body.String()
if actual != expected { t.Errorf("Handler returned wrong body  : got %v  expected %v",actual,expected)}}