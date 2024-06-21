package auxiliaries

import (
	"basic-go-rest-api/auxiliary_types"
	"encoding/json"
)

func GenerateResponseMessage(message string) (responseMessage []byte) {
	messageResponse := auxiliary_types.ResponseMessage{Message: message}
	responseMessage, _ = json.Marshal(messageResponse)

	return
}
