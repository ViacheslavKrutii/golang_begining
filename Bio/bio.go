package bio

import (
	"fmt"
)

type profession string

type bio struct {
	name       string
	surname    string
	alive      bool
	age        int16
	profession profession
	height     float32
}

func BioConst(name, surname string, alive bool, age int16, profession profession, height float32) bio {
	return bio{
		name:       name,
		surname:    surname,
		alive:      alive,
		age:        age,
		profession: profession,
		height:     height,
	}
}

func IsAlive(bio bio) string {
	status := "помер"
	switch bio.alive {
	case true:
		status = "живий"
	default:
		status = "помер"
	}
	return status
}

func Bio() {
	Borges := BioConst("Хорхе", "Борхес", false, 86, "письменник", 167.5)
	fmt.Printf("Ім'я %s, прізвище %s, %s, вік %d, рід діяльності %v, зріст %f.", Borges.name, Borges.surname, IsAlive(Borges), Borges.age, Borges.profession, Borges.height)
}
