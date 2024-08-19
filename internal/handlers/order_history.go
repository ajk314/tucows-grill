package handlers

import (
    "encoding/json"
    "math"
    "net/http"
    "strconv"
    "sync"
    
	"tucows-grill-service/internal/data"
)

const batchSize = 10 // small batch size cause the sample data is only 50 rows

func GetTotalCostForItem(repo *data.Repository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        itemIDStr := r.URL.Query().Get("item_id")
        if itemIDStr == "" {
            http.Error(w, "item_id query parameter is required", http.StatusBadRequest)
            return
        }

        itemID, err := strconv.Atoi(itemIDStr)
        if err != nil {
            http.Error(w, "Invalid item_id", http.StatusBadRequest)
            return
        }

        totalCost, err := repo.GetTotalCostForItem(itemID)
        if err != nil {
            http.Error(w, "Error querying the database", http.StatusInternalServerError)
            return
        }

        totalCost = math.Round(totalCost*100) / 100
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]float64{"total_cost": totalCost})
    }
}

// This method will give the same result as the method above. The only difference is
// it will query in batches and use concurrency to calculate the result. This is ideal
// for large queries, complicated, or expensive computations. In this case, batch size
// is 10 because the sample data is small.
func GetTotalCostForItemAsync(repo *data.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		itemIDStr := r.URL.Query().Get("item_id")
		if itemIDStr == "" {
			http.Error(w, "item_id query parameter is required", http.StatusBadRequest)
			return
		}

		itemID, err := strconv.Atoi(itemIDStr)
		if err != nil {
			http.Error(w, "Invalid item_id", http.StatusBadRequest)
			return
		}

		resultChan := repo.GetTotalCostForItemInBatches(itemID, batchSize)

		var totalCost float64
		var wg sync.WaitGroup

		// if a waitgroup is overkill, could just use another "done" channel or something
		// but if we have limits on go routines we want to use, then can control that here
		wg.Add(1)
		
		go func() {
			defer wg.Done()
			for cost := range resultChan {
				totalCost += cost
			}
		}()

		wg.Wait()

		totalCost = math.Round(totalCost*100) / 100
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]float64{"total_cost": totalCost})
	}
}