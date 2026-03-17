package models

import (
    "encoding/json"
)

type Route struct {
    ID       string `json:"id"`
    Origin   string `json:"origin"`
    Destination string `json:"destination"`
}

func GetRoutes() []Route {
    // Simulate a database query
    return []Route{{ID: "1", Origin: "New York", Destination: "Los Angeles"}, {ID: "2", Origin: "Chicago", Destination: "Houston"}}
}

func CreateRoute(route Route) {
    // Simulate a database insertion
    fmt.Printf("Created route: %+v\n", route)
}
