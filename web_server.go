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

    // Handle resource user by id
    r.GET("/user/:id", func (w http.ResponseWriter, r *http.Request, p httprouter.Params) {

        // Simply write some test data for now
        new_user := models.User{
            Name:   "John Doe",
            Gender: "Male",
            Age:    34,
            Id:     p.ByName("id"),
        }

        // Marshal provided interface into JSON structure
        uj, _ := json.Marshal(new_user)

        // Write content-type, stautscode and payload
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(200)
        fmt.Fprintf(w, "%s", uj)
    })

    // Fire up the server
    http.ListenAndServe("localhost:3000", r)
}
