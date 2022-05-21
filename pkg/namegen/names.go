package namegen

import "math/rand"

var mascNames = []string{
	"Edeard",
	"Jim",
	"Bob",
	"Bill",
	"Hank",
	"Frank",
	"Jerome",
	"Silens",
	"Harold",
	"Wayne",
}

var femmeNames = []string{
	"Violet",
	"Samantha",
	"Lucy",
	"Margarette",
	"Louise",
	"Ella",
	"Jeanette",
	"Eva",
	"Dolly",
	"Sue",
}

func RandomMascName() string {
	max := len(mascNames)
	i := rand.Intn(max)
	return mascNames[i]
}

func RandomFemmeName() string {
	max := len(femmeNames)
	i := rand.Intn(max)
	return femmeNames[i]
}
