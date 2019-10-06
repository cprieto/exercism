package robotname

import (
	"math/rand"
	"fmt")

type Robot struct {
	initialized bool
	name        string
}

func (r *Robot) Reset() {
	r.initialized = false
}

var letters = "ABCDEFGHIJKLMNOPQRSTUWXYZ"
var names map[string]bool

func generateName() string {
	n1 := rand.Intn(len(letters))
	n2 := rand.Intn(len(letters))
	n3 := rand.Intn(1000)

	return fmt.Sprintf("%s%s%03d", string(letters[n1]), string(letters[n2]), n3)
}

func (r *Robot) Name() (string, error) {
	r.name = generateName()
	if !r.initialized {
		for _, ok := names[r.name]; ok == true {

		}
		r.initialized = true

		names = append(names, r.name)
	}

	return r.name, nil
}
