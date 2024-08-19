package data

import (
    "database/sql"
    "tucows-grill-service/internal/models"

	"log"
)

type Repository struct {
    DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
    return &Repository{DB: db}
}

func (repo *Repository) GetIngredientByID(id int) (*models.Ingredient, error) {
	query := "SELECT id, name, quantity FROM ingredients WHERE id = ?"
	row := repo.DB.QueryRow(query, id)

	var ingredient models.Ingredient
	if err := row.Scan(&ingredient.ID, &ingredient.Name, &ingredient.Quantity); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &ingredient, nil
}

func (repo *Repository) AddIngredient(ingredient *models.Ingredient) error {
    query := "INSERT INTO ingredients (name, quantity) VALUES (?, ?)"
    _, err := repo.DB.Exec(query, ingredient.Name, ingredient.Quantity)
    return err
}

func (repo *Repository) GetTotalCostForItem(itemID int) (float64, error) {
    var totalCost float64
    query := `SELECT SUM(cost) FROM order_history WHERE item_id = ?`
    err := repo.DB.QueryRow(query, itemID).Scan(&totalCost)
    if err != nil {
        return 0, err
    }
    return totalCost, nil
}

func (repo *Repository) GetTotalCostForItemInBatches(itemID int, batchSize int) <-chan float64 {
	resultChan := make(chan float64)
	query := `SELECT cost FROM order_history WHERE item_id = ? LIMIT ? OFFSET ?`

	go func() {
		defer close(resultChan)

		offset := 0
		for {
			rows, err := repo.DB.Query(query, itemID, batchSize, offset)
			if err != nil {
				log.Printf("Error querying database: %v", err)
				return
			}

			batchCount := 0 // counts how many rows were actually returned
			for rows.Next() {
				var cost float64
				if err := rows.Scan(&cost); err != nil {
					log.Printf("Error scanning row: %v", err)
					rows.Close()
					return
				}
				resultChan <- cost
				batchCount++
			}
			rows.Close()

			// if the total number of rows returned were smaller than the batch size,
			// that means that we finished reading all the data. so break out.
			if batchCount < batchSize {
				break
			}

			offset += batchSize
		}
	}()

	return resultChan
}