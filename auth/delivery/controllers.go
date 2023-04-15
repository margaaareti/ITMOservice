package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"itmo/auth"
	"itmo/models"
	"net/http"
)

type Handler struct {
	auth auth.AuthService
}

func CreateNewHandler(auth auth.AuthService) *Handler {
	return &Handler{auth: auth}
}

func (h *Handler) SignUp(c *gin.Context) {

	if c.Request.Method != "POST" {
		newErrorResponse(c, http.StatusMethodNotAllowed, "ForbiddenMethod")
		return
	}

	var input models.User

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Info("Ошибка в привязке данных пользователя к объекту")
	}

	id, err := h.auth.SignUp(c.Request.Context(), input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Создан новый пользователь!": id,
	})
}
