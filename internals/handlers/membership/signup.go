package membership

import (
	"net/http"

	"simple-forum/internals/models/membership"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Signup(c *gin.Context) {
	context := c.Request.Context()

	var request membership.SignUpRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.membershipService.Signup(context, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}
