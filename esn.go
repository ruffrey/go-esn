package main

import (
	"io/ioutil"
	"math/rand"
)

/*
Adapted from an Echo State Network in python:

    A minimalistic Echo State Networks demo with Mackey-Glass (delay 17) data
    in "plain" scientific Python.
    by Mantas LukoĹĄeviÄ?ius 2012
    http://minds.jacobs-university.de/mantas
*/
const (
	trainLen = 2000
	testLen  = 2000
	initLen  = 100
)

func main() {
	data, err := ioutil.ReadFile("MackeyGlass_t17.txt")
	if err != nil {
		panic(err)
	}

	// generate the ESN reservoir
	inSize := 1
	outSize := 1
	resSize := 1000
	a := 0.3 // leaking rate

	rand.Seed(42)
}
