package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseInfix1(t *testing.T) {
	str := []string{"(", "8", "-", "1", ")", "*", "5"}
	res := parseInfix(str)
	exp := "8 1 - 5 *"
	assert.Equal(t, exp, res, "error message %s", "formatted")
}
func Test_parseInfix2(t *testing.T) {
	str := []string{"8", "-", "1", "*", "5"}
	res := parseInfix(str)
	exp := "8 1 5 * -"
	assert.Equal(t, exp, res, "error message %s", "formatted")
}

func Test_calc1(t *testing.T) {
	str := []string{"8", "1", "-", "5", "*"}
	res, err := calc(str)
	exp := "35"
	assert.Equal(t, exp, res, "error message %s", "formatted")
	assert.Equal(t, err, nil, "error message %s", "formatted")
}
func Test_calc2(t *testing.T) {
	str := []string{"8", "1", "5", "*", "-"}
	res, err := calc(str)
	exp := "3"
	assert.Equal(t, exp, res, "error message %s", "formatted")
	assert.Equal(t, err, nil, "error message %s", "formatted")
}
