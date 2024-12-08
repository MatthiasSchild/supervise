package errs_test

import (
	"errors"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/MatthiasSchild/supervise/errs"
)

func TestExtractMessages(t *testing.T) {
	baseError := errors.New("base")
	wrap1 := errs.Wrap(baseError, errs.Details{Message: "wrap1"})
	wrap2 := errs.Wrap(wrap1, errs.Details{Message: "wrap2"})

	c := &gin.Context{Errors: []*gin.Error{
		{Err: wrap2},
	}}
	extract := errs.ExtractMessages(c)

	if len(extract) != 3 {
		t.Errorf("Expected 3 errors, got %d", len(extract))
	}
	if extract[0] != "wrap2" {
		t.Errorf("Expected 'wrap2', got %s", extract[0])
	}
	if extract[1] != "wrap1" {
		t.Errorf("Expected 'wrap1', got %s", extract[1])
	}
	if extract[2] != "base" {
		t.Errorf("Expected 'base', got %s", extract[2])
	}
}
