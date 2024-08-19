package models

type Ingredient struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Quantity int    `json:"quantity"`
}
