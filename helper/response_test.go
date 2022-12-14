package helper

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResponse(t *testing.T) {
	data := map[string]interface{}{
		"test": "ok",
	}
	expected := Response{
		StatusCode: http.StatusOK,
		Message:    "ok",
		Data:       data,
	}

	resp := NewResponse(http.StatusOK, "ok", data)

	assert.Equal(t, expected, resp)
}

func TestNewErrorResponse(t *testing.T) {
	data := map[string]interface{}{
		"test": "failed",
	}
	expected := ErrorResponse{
		StatusCode: http.StatusBadRequest,
		Message:    "failed",
		Errors:     data,
	}

	resp := NewErrorResponse(http.StatusBadRequest, "failed", data)

	assert.Equal(t, expected, resp)
}
