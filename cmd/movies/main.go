package main

import (
	"encoding/json"
	"fmt"
	"go-tutorials/internal/movie"
	"log"
)

func main() {
	var movies = []movie.Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Bullitt", Year: 1968, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s\n", err)
	}
	fmt.Printf("%s\n", data)

	fmt.Println()

	data, err = json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s\n", err)
	}
	fmt.Printf("%s\n", data)

	fmt.Println()

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s\n", err)
	}
	fmt.Println(titles)
}
