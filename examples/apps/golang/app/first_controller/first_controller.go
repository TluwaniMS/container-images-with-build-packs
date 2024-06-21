package first_controller

import (
	"basic-go-rest-api/auxiliaries"
	"basic-go-rest-api/auxiliary_types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)

func GetHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	message := auxiliaries.GenerateResponseMessage("Response from Get Method on the First Controller ;) ...!!")

	response.Write(message)
}

func PostHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	body, _ := ioutil.ReadAll(request.Body)

	var person auxiliary_types.Person

	error := json.Unmarshal(body, &person)

	if error != nil {
		fmt.Println("There was an error.")
	}

	message := auxiliaries.GenerateResponseMessage("Hello " + person.Name + " " + person.Surname + ", this is a response from the first controller.")

	response.Write(message)
}

func PutHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	id := mux.Vars(request)["id"]
	gender := mux.Vars(request)["gender"]

	message := auxiliaries.GenerateResponseMessage("The id provided is: " + id + " and the gender is: " + gender + " from the first controller.")

	response.Write(message)
}

func DeletHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	id := mux.Vars(request)["id"]
	gender := mux.Vars(request)["gender"]

	message := auxiliaries.GenerateResponseMessage("The id provided for deletion is :" + id + " and the gender is: " + gender + " from the first controller.")

	response.Write(message)
}
