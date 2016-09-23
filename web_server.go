package main

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

func main() {  
    // Instantiate a new router
    r := httprouter.New()

    // Add a handler on /status
    r.GET("/status", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        // Simply write some test data for now
        fmt.Fprint(w, "live!\n")
    })

    // Get resource user by id
    r.GET("/user/:id", func (w http.ResponseWriter, r *http.Request, p httprouter.Params) {

        // Simply write some test data for now
        found_user := models.User{
            Name:   "John Doe",
            Gender: "male",
            Age:    34,
            Id:     p.ByName("id"),
        }

        // Marshal provided interface into JSON structure
        uj, _ := json.Marshal(found_user)

        // Write content-type, stautscode and payload
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(200)
        fmt.Fprintf(w, "%s", uj)
    })

    // Creates resource user by id
    r.POST("/user", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {  
        // Stub an user to be populated from the body
        new_user := models.User{}

        // Populate the user data
        json.NewDecoder(r.Body).Decode(&new_user)

        // Add an Id
        new_user.Id = "foo"

        // Marshal provided interface into JSON structure
        uj, _ := json.Marshal(new_user)

        // Write content-type, statuscode, payload
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(201)
        fmt.Fprintf(w, "%s", uj)
    })

    // Removes resource user by id
    r.DELETE("/user/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  

        // TODO: only write status for now
        w.WriteHeader(200)
    })

    // Fire up the server
    http.ListenAndServe("localhost:3000", r)
}
