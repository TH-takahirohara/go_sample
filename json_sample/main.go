package main

import (
	"encoding/json"
	"fmt"
	"os"
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
		return "", fmt.Errorf("cannot marshal input fruits: %w", err)
	}
	jsonStr := string(jsonData)
	return jsonStr, nil
}

// ファイルから読み込んだ文字列をUnmarshalする関数
// os.Open, f.Read を使用
func UnmarshalFruitsFile(path string) ([]Fruit, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}
	defer f.Close()

	data := make([]byte, 1024)
	count, err := f.Read(data)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}

	fruits := new([]Fruit)
	if err := json.Unmarshal(data[:count], fruits); err != nil {
		return nil, fmt.Errorf("cannot unmarshal read file: %w", err)
	}
	return *fruits, nil
}
