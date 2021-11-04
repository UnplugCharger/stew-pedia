package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Stew struct {
	Dish        string `json:"stew"`
	Description string `json:description`
}

var dishes []Stew

func GetStewHandler(w http.ResponseWriter, r *http.Request) {
	// convert the data into json

	dishesListbytes, err := json.Marshal(dishes)

	// if there is any error print it to the console and return a server error to the user

	if err != nil {
		fmt.Println(fmt.Errorf("Error : %v ", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(dishesListbytes)
}


func CreateStewHandle(w http.ResponseWriter , r *http.Request)  {
	dish := Stew{}



	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error : %v ", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	dish.Dish= r.Form.Get("dishes")
	dish.Description= r.Form.Get("description")

	dishes = append(dishes, dish)

	http.Redirect(w,r,"/assets/", http.StatusFound) 
} 