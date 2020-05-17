package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type response1 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

type batter struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
}

func main() {
	r, _ := json.Marshal(2.34)
	fmt.Println(string(r))

	fruits := []string{"a", "b", "c"}
	rFruits, _ := json.Marshal(fruits)
	fmt.Println(string(rFruits))

	res1D := &response1{
		Page:   1,
		Fruits: fruits,
	}

	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	var data map[string]interface{}

	fB, _ := ioutil.ReadFile("json/sample.json")
	if err := json.Unmarshal(fB, &data); err != nil {
		panic(err)
	}
	d := data["batters"].(map[string]interface{})["batter"].([]interface{})[0].(map[string]interface{})["id"]

	fmt.Println(d)

}
