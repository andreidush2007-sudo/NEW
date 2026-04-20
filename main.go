package main

import (
	"fmt"
	"math/rand"
	"time"
)

type NameGenerator struct {
	prefixes []string
	suffixes []string
	used     map[string]bool
}

func NewNameGenerator() *NameGenerator {
	return &NameGenerator{
		prefixes: []string{"Swift", "Brave", "Calm", "Dark", "Echo", "Fierce", "Gentle", "Holy"},
		suffixes: []string{"Wolf", "Fox", "Bear", "Hawk", "Lion", "Tiger", "Falcon", "Raven"},
		used:     make(map[string]bool, 20),
	}
}

func (ng *NameGenerator) generateOne() string {
	idx1 := rand.Intn(len(ng.prefixes))
	idx2 := rand.Intn(len(ng.suffixes))
	return ng.prefixes[idx1] + ng.suffixes[idx2]
}

func (ng *NameGenerator) GenerateUnique(count int) []string {
	defer fmt.Println("Genetaion complite") 
	
	names := make([]string, 0, count) 
	
	for len(names) < count {
		candidate := ng.generateOne()
		if !ng.used[candidate] {
			ng.used[candidate] = true
			names = append(names, candidate)
		}
	}
	return names
}

func printNames(names []string, prefix string) {
	for i, name := range names {
		fmt.Printf("%s [%d] %s\n", prefix, i+1, name)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	gen := NewNameGenerator()
	
	uniqueNames := gen.GenerateUnique(10)
	printNames(uniqueNames, ".")
}