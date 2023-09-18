package main

import (
	"Proj/HomeWork/golang_structs/zoo"
	"fmt"
)

/*
0. Declaration
1. Different types
3. Creating instances, instantiation
4. Zero values
5. Consecutive fields with same type
6. User defined type. Structs are the only way to create concrete user-defined types in Golang.
7. Struct Visibility
8. Fields Visibility
9. Embedding
10. Recursive (almost)
11. Anonymous/unnamed structs
12. Structural comparison
13. methods or pointer and value receivers
14. Tagging
15. Implicit dereferencing
16. Assignment
*/

func main() {
	animalList := map[string][]string{
		"Elephant": {"Jack", "Alphred"},
		"Giraffe":  {"Steeve"},
		"Camel":    {"Fred"},
		"Tiger":    {"Eric"},
		"Lion":     {"Alex"},
		"Monkey":   {"Kevin"},
	}

	animals := zoo.CreateAnimalsFromMap(animalList)
	jails := zoo.CreateJailsFromMap(animalList)
	George := zoo.Zookeeper{Name: "George"}
	kpiZoo := zoo.Zoo{Name: "KPI Zoo", Zookeeper: George, Animals: animals, Jails: jails}
	fmt.Printf("В зоопарку %s, %d пустих кліток і %d тварин на волі.\n", kpiZoo.Name, len(jails), len(animals))
	for i, v := range animals {
		George.Imprison(&v, &jails[i])
	}
	fmt.Printf("В зоопарку %s, всі в клітках", kpiZoo.Name)
}
