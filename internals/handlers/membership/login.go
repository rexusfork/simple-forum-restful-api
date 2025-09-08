package membership

import (
	"net/http"
	"simple-forum/internals/models/membership"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	context := c.Request.Context()

	var request membership.LoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, err := h.membershipService.Login(context, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response:= membership.LoginResponse{
		AccessToken: accessToken,
	}

	c.JSON(http.StatusOK, response)
}