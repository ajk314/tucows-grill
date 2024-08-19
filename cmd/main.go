package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "tucows-grill-service/internal/auth"
    "tucows-grill-service/internal/database"
    "tucows-grill-service/internal/data"
    "tucows-grill-service/internal/handlers"
)

func main() {
    db, err := database.Connect()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    repo := data.NewRepository(db)

    r := mux.NewRouter()
    protected := r.PathPrefix("/").Subrouter()

    r.HandleFunc("/login", handlers.Login).Methods("POST")

    protected.Use(auth.JWTMiddleware)
    // require JWT
    protected.HandleFunc("/ingredients/{id}", handlers.GetIngredientByID(repo)).Methods("GET")
    protected.HandleFunc("/ingredients", handlers.AddIngredient(repo)).Methods("POST")
    protected.HandleFunc("/total-cost", handlers.GetTotalCostForItem(repo)).Methods("GET")
    protected.HandleFunc("/total-cost-async", handlers.GetTotalCostForItemAsync(repo)).Methods("GET")

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
