package serialization

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"users_service/models"
)

func DeserializeRequest(req *http.Request, result models.Validatable) (status int) {
	if req == nil {
		return http.StatusBadRequest
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return http.StatusBadRequest
	}
	if err := json.Unmarshal(body, result); err != nil {
		return http.StatusBadRequest
	}
	if !result.Validate() {
		return http.StatusBadRequest
	}
	return http.StatusOK
}

func SerializeResponse(writer http.ResponseWriter, response any) {
	marshalled, err := json.Marshal(response)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(marshalled)
}
