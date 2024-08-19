package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"

    "tucows-grill-service/internal/data"
    "tucows-grill-service/internal/models"
)

func GetIngredientByID(repo *data.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ingredientID, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid ingredient ID", http.StatusBadRequest)
			return
		}

		ingredient, err := repo.GetIngredientByID(ingredientID)
		if err != nil {
			http.Error(w, "Error fetching ingredient", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ingredient)
	}
}

func AddIngredient(repo *data.Repository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var ingredient models.Ingredient
        if err := json.NewDecoder(r.Body).Decode(&ingredient); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        if err := repo.AddIngredient(&ingredient); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
    }
}