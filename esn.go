package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"

	"github.com/gonum/matrix/mat64"
)

/*
Adapted from an Echo State Network in python:

http://minds.jacobs-university.de/sites/default/files/uploads/mantas/code/minimalESN.py.txt

    A minimalistic Echo State Networks demo with Mackey-Glass (delay 17) data
    in "plain" scientific Python.
    by Mantas Lukoševičius 2012
    http://minds.jacobs-university.de/mantas
*/
const (
	trainLen = 2000
	testLen  = 2000
	initLen  = 100
)

func main() {
	data := loadtxt("MackeyGlass_t17.txt")
	fmt.Println(data)

	// generate the ESN reservoir
	inSize := 1
	outSize := 1
	resSize := 1000
	a := 0.3 // leaking rate

	rand.Seed(42)

	Win := randMatrix(resSize, 1+inSize)
	Win.Sub(Win, copyAndFill(Win, 0.5))
	Win.MulElem(Win, copyAndFill(Win, 1)) // necessary to multiply by 1?

	W := randMatrix(resSize, resSize)
	W.Sub(W, copyAndFill(W, 0.5))

	// Option 1 - direct scaling (quick&dirty, reservoir-specific):
	W.MulElem(W, copyAndFill(W, 0.135))
	// Option 2 - normalizing and setting spectral radius (correct, slow):
	// fmt.Println("Computing spectral radius...")
	// lgeig := linalg.eig(W)[0]
	// rhoW := math.Max(math.Abs(lgeig))
	// fmt.Println("done")
	// W.Mul(W, copyAndFill(W, 1.25/rhoW))

	// allocated memory for the design (collected states) matrix
	X := zeros(1+inSize+resSize, trainLen-initLen)
	// set the corresponding target matrix directly
	Yt := data[initLen+1 : trainLen+1]

	// run the reservoir with the data and collect X
	x = zeros(resSize, 1)
	for t := 0; i < trainLen; i++ {
	    u := data[t]
	    x = (1-a)*x + a*math.Tanh( dot( Win, vstack((1,u)) ) + dot( W, x ) )
	    if t >= initLen {
				newX := mat64.NewDense(r, c, data)
				X[:,t-initLen] = vstack((1,u,x))[:,0]
			}
	}
}

// loadtxt behaves like numpy.loadtxt
func loadtxt(filename string) []string {
	dataBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(dataBytes), "\n")
}

// randMatrix behaves like numpy.random.rand
func randMatrix(r, c int) *mat64.Dense {
	data := make([]float64, r*c)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Float64()
	}
	mat := mat64.NewDense(r, c, data)
	return mat
}

func zeros(r, c int) *mat64.Dense {
	data := make([]float64, r*c)
	for i := 0; i < len(data); i++ {
		data[i] = 0
	}
	mat := mat64.NewDense(r, c, data)
	return mat
}

func copyAndFill(dense *mat64.Dense, value float64) *mat64.Dense {
	r, c := dense.Dims()
	data := make([]float64, r*c)
	for i := 0; i < len(data); i++ {
		data[i] = value
	}
	mat := mat64.NewDense(r, c, data)
	return mat
}
