package membership

import (
	"context"
	"simple-forum/internals/models/membership"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	*gin.Engine
	membershipService
}

type membershipService interface {
	Signup(c context.Context, req membership.SignUpRequest) error
}

func NewHandler(api *gin.Engine, membershipSvc membershipService) *Handler {
	return &Handler{
		Engine:            api,
		membershipService: membershipSvc,
	}
}

func (h *Handler) RegisterRoutes() {
	route := h.Group("membership")
	route.GET("/ping", h.Ping)
	route.POST("/signup", h.Signup)
}
