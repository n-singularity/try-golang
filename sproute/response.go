package sproute

import "encoding/json"

type Res struct {
	code      int
	typeValue string
	response  interface{}
}

func Response() Res {
	return Res{}
}

func (it *Res) SetCode(c int) {
	it.code = c
}

func (it *Res) SetType(t string) {
	it.typeValue = t
}

func (it *Res) SetResponse(r interface{}) {
	it.response = r
}

func (it Res) getResponse() string {
	if(it.typeValue == "json"){
		js, _ := json.Marshal(it.response)
		return string(js)
	}else if(it.typeValue == "string"){
		return it.response.(string)
	}

	return ""
}

func ResponseString(code int, response interface{}) Res  {
	res := Response()
	res.SetCode(code)
	res.SetType("string")
	res.SetResponse(response.(string))
	return res
}

func ResponseJson(code int, response interface{}) Res  {
	res := Response()
	res.SetCode(code)
	res.SetType("json")
	res.SetResponse(response)
	return res
}