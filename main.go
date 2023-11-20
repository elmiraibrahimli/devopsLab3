package main

import (
    "devopslab3/db"
    "devopslab3/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    db.Connect()

    r := gin.Default()

    r.GET("/items", handlers.GetAllItems)
    r.GET("/items/:id", handlers.GetItem)
    r.POST("/items", handlers.CreateItem)
    r.PUT("/items/:id", handlers.UpdateItem)
    r.DELETE("/items/:id", handlers.DeleteItem)

    r.Run() // listen and serve on 0.0.0.0:8080
}





// package main

// import (
//     "log"
//     "net/http"

//     "github.com/gorilla/mux"
// 	"devopslab3/db"
// 	"devopslab3/handlers"
// )



// func main() {
//     InitDB() // Initialize the database connection

//     router := mux.NewRouter()
//     // Routes for CRUD operations
//     router.HandleFunc("/items", createItem).Methods("POST")
//     router.HandleFunc("/items", getAllItems).Methods("GET")
//     router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
//     router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

//     log.Fatal(http.ListenAndServe(":8085", router))
// }

