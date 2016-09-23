package controllers

import (
      // Standard library packages
    "encoding/json"
    "fmt"
    "net/http"

    // Our models
    "models"

    // Third party packages
    "github.com/julienschmidt/httprouter"
)

type (
  // UserController represents the controller for operating on the User resource
  UserController struct{}
)

func NewUserController() *UserController {
  return &UserController{}
}

// GetUser retrieves an individual user resource
func (userController UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    // Simply write some test data for now
    found_user := models.User{
        Name:   "John Doe",
        Gender: "male",
        Age:    34,
        Id:     p.ByName("id"),
    }

    // Marshal provided interface into JSON structure
    user_json, _ := json.Marshal(found_user)

    // Write content-type, stautscode and payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", user_json)
}

// CreateUser creates a new user resource
func (userController UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
    // Stub an user to be populated from the body
    new_user := models.User{}

    // Populate the user data
    json.NewDecoder(r.Body).Decode(&new_user)

    // Add an Id
    new_user.Id = "foo"

    // Marshal provided interface into JSON structure
    user_json, _ := json.Marshal(new_user)

    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", user_json)
}

// RemoveUser removes an existing user resource
func (userController UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
    w.WriteHeader(200)
}