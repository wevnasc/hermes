package email

import (
	"fmt"
	"net/http"

	"github.com/wevnasc/hermes/server"
)

type controller interface {
	schedule(s *schedule) *server.ServerError
}

type EmailController struct {
	repo    repo
	storage storage
}

func NewEmailController(repo repo, storage storage) *EmailController {
	return &EmailController{repo, storage}
}

func (c *EmailController) schedule(s *schedule) *server.ServerError {
	if !c.storage.exists(s.template) {
		err := fmt.Sprintf("template not exits %s", s.template)
		return server.NewError(err, http.StatusBadRequest)
	}
	c.repo.save(s)
	return nil
}
