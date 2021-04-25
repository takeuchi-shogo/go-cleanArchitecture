package controllers

type H struct {
	Result string      `json:"result"`
	Data   interface{} `json:"data"`
}

func NewH(result string, data interface{}) *H {
	H := new(H)
	H.Result = result
	H.Data = data
	return H
}
