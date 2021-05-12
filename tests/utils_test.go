package main_test

import (
  "github.com/stretchr/testify/assert"
  utils "github.com/wrasdf/asg-node-roller/services/utils"
  "testing"
)

func TestToJsonString(t *testing.T) {
  type FruitBasket struct {
      Name    string
      Fruit   []string
      Id      int64  `json:"ref"`
      private string
  }
  basket := FruitBasket {
      Name:    "Standard",
      Fruit:   []string{"Apple", "Banana", "Orange"},
      Id:      999,
      private: "Second-rate",
  }
  result, _ := utils.ToJsonString(basket)
  assert.Equal(t, result, `{"Name":"Standard","Fruit":["Apple","Banana","Orange"],"ref":999}`)
}

func TestJsonStringToMap(t *testing.T) {
  mapJson := `{"species":"pigeon", "description": "likes to perch on rocks", "dimensions":{"height": 24, "width": 10}, "regions": ["north pole", "South Pole"]}`
  result := utils.JsonStringToMap(mapJson)
  dimensions := result["dimensions"].(map[string]interface{})
  regions := result["regions"].([]interface{})
  assert.Equal(t, result["species"], "pigeon")
  assert.Equal(t, result["species"].(string), "pigeon")
  assert.Equal(t, result["description"], "likes to perch on rocks")
  assert.Equal(t, dimensions["height"], float64(24))
  assert.Equal(t, dimensions["height"].(float64), float64(24))
  assert.Equal(t, regions[0], "north pole")
}
