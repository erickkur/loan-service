package handler

import (
	jsonEncoding "encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

var DefaultDecoder = schema.NewDecoder()

// EndpointHandler ...
type EndpointHandler func(http.ResponseWriter, *http.Request) ResponseInterface

func (fn EndpointHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	res := fn(w, r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.GetStatus())

	if res.HasError() {
		handleErrorResponse(w, r, res)
	} else {
		handleOKResponse(w, res)
	}
}

func handleOKResponse(w http.ResponseWriter, res ResponseInterface) {
	if res.HasNoContent() {
		return
	}

	resp := make(map[string]interface{})

	resp["message"] = "success"
	data := res.GetData()

	if data == nil {
		emptyData := make(map[string]string)
		resp["data"] = emptyData
	} else {
		resp["data"] = data
	}

	encodeResponse(w, resp)
}

func handleErrorResponse(w http.ResponseWriter, r *http.Request, res ResponseInterface) {
	data := map[string]interface{}{
		"data":    make(map[string]string),
		"code":    res.GetErrCode(),
		"message": res.GetErrorMessage(),
	}

	encodeResponse(w, data)
}

func encodeResponse(w http.ResponseWriter, data interface{}) {
	err := jsonEncoding.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("Error encode:", err.Error())
		http.Error(w, "Error encode response", http.StatusInternalServerError)
	}
}
