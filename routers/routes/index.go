package routes

import (
	"fmt"
	"net/http"
)

// IndexHandler is a handler for index
type IndexHandler struct {
}

// NewIndexHandler create new index handler
func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

// GetRoute get route for index
func (h *IndexHandler) GetRoute() string {
	return "/"
}

// HandleRequest handle request for index
func (h *IndexHandler) HandleRequest(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "Hello")

	return nil
}
