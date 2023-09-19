package textgame

type character struct {
	name           string
	areYouKing     bool
	hp             int
	hunger         int
	curentLocation location
}

func (c *character) takeDamage() {
	c.hp -= 1
}
func (c *character) eat() {
	c.hunger += 1
}

func (c *character) becameTheKing() {
	c.areYouKing = true
}

type action struct {
	discroption string
	choise      []string
}

type location struct {
	name    string
	actions []action
}
