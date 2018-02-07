package model

type gopher struct {
	name    string
	age     int
	isAdult bool
}

type horse struct {
	name   string
	weight int
}

type jumper interface {
	Jump() string
}

func (g gopher) Jump() string {
	if g.age < 65 {
		return g.name + " can jump HIGH"
	}
	return g.name + " can still jump"
}

func validateAge(g *gopher) {
	g.isAdult = g.age >= 21
}

func (h horse) Jump() string {
	if h.weight < 2500 {
		return "I'm too heavy, can't jump"
	}
	return "I will jump, neigh!!!"
}

func GetList() []jumper {
	phil := &gopher{name: "Phil", age: 30}
	noodles := &gopher{name: "Noodles", age: 30}
	barbaro := &horse{name: "Barbaro", weight: 2000}

	list := []jumper{phil, noodles, barbaro}
	return list
}
