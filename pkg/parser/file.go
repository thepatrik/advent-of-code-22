package parser

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadIntFile reads file and returns the content as a int slice
func ReadIntFile(filename string) []int {
	var input = make([]int, 0)
	t, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	in := strings.Split(string(t), " ")
	for _, i := range in {
		i = strings.TrimSuffix(i, "\n")
		n, _ := strconv.ParseInt(i, 10, 32)
		input = append(input, int(n))
	}

	return input
}

// ReadFile reads file and returns the content as a string slice
func ReadFile(name string) []string {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	strslice := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		strslice = append(strslice, line)
	}
	return strslice
}
