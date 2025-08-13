package main

import (
	"fmt"
)

func GetGroupByCreationYear(year int) ([]Artist, error) {
	artistsData, err := Artists()
	if err != nil {
		fmt.Print(err.Error())
		return []Artist{}, err
	}

	var artistsByYear []Artist

	for _, artist := range artistsData {
		if artist.CreationDate == year {
			artistsByYear = append(artistsByYear, artist)
		}
	}
	return artistsByYear, nil
}
