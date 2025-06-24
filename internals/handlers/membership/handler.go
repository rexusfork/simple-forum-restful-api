package membership

import (
	"simple-forum/internals/models/membership"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	*gin.Engine
	membershipService
}

type membershipService interface {
	Signup(c *gin.Context, req membership.SignUpRequest) error
}

func NewHandler(api *gin.Engine, membershipSvc membershipService) *Handler {
	return &Handler{
		Engine:            api,
		membershipService: membershipSvc,
	}
}

func (h *Handler) RegisterRoutes() {
	route := h.Group("membership")
	route.GET("/Ping", h.Ping)
}
