package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type JsonObj map[string]interface{}

func JsonResponse(w http.ResponseWriter, resMap JsonObj, code int) {
	jsonResponse, err := json.Marshal(resMap)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonResponse)
}

func GetIdParam(path string) (int64, error) {
	strId := path[len("/todos/"):]
	if strId == "" {
		return 0, errors.New("invalid id")
	}

	intId, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		return 0, errors.New("invalid id")
	}

	return intId, nil
}