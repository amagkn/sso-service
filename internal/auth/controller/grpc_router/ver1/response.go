package ver1

import (
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type errorPayload struct {
	Type    error
	Details any
}

type errorJSONPayload struct {
	Type    string `json:"type"`
	Details any    `json:"details"`
}

type errorJSONBody struct {
	Error errorJSONPayload `json:"error"`
}

func errorResponse(statusCode codes.Code, payload errorPayload) error {
	jsonFields, _ := json.Marshal(errorJSONBody{
		errorJSONPayload{
			Type:    payload.Type.Error(),
			Details: payload.Details,
		},
	})

	return status.Error(statusCode, string(jsonFields))
}
