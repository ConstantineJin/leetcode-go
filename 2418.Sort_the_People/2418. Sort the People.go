package main

import "sort"

type person struct {
	name   string
	height int
}

func sortPeople(names []string, heights []int) []string {
	people := make([]person, len(names))
	for i, name := range names {
		people[i] = person{name, heights[i]}
	}
	sort.Slice(people, func(i, j int) bool { return people[i].height > people[j].height })
	for i, p := range people {
		names[i] = p.name
	}
	return names
}
