package controller

import (
	"net/http"
	"encoding/json"
)

func RegisterControllers() {
	uc := newUserController()
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface, w io.Writer) {
	enc := json.NewEncode(w)
	enc.Encode(data)
}
