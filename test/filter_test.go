package main

import (
	"lotr"
	"sort"
	"testing"
)

const (
	FEMALE_HOBBITS = 42
	FEMALE_ELFS    = 24

	RUNTIME_BETWEEN_150_250 = 5
)

func TestMatch(t *testing.T) {
	app := lotr.New(ACCESS_CODE)

	options := lotr.NewGetOptions()
	options.Match = map[string][]string{
		"gender": []string{"Female"},
		"race":   []string{"Hobbit", "Elf"},
	}

	charList, err := app.ListCharactersOptions(options)
	if err != nil {
		t.Fatal("Get error: ", err)
	}

	if charList.Total != FEMALE_ELFS+FEMALE_HOBBITS {
		t.Errorf("Filter: Female Hobbits and Elf : %d != %d (Should be)", charList.Total, FEMALE_HOBBITS+FEMALE_ELFS)
	}

	options.NotMatch = map[string][]string{
		"name": []string{"/\\(.*/", "/^A.*/i"}, // regex match - have a '(' in the name and don't start with A
	}

	charList, err = app.ListCharactersOptions(options)
	if err != nil {
		t.Fatal("Get error: ", err)
	}

	if charList.Total != 37 { // hard coding for now
		t.Errorf("Filter: Female Hobbits and Elf : %d != %d (Should be)", charList.Total, FEMALE_HOBBITS+FEMALE_ELFS)
	}
}

func TestInEquality(t *testing.T) {
	app := lotr.New(ACCESS_CODE)

	options := lotr.NewGetOptions()
	options.GreaterThan = map[string]int{
		"runtimeInMinutes": 150,
	}
	options.LessThan = map[string]int{
		"runtimeInMinutes": 250,
	}

	movieList, err := app.ListMoviesOptions(options)
	if err != nil {
		t.Fatal("Get error: ", err)
	}

	if movieList.Total != RUNTIME_BETWEEN_150_250 { // hard coding for now
		t.Errorf("Filter: Runtime between 150 and 250 : %d != %d (Should be)", movieList.Total, RUNTIME_BETWEEN_150_250)
	}
}

func TestSort(t *testing.T) {
	app := lotr.New(ACCESS_CODE)

	options := lotr.NewGetOptions()
	options.SortOnKey("name", true)

	movieList, err := app.ListMoviesOptions(options)
	if err != nil {
		t.Fatal("Get error: ", err)
	}

	var names []string
	for _, movie := range movieList.Movies {
		names = append(names, movie.Name)
	}

	if len(names) != 8 {
		t.Error("Wrong movie count")
	}

	isSorted := sort.SliceIsSorted(names, func(a, b int) bool {
		return names[a] <= names[b]
	})

	if !isSorted {
		t.Error("Movie List is not sorted on names")
	}
}
