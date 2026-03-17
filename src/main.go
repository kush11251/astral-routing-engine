package main

import (
    "astral-routing-engine/src/controllers"
    "astral-routing-engine/src/models"
    "astral-routing-engine/src/utils"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/routes", controllers.GetRoutes).Methods("GET")
    router.HandleFunc("/routes", controllers.CreateRoute).Methods("POST")
    log.Fatal(http.ListenAndServe(":8080", router))
}
