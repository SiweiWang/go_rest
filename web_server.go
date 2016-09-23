package main

import ( 
    // Standard library packages
    "net/http"

    // Our controller
    "controllers"

    // Third party packages
    "github.com/julienschmidt/httprouter"
)

func main() {  
    // Instantiate a new router
    router := httprouter.New()

    // Get a UserController instance
    userController := controllers.NewUserController()

    // Get a user resource
    router.GET("/user/:id", userController.GetUser)

    router.POST("/user", userController.CreateUser)

    router.DELETE("/user/:id", userController.RemoveUser)

    // Fire up the server
    http.ListenAndServe("localhost:3000", router)
}
