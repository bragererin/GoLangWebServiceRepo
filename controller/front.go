package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	cc := newComputeController()

	http.Handle("/compute", *cc)
	http.Handle("/compute/", *cc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
