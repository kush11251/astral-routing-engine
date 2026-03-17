package controllers

import (
    "astral-routing-engine/src/models"
    "encoding/json"
    "net/http"
)

func GetRoutes(w http.ResponseWriter, r *http.Request) {
    routes := models.GetRoutes()
    json.NewEncoder(w).Encode(routes)
}

func CreateRoute(w http.ResponseWriter, r *http.Request) {
    var route models.Route
    _ = json.NewDecoder(r.Body).Decode(&route)
    models.CreateRoute(route)
    w.WriteHeader(http.StatusCreated)
}
