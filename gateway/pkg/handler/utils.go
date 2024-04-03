package handler

import (
	"errors"

	"github.com/valyala/fasthttp"
)

func setErrorResponse(res *fasthttp.Response, statusCode int, message string) {
	res.SetStatusCode(statusCode)
	res.SetBody([]byte(message))
}

func handleClientError(res *fasthttp.Response, err error) {
	if errors.Is(err, fasthttp.ErrTimeout) {
		setErrorResponse(res, fasthttp.StatusRequestTimeout, err.Error())
	} else {
		setErrorResponse(res, fasthttp.StatusInternalServerError, err.Error())
	}
}
