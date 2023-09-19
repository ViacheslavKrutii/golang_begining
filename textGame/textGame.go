package textgame

type weapon struct {
	string
}

type tool struct {
	string
}

type food struct {
	string
}

type inventory struct {
	weapons []weapon
	tools   []tool
	food    []food
}

type character struct {
	name           string
	areYouKing     bool
	hp             int
	hunger         int
	inventory      inventory
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
