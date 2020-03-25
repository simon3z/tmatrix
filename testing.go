package main

import "strings"

// TestDefinition represents a test definition
type TestDefinition struct {
	Name  string `yaml:"Name"`
	Score int    `yaml:"Score"`
}

// TestDefinitionSet represents a set of test definitions
type TestDefinitionSet []TestDefinition

// TestPermutationItem represents a item of a test permutation
type TestPermutationItem struct {
	Name  string            `yaml:"Name"`
	Items TestDefinitionSet `yaml:"Items"`
}

// TestPermutationSet represents a set of item for a test permutation
type TestPermutationSet []TestPermutationItem

// ByScore can be used as a sorting strategy for a list of TestDefinitionSet
type ByScore []TestDefinitionSet

// Permutations returns all the permustations for a list of TestDefinitionSet
func (p TestPermutationSet) Permutations() []TestDefinitionSet {
	s := []TestDefinitionSet{}
	c := make([]int, len(p))

	// TODO: first attempt for permutation algorithm, feel free to improve
	for {
		o := make(TestDefinitionSet, len(p))

		for j := range p {
			o[j] = p[j].Items[c[j]]
		}

		s = append(s, o)

		complete := true

		for j := len(p) - 1; j >= 0; j-- {
			c[j]++

			if c[j] >= len(p[j].Items) {
				c[j] = 0
			} else {
				complete = false
				break
			}
		}

		if complete {
			break
		}
	}

	return s
}

func (p TestDefinitionSet) String() string {
	s := make([]string, len(p))

	for i := range p {
		s[i] = p[i].Name
	}

	return strings.Join(s, ", ")
}

// Score returns the combined total score of the tests definitions in the set
func (p TestDefinitionSet) Score() int {
	v := 0

	for i := range p {
		v += p[i].Score
	}

	return v
}

func (p ByScore) Len() int {
	return len(p)
}

func (p ByScore) Less(i, j int) bool {
	return p[i].Score() < p[j].Score()
}

func (p ByScore) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
