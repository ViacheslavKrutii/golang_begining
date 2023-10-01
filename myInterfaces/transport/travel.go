package mytransport

import (
	"fmt"
	"math/rand"
)

func Travel(trname string, c []City) {
	r := route{cities: c}
	t := traveler{name: trname}
	aviliableTrans := []transport{car{name: "car"}, train{name: "train"}, plane{name: "plane"}}
	i := 0

	fmt.Printf("%s start travel from %s.\n", t.name, r.cities[i])
	for i := range r.cities {
		curentTrans := aviliableTrans[rand.Intn(len(aviliableTrans))]
		r.addTransToRoute(curentTrans)
		switch {
		case i < (len(r.cities) - 1):
			fmt.Printf("%s travel from %s to %s by %v.\n", t.name, r.cities[i], r.cities[(i+1)], curentTrans.returnName())
		case i == (len(r.cities) - 1):
			fmt.Printf("%s arrived to %s by %s.\n", t.name, r.cities[i], curentTrans.returnName())
		}
	}

	r.showListOfTrans()

}
