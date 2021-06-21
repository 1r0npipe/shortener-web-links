package handler

import (
	"encoding/json"
)

func JsonVerify(data []byte) bool {
	return json.Valid(data)
}
