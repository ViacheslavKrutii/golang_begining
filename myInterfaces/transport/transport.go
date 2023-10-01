package mytransport

import "fmt"

type transport interface {
	takePassenger(traveler)
	dropPassengers(traveler)
	returnName() string
}
type traveler struct {
	name string
}
type City string

type route struct {
	cities      []City
	listOfTrans []transport
}

func (r *route) addTransToRoute(t transport) {
	r.listOfTrans = append(r.listOfTrans, t)
}
func (r *route) showListOfTrans() {
	fmt.Println(r.listOfTrans)
}
