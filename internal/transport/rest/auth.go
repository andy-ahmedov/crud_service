package rest

import (
	"net/http"

	"github.com/andy-ahmedov/crud_service/internal/domain"
	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags auth
// @Description Registration of a new user in the system.
// @ID sign-up
// @Accept json
// @Produce json
// @Param input body domain.User true "User info"
// @Success 200 {string} gin.H "The user has been successfully registered."
// @Failure 400 {object} errResponse "Bad Request"
// @Failure 500 {object} errResponse "Internal Server Error"
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var user domain.SignUpInput

	if err := c.ShouldBindJSON(&user); err != nil {
		logError("signUp", "incorrect format", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userService.SignUp(c.Request.Context(), user); err != nil {
		logError("signUp", "Internal Service Error", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
