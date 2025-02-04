package robotname

import (
	"fmt"
	"math/rand"
)

const (
	ceil        = 10
	lenAlpha    = 26
	alphaOffset = int('A')
	template    = "%s%s%d%d%d"
)

// Define the Robot type here.
type Robot struct {
	name string
}

var names = map[string]bool{}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = genName()
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = genName()
}

func randLetter() string {
	n := rand.Intn(lenAlpha) + alphaOffset
	return string(rune(n))
}

func genName() string {
	name := createName()
	for names[name] {
		name = createName()
	}
	names[name] = true
	return name
}

func createName() string {
	return fmt.Sprintf(
		template,
		randLetter(),
		randLetter(),
		rand.Intn(ceil),
		rand.Intn(ceil),
		rand.Intn(ceil),
	)
}
