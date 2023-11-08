package serialization

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

type deserializableStruct struct {
	Field string `json:"field"`
}

func (obj *deserializableStruct) Validate() bool {
	return obj.Field[0:4] == "test"
}

func TestDeserializationOK(t *testing.T) {
	obj := deserializableStruct{Field: "test"}
	marshalled, _ := json.Marshal(obj)
	req, _ := http.NewRequest("POST", "", bytes.NewReader(marshalled))
	var newObj deserializableStruct
	status := DeserializeRequest(req, &newObj)
	if status != 200 || newObj.Field != obj.Field {
		t.Fatal("Deserialization failed")
	}
}

func TestDeserializationFail(t *testing.T) {
	obj := deserializableStruct{Field: "badtest"}
	marshalled, _ := json.Marshal(obj)
	req, _ := http.NewRequest("POST", "", bytes.NewReader(marshalled))
	var newObj deserializableStruct
	status := DeserializeRequest(req, &newObj)
	if status == 200 {
		t.Fatal("Deserialization succeded, but had not to")
	}
}
