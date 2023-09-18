package zoo

import "fmt"

type animal struct {
	species string
	name    string
	size    string
	loose   bool
}

func (a animal) beAnimal() {

}

type AnyAnimal interface {
	beAnimal()
}

type Elephant struct {
	animal
}

type Camel struct {
	animal
}

type Tiger struct {
	animal
}

type Lion struct {
	animal
}

type Giraffe struct {
	animal
}

func AnimalConst(species, name string) AnyAnimal {
	switch species {
	case "Elephant":
		fmt.Println("Слон втік")
		return Elephant{animal: animal{species: species, name: name, size: "Large", loose: true}}
	case "Giraffe":
		fmt.Println("Жираф втік")
		return Giraffe{animal: animal{species: species, name: name, size: "Large", loose: true}}
	case "Camel":
		fmt.Println("Верблюд втік")
		return Camel{animal: animal{species: species, name: name, size: "Medium", loose: true}}
	case "Tiger":
		fmt.Println("Тигр втік")
		return Tiger{animal: animal{species: species, name: name, size: "Medium", loose: true}}
	case "Lion":
		fmt.Println("Лев втік")
		return Lion{animal: animal{species: species, name: name, size: "Medium", loose: true}}
	default:
		return Elephant{animal: animal{species: species, name: name, size: "Large", loose: true}}
	}
}

func CreateAnimalsFromMap(animalList map[string][]string) (animals []AnyAnimal) {
	for species, nameList := range animalList {
		for _, name := range nameList {
			animals = append(animals, AnimalConst(species, name))
		}

	}
	return animals

}
