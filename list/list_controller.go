package list

import (
	"github.com/gin-gonic/gin"
	"github.com/h00s/todo-api/db"
	"github.com/h00s/todo-api/logger"
)

// Controller for Link methods
type Controller struct {
	db     *db.Database
	logger *logger.Logger
}

type errorResponse struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}

// NewController creates new link controller
func NewController(db *db.Database, logger *logger.Logger) *Controller {
	return &Controller{db: db, logger: logger}
}

// GetLists return all lists
func (lc *Controller) GetLists(c *gin.Context) {
	c.String(200, "Hello world")
}
