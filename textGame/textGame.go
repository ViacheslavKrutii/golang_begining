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
	hp             int
	hunger         int
	inventory      inventory
	curentLocation location
}

type location struct {
	name string
}
