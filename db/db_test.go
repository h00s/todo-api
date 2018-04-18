package db

import (
	"testing"

	"github.com/h00s/todo-api/config"
)

func TestDB(t *testing.T) {
	c, err := config.Load("../configuration_test.json")
	if err != nil {
		t.Error("Unable to load configuration")
	}

	_, err = Connect(c.Database)
	if err != nil {
		t.Error("Unable to connect to DB", err)
	}
}
