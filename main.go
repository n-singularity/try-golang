package main

import (
	"bufio"
	"firstProject/routes"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func main() {

	err := os.Setenv("FOO", "1")
	check(err)

	writeEnv()

	g := gin.Default()
	g.LoadHTMLGlob("resources/templates/**/*")

	g = routes.Route(g)

	err = g.Run() // listen and serve on 0.0.0.0:8080

	check(err)
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeEnv()  {
	file, err := os.Open(".env")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text()!=""{
			s := strings.Split(scanner.Text(), "=")
			if len(s)==2{
				err = os.Setenv(s[0], s[1])
				check(err)
			}
		}

	}
}