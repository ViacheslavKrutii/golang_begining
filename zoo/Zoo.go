package zoo

import "fmt"

type Zoo struct {
	Name      string
	Zookeeper Zookeeper
	Animals   []AnyAnimal
	Jails     []Jail
}

type Jail struct {
	species string
	empty   bool
	AnyAnimal
}

func JailConst(species string) Jail {
	return Jail{species: species, empty: true}
}

func CreateJailsFromMap(animalList map[string][]string) (jails []Jail) {
	for species, nameList := range animalList {
		for i := 0; i < len(nameList); i++ {
			jails = append(jails, JailConst(species))
		}

	}
	return jails
}

type Zookeeper struct {
	Name string
}

func (z Zookeeper) Imprison(a *AnyAnimal, j *Jail) {
	if j.empty {
		j.AnyAnimal = *a
		j.empty = false
		fmt.Println("Спіймав!!!")
	}
}
