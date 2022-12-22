package main

import (
	"encoding/json"
	"fmt"
)

type Fruit struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func UnmarshalFruits(jsonStr string) ([]Fruit, error) {
	fruits := new([]Fruit)
	jsonBytes := []byte(jsonStr)
	if err := json.Unmarshal(jsonBytes, fruits); err != nil {
		return nil, fmt.Errorf("cannot unmarshal input string: %w", err)
	}
	return *fruits, nil
}

func MarshalFruits(fruits []Fruit) (string, error) {
  jsonData, err := json.Marshal(fruits)
  if err != nil {
		return "", fmt.Errorf("cannnot marshal input fruits: %w", err)
	}
	jsonStr := string(jsonData)
	return jsonStr, nil
}
