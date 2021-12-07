package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func step1() {
	content, err := ioutil.ReadFile("01/01.input")
	if err != nil {
		log.Fatal(err)
	}
	numbers := strings.Split(string(content), "\n")
	result := 0
	for _, number := range numbers {
		converted, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}
		result += converted/3 - 2
	}
	log.Default().Println(result)
}

func step2() {
	content, err := ioutil.ReadFile("01/01.input")
	if err != nil {
		log.Fatal(err)
	}
	numbers := strings.Split(string(content), "\n")
	result := 0
	for _, number := range numbers {
		converted, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}
		for converted/3-2 > 0 {
			converted = converted/3 - 2
			result += converted
		}
	}
	log.Default().Println(result)
}

func main() {
	step1()
	step2()
}
