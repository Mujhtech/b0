package services

import (
	"testing"

	"github.com/mujhtech/b0/database/models"
)

func stripVariableFields(t *testing.T, obj string, v interface{}) {
	switch obj {
	case "project":
		g := v.(*models.Project)

		g.Slug = ""
	case "endpoint":
		e := v.(*models.Endpoint)

		e.Slug = ""
	default:
		t.Errorf("invalid data body - %v of type %T", obj, obj)
		t.FailNow()
	}
}
