package main

import (
	"encoding/json"
	"net/http"
)

func GetRelation(url string) (Relation, error) {
	response, err := http.Get(url)
	if err != nil {
		return Relation{}, err
	}

	defer response.Body.Close()

	var relation Relation
	if err := json.NewDecoder(response.Body).Decode(&relation); err != nil {
		return Relation{}, err
	}
	return relation, nil

}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
