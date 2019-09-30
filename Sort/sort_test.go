package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_sortReverse(t *testing.T) {
	arr := []string{"a", "c", "b", "f"}
	exp := []string{"f", "b", "c", "a"}
	res := sortReverse(arr)
	assert.Equal(t, exp, res, "error message %s", "formatted")
	fmt.Println("Test_sortReverse -- OK")
}
func Test_sortCase(t *testing.T) {
	arr := []string{"Napkin", "Book", "BOOK", "Apple"}
	exp := []string{"Apple", "Book", "BOOK", "Napkin"}
	res := sortCase(arr)
	assert.Equal(t, exp, res, "error message %s", "formatted")
	fmt.Println("Test_sortCase -- OK")
}
func Test_sortColumn(t *testing.T) {
	arr := []string{"Hauptbahnhof 5", "Book 3", "Go 4"}
	exp := []string{"Book 3", "Go 4", "Hauptbahnhof 5"}
	res := sortColumn(arr, 2)
	assert.Equal(t, exp, res, "error message %s", "formatted")
	fmt.Println("Test_sortColumn -- OK")
}
func Test_sortUnique(t *testing.T) {
	arr := []string{"Apple", "BOOK", "Book", "Go", "January", "January", "Napkin"}
	exp := []string{"Apple", "BOOK", "Book", "Go", "January", "Napkin"}
	res := sortUnique(arr)
	assert.Equal(t, exp, res, "error message %s", "formatted")
	fmt.Println("Test_sortUnique -- OK")
}
func Test_sortUniqueByColumn(t *testing.T) {
	arr := []string{"6 Apple", "7 BOOK", "5 Book", "8 Go", "3 Hauptbahnhof", "1 January", "2 January", "4 Napkin"}
	exp := []string{"6 Apple", "7 BOOK", "5 Book", "8 Go", "3 Hauptbahnhof", "1 January", "4 Napkin"}
	res := sortUniqueByColumn(arr, 2)
	assert.Equal(t, exp, res, "error message %s", "formatted")
	fmt.Println("Test_sortUniqueByColumn -- OK")
}
func Test_sortCaseUnique(t *testing.T) {
	arr := []string{"Apple", "BOOK", "Book", "Go", "January", "January", "Napkin"}
	exp := []string{"Apple", "BOOK", "Go", "January", "Napkin"}
	res := sortCaseUnique(arr)
	assert.Equal(t, exp, res, "error message %s", "formatted")
	fmt.Println("Test_sortCaseUnique -- OK")
}
func Test_sortCaseUniqueByColumn(t *testing.T) {
	arr := []string{"6 Apple", "7 BOOK", "5 Book", "8 Go", "3 Hauptbahnhof", "1 January", "2 January", "4 Napkin"}
	exp := []string{"6 Apple", "7 BOOK", "8 Go", "3 Hauptbahnhof", "1 January", "4 Napkin"}
	res := sortCaseUniqueByColumn(arr, 2)
	assert.Equal(t, exp, res, "error message %s", "formatted")
	fmt.Println("Test_sortCaseUniqueByColumn -- OK")
}
func Test_writeFile(t *testing.T) {
	arr := []string{"Apple", "BOOK", "Book", "Go", "January", "January", "Napkin"}
	file := "file.txt"
	writeFile(arr, file)
	content, err := ioutil.ReadFile(file)
	res := strings.Split(string(content), "\n")
	res = res[:len(res)-1]
	assert.Equal(t, arr, res, "error message %s", "formatted")
	assert.Equal(t, err, nil, "error message %s", "formatted")
	fmt.Println("Test_writeFile -- OK")
}

func Test_sortBuKey1(t *testing.T) {
	keys := Keys{false, false, false, false, 0, "data.txt"}
	arr := []string{"Apple", "BOOK", "January", "Book", "Napkin", "Go", "January"}
	exp := []string{"Apple", "BOOK", "Book", "Go", "January", "January", "Napkin"}
	res := sortByKey(arr, keys)
	assert.Equal(t, exp, res, "error message %s", "formatted")
	fmt.Println("Test_sortBuKey1 -- OK")
}
func Test_sortBuKey2(t *testing.T) {
	keys := Keys{true, false, false, true, 1, "data.txt"}
	arr := []string{"6 Hauptbahnhof", "3 Book", "4 Go", "5 BOOK"}
	exp := []string{"3 Book", "4 Go", "5 BOOK", "6 Hauptbahnhof"}
	res := sortByKey(arr, keys)
	assert.Equal(t, exp, res, "error message %s", "formatted")
	fmt.Println("Test_sortBuKey2 -- OK")
}
func Test_sortBuKey3(t *testing.T) {
	keys := Keys{true, true, false, false, 0, "data.txt"}
	arr := []string{"Napkin", "Apple", "January", "BOOK", "January", "Hauptbahnhof", "Book", "Go"}
	exp := []string{"Apple", "Book", "Go", "Hauptbahnhof", "January", "Napkin"}
	res := sortByKey(arr, keys)
	assert.Equal(t, exp, res, "error message %s", "formatted")
	fmt.Println("Test_sortBuKey3 -- OK")
}
func Test_sortBuKey4(t *testing.T) {
	keys := Keys{false, true, true, false, 0, "data.txt"}
	arr := []string{"Apple", "BOOK", "January", "Book", "Napkin", "Go", "January"}
	exp := []string{"Napkin", "January", "Go", "Book", "BOOK", "Apple"}
	res := sortByKey(arr, keys)
	assert.Equal(t, exp, res, "error message %s", "formatted")
	fmt.Println("Test_sortBuKey4 -- OK")
}
func Test_sortBuKey5(t *testing.T) {
	keys := Keys{true, true, true, true, 2, "data.txt"}
	arr := []string{"4 Napkin", "6 Apple", "1 January", "7 BOOK", "2 January", "3 Hauptbahnhof", "5 Book", "8 Go"}
	exp := []string{"4 Napkin", "1 January", "3 Hauptbahnhof", "8 Go", "5 Book", "6 Apple"}
	res := sortByKey(arr, keys)
	assert.Equal(t, exp, res, "error message %s", "formatted")
	fmt.Println("Test_sortBuKey5 -- OK")
}
