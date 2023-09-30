package main

import "fmt"

type Developer struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Filter Unique Challenge")
	devs := []Developer{
		{Name: "Elliot"},
		{Name: "Alan"},
		{Name: "Jennifer"},
		{Name: "Graham"},
		{Name: "Paul"},
		{Name: "Alan"},
	}

	fmt.Println(devs)
	result := FilterUnique(devs)
	fmt.Println(result)
}

func FilterUnique(developers []Developer) []string {
	var uniqueNames []string
	for _, v := range developers {
		if contains(uniqueNames, v.Name) {
			uniqueNames = append(uniqueNames, v.Name)
		}
	}

	return uniqueNames
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return false
		}
	}
	return true
}
