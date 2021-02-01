package main

import (
	"github.com/nod/cowyo/config"
)

func main() {
	c, err := config.ParseFile("multisite_sample.toml")
	if err != nil {
		panic(err)
	}

	panic(c.ListenAndServe())
}
