package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbarroso23/baseline-GoAPI/src/configuration/validation"
	"github.com/lucasbarroso23/baseline-GoAPI/src/controller/model/request"
	"github.com/lucasbarroso23/baseline-GoAPI/src/controller/model/response"
)

func CreateUser(c *gin.Context) {

	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	c.JSON(http.StatusOK, response)
}
