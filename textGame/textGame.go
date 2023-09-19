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

func (c *character) heal() {
	c.hp -= 1
}

func (c *character) eat() {
	c.hunger += 1
}

func (c *character) becameTheKing() {
	c.areYouKing = true
}

type event struct {
	discroption string
	action      string
	nextEvent   *event
}

type location struct {
	name   string
	events []event
}

const takeDamage = "takeDamage"
const heal = "heal"
const eat = "eat"
const becameTheKing = "becameTheKing"
