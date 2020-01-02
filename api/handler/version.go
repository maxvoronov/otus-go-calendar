package handler

import (
	"github.com/maxvoronov/otus-go-calendar/internal/version"
	"net/http"
)

// VersionHandler Return current application version
// GET /
func (h *Handler) VersionHandler(_ *http.Request) APIResponse {
	return h.JSON(http.StatusOK, map[string]string{"version": version.Version})
}
