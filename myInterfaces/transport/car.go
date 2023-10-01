package mytransport

// car
type car struct {
	name      string
	passenger traveler
}

func (c car) returnName() string {
	return c.name
}

func (c car) move() {

}
func (c car) stop() {

}
func (c car) changeSpeed() {

}
func (c car) takePassenger(traveler) {

}
func (c car) dropPassengers(traveler) {

}
