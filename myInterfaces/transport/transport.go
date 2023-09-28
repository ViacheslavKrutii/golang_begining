package mytransport

type transport interface {
	takePassenger()
	dropPassengers()
}
type traveler struct {
	name string
}
type city string

type route struct {
	cities      []city
	listOfTrans []transport
}

func (r *route) addTransToRoute(t transport) {
	r.listOfTrans = append(r.listOfTrans, t)
}
func (r *route) showListOfTrans() {
	println(r.listOfTrans)
}
