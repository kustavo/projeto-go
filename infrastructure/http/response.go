package http

import (
	"encoding/json"
	nethttp "net/http"

	"github.com/kustavo/benchmark/go/domain"
	"github.com/kustavo/benchmark/go/domain/message"
)

// swagger:model responseMessage
type responseMessage struct {
	Errors   []string    `json:"errors"`
	Messages []string    `json:"messages"`
	Data     interface{} `json:"data"`
}

func newResponseMessage(err error, messages []string, data interface{}) responseMessage {
	responseMessage := responseMessage{
		Errors:   nil,
		Data:     data,
		Messages: messages,
	}

	if err == nil {
		return responseMessage
	}

	ce, ok := err.(*domain.ErrorsList)
	if ok {
		if ce == nil {
			return responseMessage
		}

		var errors []string
		for _, er := range ce.Errs {
			errors = append(errors, er.Error())
		}
		responseMessage.Errors = errors
	} else {
		responseMessage.Errors = []string{err.Error()}
	}

	if responseMessage.Messages == nil {
		responseMessage.Messages = responseMessage.Errors
	}

	return responseMessage
}

func Response(w nethttp.ResponseWriter, er error, msg []string, data interface{}) {
	if er != nil {
		ResponseBadRequest(w, er, msg)
	} else {
		ResponseSuccess(w, msg, data)
	}
}

func ResponseSuccess(w nethttp.ResponseWriter, msg []string, data interface{}) {
	responseSender(w, nil, msg, data, nethttp.StatusOK)
}

func ResponseNotFound(w nethttp.ResponseWriter, msg []string) {
	responseSender(w, message.ErrPageNotFound, msg, nil, nethttp.StatusNotFound)
}

func ResponseForbidden(w nethttp.ResponseWriter, er error, msg []string) {
	responseSender(w, er, msg, nil, nethttp.StatusForbidden)
}

func ResponseBadRequest(w nethttp.ResponseWriter, er error, msg []string) {
	responseSender(w, er, msg, nil, nethttp.StatusBadRequest)
}

func ResponseUnauthenticated(w nethttp.ResponseWriter, er error, msg []string) {
	responseSender(w, er, msg, nil, nethttp.StatusUnauthorized)
}

func ResponseServerError(w nethttp.ResponseWriter, er error, msg []string) {
	responseSender(w, er, msg, nil, nethttp.StatusInternalServerError)
}

func responseSender(w nethttp.ResponseWriter, er error, msg []string, data interface{}, code int) {
	m := newResponseMessage(er, msg, data)

	res, err := json.Marshal(&m)
	if err != nil {
		w.WriteHeader(nethttp.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseMessage{Errors: []string{message.ErrCreatingMessage.Error()}, Data: nil})
	} else {
		w.WriteHeader(code)
		w.Write(res)
	}
}
