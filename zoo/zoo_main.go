package zoo

import "fmt"

func zoo_main() {
	animalList := map[string][]string{
		"Elephant": {"Jack", "Alphred"},
		"Giraffe":  {"Steeve"},
		"Camel":    {"Fred"},
		"Tiger":    {"Eric"},
		"Lion":     {"Alex"},
		"Monkey":   {"Kevin"},
	}

	animals := CreateAnimalsFromMap(animalList)
	jails := CreateJailsFromMap(animalList)
	George := Zookeeper{Name: "George"}
	kpiZoo := Zoo{Name: "KPI Zoo", Zookeeper: George, Animals: animals, Jails: jails}
	fmt.Printf("В зоопарку %s, %d пустих кліток і %d тварин на волі.\n", kpiZoo.Name, len(jails), len(animals))
	for i, v := range animals {
		George.Imprison(&v, &jails[i])
	}
	fmt.Printf("В зоопарку %s, всі в клітках", kpiZoo.Name)
}
