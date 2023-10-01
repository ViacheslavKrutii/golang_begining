package mytransport

// train
type train struct {
	name      string
	passenger traveler
}

func (t train) returnName() string {
	return t.name
}

func (t train) move() {

}
func (t train) stop() {

}
func (t train) changeSpeed() {

}

func (t train) takePassenger(traveler) {

}
func (t train) dropPassengers(traveler) {

}
