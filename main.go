package main

import (
	"bufio"
	"firstProject/app/Http/Controller"
	"firstProject/app/Http/Middlewares"
	"firstProject/sproute"
	"net/http"
	"os"
	"strings"
)

func main() {

	err := os.Setenv("FOO", "1")
	check(err)

	writeEnv()

	route := sproute.Build()

	route.GET("/api/:word", func(request *http.Request, params sproute.H) sproute.Res {
		return sproute.ResponseString(200, params["word"])
	}).Middleware(middleware.FirstMiddleware{Params: sproute.H{"name":"ok"}}).Middleware(middleware.AnotherMiddleware{})

	route.GET("/api2/:word", Controller.IndexWeb1)

	route.Listen(":9900")
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