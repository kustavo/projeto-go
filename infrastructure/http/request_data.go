package http

import (
	"encoding/json"
	"io/ioutil"
	nethttp "net/http"
	"path"
	"strconv"
	"strings"

	"github.com/kustavo/benchmark/go/domain/model"
)

func RequestId(w nethttp.ResponseWriter, r *nethttp.Request) (uint64, error) {
	idStr := path.Base(r.URL.Path)
	id, err := strconv.ParseUint(string(idStr), 10, 64)
	return id, err
}

func RequestBody(w nethttp.ResponseWriter, r *nethttp.Request) (*[]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return &body, err
}

func RequestModel(w nethttp.ResponseWriter, r *nethttp.Request, strc interface{}) error {
	body, err := RequestBody(w, r)
	if err != nil {
		return err
	}

	err = json.Unmarshal(*body, strc)
	return err
}

func RequestEntity(w nethttp.ResponseWriter, r *nethttp.Request, entity model.Entity) error {
	if err := RequestModel(w, r, entity); err != nil {
		return err
	}

	id, _ := RequestId(w, r)

	entity.SetId(id)
	return nil
}

func RequestToken(r *nethttp.Request) string {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
