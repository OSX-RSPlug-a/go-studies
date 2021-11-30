package main

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	ID        string  `json:"id"`
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status,omitempty"`
}

func main() {
	o := Order{
		ID:        "1",
		AccountID: "dudeBug",
		Amount:    200,
		Status:    "pending",
	}

	jsonOutput, _ := json.Marshal(o)

	fmt.Println(string(jsonOutput))

}
