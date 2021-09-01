package schedule

import (
	"fmt"
	"net/http"

	"github.com/wevnasc/hermes/server"
	"github.com/wevnasc/hermes/storage"
)

type controller struct {
	db      *memoDB
	storage storage.Storager
}

func newController(db *memoDB, storage storage.Storager) *controller {
	return &controller{db, storage}
}

func (c *controller) schedule(s *schedule) *server.ServerError {
	if !c.storage.Exists(s.template) {
		err := fmt.Sprintf("template not exits %s", s.template)
		return server.NewError(err, http.StatusBadRequest)
	}
	c.db.save(s)
	return nil
}
