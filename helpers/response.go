package helpers

import (
	"encoding/json"
	"net/http"
)

type Respond struct {
	Status      int         `json:"status"`
	IsError     bool        `json:"isError"`
	Description string      `json:"desc,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

func (res *Respond) Send(w http.ResponseWriter) {
	err := json.NewEncoder(w).Encode(&Respond{
		Status:      res.Status,
		IsError:     res.IsError,
		Description: getStatus(res.Status),
		Data:        res.Data,
	})
	if err != nil {
		w.Write([]byte("Error When Encode response"))
	}
}

func Response(w http.ResponseWriter, data interface{}, code int, isError bool) {

	if isError {
		json.NewEncoder(w).Encode(&Respond{
			Status:      code,
			IsError:     isError,
			Description: getStatus(code),
			Data:        data,
		})
		return
	} else {
		json.NewEncoder(w).Encode(&Respond{
			Status:      code,
			IsError:     isError,
			Description: getStatus(code),
			Data:        data,
		})
	}
}

func getStatus(status int) string {
	var desc string

	switch status {
	case 200:
		desc = "OK"
	case 201:
		desc = "Created"
	case 400:
		desc = "Bad Request"
	case 401:
		desc = "Unauthorized"
	case 500:
		desc = "Internal Server Error"
	case 501:
		desc = "Bad Gateway"
	case 304:
		desc = "Not Modified"
	default:
		desc = ""
	}

	return desc
}
