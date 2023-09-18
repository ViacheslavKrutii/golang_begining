package textgame

type weapon struct{
	string
}

type tool struct{
	string
}

type food struct{
	string
}

type inventory struct{
	weapons [] weapon
	tools []tool
	food [] food
}

type palyer struct {
	name string
	hp int
	hunger int
	inventory struct
}

