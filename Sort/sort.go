package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Keys struct {
	caseIgnore bool
	unique bool
	reverse bool
	inFile bool
	numbers bool
	byColumn bool
	column int64
	fileName string
}
func sortReverse(elements []string) [] string{
	for i := len(elements)/2-1; i >= 0; i-- {
		opp := len(elements)-1-i
		elements[i], elements[opp] = elements[opp], elements[i]
	}
	return elements
}
func sortCase(elements[]string) []string{
	sort.Slice(elements, func(i, j int) bool { return strings.ToLower(elements[i]) < strings.ToLower(elements[j]) })
	return elements
}
func sortColumn(elements[]string, key int64)[]string{
	sort.Slice(elements, func(i, j int) bool {
		return strings.Fields(elements[i])[key-1]<strings.Fields(elements[j])[key-1]
	})
	return elements
}
func sortUnique(elements []string) []string {
	encountered := map[string]bool{}
	var result []string

	for v := range elements {
		if encountered[elements[v]] == true {
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}
func sortUniqueByColumn(elements []string,key int64) []string {
	encountered := map[string]bool{}
	var result []string

	for v := range elements {
		if encountered[strings.Fields(elements[v])[key-1]] == true {
		} else {
			encountered[strings.Fields(elements[v])[key-1]] = true
			result = append(result, elements[v])
		}
	}
	return result
}
func sortCaseUnique(elements []string) []string{
	encountered := map[string]bool{}
	encountered[elements[0]] = true

	for i:=1;i<len(elements);i++ {

		encountered[elements[i]] = true
		if strings.ToLower(elements[i-1])==strings.ToLower(elements[i])&&
			 				elements[i-1]!=elements[i]{
			encountered[elements[i]] = false
		}
	}

	var result []string
	for key, b := range encountered {
		if b{
			result = append(result, key)
		}
	}
	sort.Strings(result)
	return result
}
func sortCaseUniqueByColumn(elements []string, key int64) []string{
	encountered := map[string]bool{}
	encountered[elements[0]] = true
	for i:=1;i<len(elements);i++ {
		encountered[elements[i]] = true
		if strings.Fields(strings.ToLower(elements[i-1]))[key-1]==strings.Fields(strings.ToLower(elements[i]))[key-1]{
			encountered[elements[i]] = false
		}
	}

	var result []string
	for key, b := range encountered {
		if b{
			result = append(result, key)
		}
	}
	sortColumn(result,key)
	return result
}
func writeFile(elements[]string, file string){
	f, err := os.OpenFile(file,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	for _, line := range elements{
		if _, err := f.WriteString(line); err != nil {
			log.Println(err)
		}
		if _, err := f.WriteString("\n"); err != nil {
			log.Println(err)
		}
	}



}

func sortByKey(elements []string, keys Keys) []string {
	if keys.column != 0{
		if keys.caseIgnore{
			sort.Slice(elements, func(i, j int) bool {
				return  strings.ToLower( strings.Fields(elements[i])[keys.column-1]) <
					strings.ToLower(strings.Fields(elements[j])[keys.column-1])
			})
			if keys.unique{
				elements = sortCaseUniqueByColumn(elements,keys.column)
			}
		}else {
			elements = sortColumn(elements, keys.column)
			if keys.unique{
				elements = sortUniqueByColumn(elements,keys.column)
			}
		}
	}else {
		if keys.caseIgnore {
			elements = sortCase(elements)
			if keys.unique{
				elements = sortCaseUnique(elements)
			}
		}else {
			sort.Strings(elements)
			if keys.unique{
				elements = sortUnique(elements)
			}
		}
	}
	if keys.reverse{
		elements = sortReverse(elements)
	}
	if keys.inFile{
		writeFile(elements,keys.fileName)
	}

	return elements
}

func main() {

	args := os.Args[1:]
	content, _ := ioutil.ReadFile(args[len(args)-1])
	lines := strings.Split(string(content), "\n")

	var keys Keys
	for i := 0; i < len(args); i++{
		switch args[i] {

		case "-f":
			keys.caseIgnore = true

		case "-u":
			keys.unique = true

		case "-r":
			keys.reverse = true

		case "-o":
			keys.inFile = true
			keys.fileName = args[i+1]

		case "-n":
			keys.numbers = true

		case "-k":
			keys.byColumn = true
			keys.column, _ = strconv.ParseInt(args[i+1], 10, 64)
		}
	}

	lines = sortByKey(lines, keys)
	if !keys.inFile{
		fmt.Println(lines)
	}

	return
}

