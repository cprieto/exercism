package robotname

import (
	"fmt"
	"math/rand"
)

type Robot struct {
	initialized bool
	name        string
}

func (r *Robot) Reset() {
	r.initialized = false
}

var letters = "ABCDEFGHIJKLMNOPQRSTUWXYZ"
var names = make(map[string]bool)

func generateName() string {
	var name string
	n1 := rand.Intn(len(letters))
	n2 := rand.Intn(len(letters))
	n3 := rand.Intn(1000)
	name = fmt.Sprintf("%s%s%03d", string(letters[n1]), string(letters[n2]), n3)

	return name
}

func (r *Robot) Name() (string, error) {
	if !r.initialized {
		r.name = generateName()
		r.initialized = true
		names[r.name] = true
	}

	return r.name, nil
}
