package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	app_err "github.com/nnnn24/go-gin-boilerplate/pkg/err"
)

func BindAndValidate(c *gin.Context, form interface{}) (int, int) {
	if err := c.ShouldBindJSON(form); err != nil {
		return http.StatusBadRequest, app_err.INVALID_PARAMS
	}

	return http.StatusOK, app_err.SUCCESS
}
