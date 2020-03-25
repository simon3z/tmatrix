package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"sort"

	"gopkg.in/yaml.v2"
)

var commandFlags = struct {
	MatrixFile string
}{}

// ErrNoMatrixFile is the error returned when no matrix file was provided
var ErrNoMatrixFile = errors.New("No matrix file specified")

func init() {
	flag.StringVar(&commandFlags.MatrixFile, "matrix", "", "path to matrix file")
}

func main() {
	var permutations TestPermutationSet

	flag.Parse()

	if commandFlags.MatrixFile == "" {
		panic(ErrNoMatrixFile)
	}

	matrix, err := ioutil.ReadFile(commandFlags.MatrixFile)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(matrix, &permutations)

	if err != nil {
		panic(err)
	}

	s := permutations.Permutations()

	sort.Sort(sort.Reverse(ByScore(s)))

	for _, i := range s {
		fmt.Printf("%s (Score: %d)\n", i, i.Score())
	}
}
