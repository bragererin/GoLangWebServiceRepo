package controllers

import (
	"Itential/AverageComputation/GoLangWebService/models"
	"encoding/json"
	"net/http"
	"regexp"
)

type computeController struct {
	pattern *regexp.Regexp
}

func (cc computeController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/compute" {
		switch r.Method {
		case http.MethodPost:
			cc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := cc.pattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		switch r.Method {
		case http.MethodPost:
			cc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (cc *computeController) post(w http.ResponseWriter, r *http.Request) {
	c, err := cc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse object"))
		return
	}
	c, err = models.AddAssignment(c)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(c, w)
}

func (cc *computeController) parseRequest(r *http.Request) (models.Assignment, error) {
	dec := json.NewDecoder(r.Body)
	var a models.Assignment
	err := dec.Decode(&a)
	if err != nil {
		return models.Assignment{}, err
	}
	return a, nil
}

func newComputeController() *computeController {
	return &computeController{
		pattern: regexp.MustCompile(`^/compute/(\d+)/?`),
	}
}
