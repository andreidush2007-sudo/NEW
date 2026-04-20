package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Generator struct {
	ft []string
	lt []string
	used     map[string]bool
}

func gen() *Generator {
	return &Generator{
		ft: []string{"James", "John", "Robert", "Michael", "William", "David", "Richard", "Thomas"},
		lt:  []string{"Smith", "Johnson", "Brown", "Jones", "Garcia", "Miller", "Davis", "Wilson"},
		used:     make(map[string]bool, 20),
	}
}

func (ng *Generator) gen_names() string {
	idx1 := rand.Intn(len(ng.ft))
	idx2 := rand.Intn(len(ng.lt))
	return ng.ft[idx1] + " " + ng.lt[idx2]
}

func (ng *Generator) unique(count int) []string {
	defer fmt.Println("Genetaion complite") 
	
	names := make([]string, 0, count) 
	
	for len(names) < count {
		candidate := ng.gen_names()
		if !ng.used[candidate] {
			ng.used[candidate] = true
			names = append(names, candidate)
		}
	}
	return names
}

func print_all(names []string) {
	for i, name := range names {
		fmt.Printf("%d - %s\n",i+1, name)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	gen := gen()
	
	uniqueNames := gen.unique(10)
	print_all(uniqueNames)
}