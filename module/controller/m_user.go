package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"swagger-gin/module/dao"
	"swagger-gin/module/utils"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary get user
// @Schemes
// @Security token
// @Description get user
// @Tags m_user
// @Accept json
// @Produce json
// @Success 200 {array} utils.Respone
// @Router /m_user/get-user [get]
func UserGet(c *gin.Context) {
	var respone utils.Respone

	id_user, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("id user from token %d", id_user)

	result, message, data := dao.UserGet()

	fmt.Println(message)
	// fmt.Println(result)

	respone.Result = result
	respone.Message = message
	respone.Data = data
	c.JSON(http.StatusOK, respone)
}

// PingExample godoc
// @Summary get user by id
// @Schemes
// @Security token
// @Description get user by id
// @Tags m_user
// @Param id_user path int true "id_user"
// @Accept json
// @Produce json
// @Success 200 {array} utils.Respone
// @Router /m_user/get-user/{id_user} [get]
func UserGetById(c *gin.Context) {
	var respone utils.Respone
	id, err := strconv.Atoi(c.Param("id_user"))
	if err != nil {
		return
	}
	fmt.Println(id)

	result, message, data := dao.UserGetById(id)

	fmt.Println(message)
	// fmt.Println(result)

	respone.Result = result
	respone.Message = message
	respone.Data = data
	c.JSON(http.StatusOK, respone)
}

// PingExample godoc
// @Summary insert user
// @Schemes
// @Security token
// @Description insert user
// @Tags m_user
// @Param request body dao.M_User true "Payload Body [RAW]"
// @Accept json
// @Produce json
// @Success 200 {array} utils.Respone
// @Router /m_user/insert-user [post]
func User(c *gin.Context) {
	var respone utils.Respone
	var input dao.M_User

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := dao.M_User{}

	u.Username = input.Username
	u.Fullname = input.Fullname
	u.Password = input.Password
	// fmt.Println(input)
	result, message, id := dao.SaveUser(u)

	fmt.Println(message)
	// fmt.Println(result)

	respone.Result = result
	respone.Message = message
	respone.Data = id
	c.JSON(http.StatusOK, respone)
}

// PingExample godoc
// @Summary udate user
// @Schemes
// @Security token
// @Description update user
// @Tags m_user
// @Param request body dao.UserGetModel true "Payload Body [RAW]"
// @Accept json
// @Produce json
// @Success 200 {array} utils.Respone
// @Router /m_user/update-user [put]
func UserUpdate(c *gin.Context) {
	var respone utils.Respone
	var input dao.UserGetModel

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := dao.UserGetModel{}

	u.Username = input.Username
	u.Fullname = input.Fullname
	u.Id_user = input.Id_user
	// fmt.Println(input)
	result, message, id := dao.UpdateUser(u)

	fmt.Println(message)
	// fmt.Println(result)

	respone.Result = result
	respone.Message = message
	respone.Data = id
	c.JSON(http.StatusOK, respone)
}

// PingExample godoc
// @Summary delete user by id
// @Schemes
// @Security token
// @Description delete user by id
// @Tags m_user
// @Param id_user path int true "id_user"
// @Accept json
// @Produce json
// @Success 200 {array} utils.Respone
// @Router /m_user/delete-user/{id_user} [delete]
func UserDeleteById(c *gin.Context) {
	var respone utils.Respone
	id, err := strconv.Atoi(c.Param("id_user"))
	if err != nil {
		return
	}
	fmt.Println(id)

	result, message, data := dao.UserDeleteById(id)

	fmt.Println(message)
	// fmt.Println(result)

	respone.Result = result
	respone.Message = message
	respone.Data = data
	c.JSON(http.StatusOK, respone)
}
