package mytransport

// plane
type plane struct {
	name      string
	passenger traveler
}

func (p plane) returnName() string {
	return p.name
}

func (p plane) move() {

}
func (p plane) stop() {

}
func (p plane) changeSpeed() {

}
func (p plane) takePassenger(traveler) {

}
func (p plane) dropPassengers(traveler) {

}
