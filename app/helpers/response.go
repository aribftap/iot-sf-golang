package helpers

type Response struct {
	Status   int         `json:"status"`
	Messages string      `json:"message"`
	Data     interface{} `json:"data"`
	Error    error       `json:"-"`
}

type ValidationResponse struct {
	Status   int         `json:"status"`
	Messages string      `json:"message"`
	Errors   interface{} `json:"error"`
}
