package Controller

import (
	"bufio"
	"firstProject/sproute"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Controller struct{}

func check(e error) (bool, sproute.Res) {
	if e != nil {
		return true, sproute.ResponseString(500, e.Error())
	}

	return false, sproute.ResponseString(200, "")
}

func ClassController() Controller {
	var controller Controller
	return controller
}

func (it Controller) IndexWeb1(request *http.Request, params sproute.H) sproute.Res {

	return sproute.ResponseString(200, params.Get("word"))
}

func (it Controller) CurlGet(request *http.Request, params sproute.H) sproute.Res {
	uri := "http://localhost:9000/api/v1/add-line-text"

	req, _ := http.NewRequest("POST", uri, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return sproute.ResponseString(200, string(body))
}

func (it Controller) CurlPost(request *http.Request, params sproute.H) sproute.Res {
	uri := "http://localhost:9000/api/v1/add-line-text"

	form := url.Values{}
	form.Add("text", request.FormValue("string"))

	req, _ := http.NewRequest("POST", uri, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)

	return sproute.ResponseString(200, string(body))
}

func (it Controller) AddValueInFileText(request *http.Request, params sproute.H) sproute.Res {
	file, err := os.OpenFile("tempfile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	errSts, _ := check(err)

	if errSts {
		file, err = os.Create("tempfile.txt")
	}

	defer file.Close()

	scanner := bufio.NewWriter(file)

	_, err = scanner.WriteString(request.FormValue("text") + "\n")

	check(err)

	defer scanner.Flush()

	return sproute.ResponseJson(200, sproute.H{
		"status": "success",
	})
}
