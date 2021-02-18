package server

import (
	"github.com/Perezonance/tinder-copycat/profile-management-service/internal/storage"
)

type (
	/*Server processes the data models and handles besiness logic for
	the server*/
	Server struct {
		db storage.Storage
	}
)

/*NewServer initializes a new server structure with the given configuration
and a database access utility*/
func NewServer(db storage.Storage, c Config) *Server {
	return &Server{db: db}

}
