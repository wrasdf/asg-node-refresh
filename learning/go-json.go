package main

import (
  "fmt"
  "time"
  "encoding/json"
)

type FruitBasket struct {
    Name    string
    Fruit   []string
    Id      int64  `json:"ref"`
    private string // An unexported field is not encoded.
    Created time.Time
}

func toJSONString(data interface{}) (string, error) {
	results, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(results), nil
}

// ==========================

type Dimensions struct {
  Height int
  Width int
}

type Bird struct {
  Species string
  Description string
  Dimensions Dimensions
}

func jsonToBirdStruct (s string) Bird {
  var bird Bird
  json.Unmarshal([]byte(s), &bird)
  return bird
}


// ========================

func jsonToMap (s string) map[string]interface{} {
  var results map[string]interface{}
  json.Unmarshal([]byte(s), &results)
  return results
}

// map[string]interface{} to store arbitrary JSON objects, and
// []interface{} to store arbitrary JSON arrays.

// ==================

func main() {

  basket := FruitBasket {
      Name:    "Standard",
      Fruit:   []string{"Apple", "Banana", "Orange"},
      Id:      999,
      private: "Second-rate",
      Created: time.Now(),
  }

  s, _ := toJSONString(basket)

  fmt.Println(s)

  birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`

  bird := jsonToBirdStruct(birdJson)
  fmt.Println(bird.Species)




  mapJson := `{"species":"pigeon", "description": "likes to perch on rocks", "dimensions":{"height": 24, "width": 10}, "level": 5, "regions": ["north pole", "South Pole"]}`
  r := jsonToMap(mapJson)
  fmt.Println(r)

  spec := r["species"].(string)
  dimensions := r["dimensions"].(map[string]interface{})
  regions := r["regions"].([]interface{})

  fmt.Println(spec)
  fmt.Println(dimensions["height"])
  fmt.Println(regions[1])

}
