package users

import (
	"github.com/gin-gonic/gin"
	"github.com/marcelo-cardozo/bookstore_users-api/domain/users"
	"github.com/marcelo-cardozo/bookstore_users-api/service"
	"github.com/marcelo-cardozo/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User

	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, err)
	//	return
	//}
	//
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//	c.JSON(http.StatusBadRequest, err)
	//	return
	//}

	// or
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestRestErr(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}

	createdUser, restErr := service.UsersService.CreateUser(user)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestRestErr("used id should be a number")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, restErr := service.UsersService.GetUser(userId)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, user)
}
