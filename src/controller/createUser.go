package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lucasbarroso23/baseline-GoAPI/src/configuration/rest_err"
	"github.com/lucasbarroso23/baseline-GoAPI/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		rest_err := rest_err.NewBadRequestError(
			fmt.Sprintf("there are some incorrect fields, error=%s\n", err.Error()),
		)

		c.JSON(rest_err.Code, rest_err)
		return
	}

	fmt.Println(userRequest)
}
