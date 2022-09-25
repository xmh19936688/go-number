package main

import (
	"encoding/json"
	"fmt"

	"github.com/xmh19936688/go-number/number"
)

type Weather struct {
	Temperature number.Number `json:"temperature"`
}

func main() {
	jsonStr := ``
	s := &Weather{}
	json.Unmarshal([]byte(jsonStr), s)

	if s.Temperature.Val == 0 && s.Temperature.Str == "0" {
		fmt.Println("temperature is 0")
	}

	if s.Temperature.Val == 0 && s.Temperature.Str == "" {
		fmt.Println("temperature is unknown")
	}
}
