package controllers

import (
	"net/http"

	"github.com/Perezonance/hello-world/util/log"
	"github.com/Perezonance/tinder-copycat/profile-management-service/internal/server"
)

type (
	/*Controller handles and processes requests and responses to
	the server as well as input validation.*/
	Controller struct {
		s *server.Server
	}
)

/*NewController creates a controller for handling and processing requests and
responses to the server*/
func NewController(s *server.Server, c *Config) *Controller {
	return &Controller{s: s}
}

/*PostProfilesHandler processes request by calling the appropriate server
operation and returns the result via the response*/
func (c *Controller) PostProfilesHandler(w http.ResponseWriter, r *http.Request) {
	writeRes(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented), w)
	return
}

/*GetProfilesHandler processes request by calling the appropriate server
operation and returns the result via the response*/
func (c *Controller) GetProfilesHandler(w http.ResponseWriter, r *http.Request) {
	writeRes(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented), w)
	return
}

func writeRes(statusCode int, message string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	res := []byte(message)
	_, err := w.Write(res)
	if err != nil {
		log.ErrorLog("Error while writing to ResponseWriter", err)
		return
	}
	return
}
