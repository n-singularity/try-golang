package main

import (
	"bufio"
	"firstProject/routes"
	"firstProject/sproute"
	"os"
	"strings"
)

func main() {

	err := os.Setenv("FOO", "1")
	check(err)

	writeEnv()

	r := sproute.Build()

	r = routes.Route(r)

	r.Listen(":"+os.Getenv("PORT"))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeEnv() {
	file, err := os.Open(".env")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			s := strings.Split(scanner.Text(), "=")
			if len(s) == 2 {
				err = os.Setenv(s[0], s[1])
				check(err)
			}
		}

	}
}
