package auth

import (
	"fmt"
	"net/http"
	"swagger-gin/module/config"
	"swagger-gin/module/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// PingExample godoc
// @Summary Login
// @Schemes
// @Description Login
// @Tags login
// @Param request body loginRequest true "Payload Body [RAW]"
// @Accept json
// @Produce json
// @Success 200 {array} utils.Respone
// @Router /login/login [post]
func Login(g *gin.Context) {
	var requestBody loginRequest

	if err := g.BindJSON(&requestBody); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(requestBody.Username)

	db, err_conn := config.ConnectDataBase()
	if err_conn != nil {
		fmt.Println(err_conn.Error())
	}
	defer db.Close()

	// var result = User{}
	var result = User{}

	var respone utils.Respone
	err := db.
		QueryRow("select * from m_user WHERE username=$1", requestBody.Username).
		Scan(&result.Id_user, &result.Username, &result.Fullname, &result.Password)
	if err != nil {
		fmt.Println(err.Error())
		respone.Result = false
		respone.Message = "User not found"
		respone.Data = ""
		g.JSON(http.StatusOK, respone)
		return
	}

	fmt.Println(result)
	fmt.Println(requestBody.Password)

	err_verifiy := VerifyPassword(requestBody.Password, result.Password)

	if err_verifiy != nil && err_verifiy == bcrypt.ErrMismatchedHashAndPassword {
		respone.Result = false
		respone.Message = "Password tidak ditemukan"
		respone.Data = ""
		g.JSON(http.StatusOK, respone)
		return
	}

	token, err := utils.GenerateToken(uint(result.Id_user))

	if err != nil {
		respone.Result = false
		respone.Message = "Token gagal"
		respone.Data = ""
		g.JSON(http.StatusOK, respone)
		return
	}

	respone.Result = true
	respone.Message = "Login Berhasil"
	respone.Data = token
	g.JSON(http.StatusOK, respone)
	return

}

type User struct {
	Id_user  int    `json:"id_user"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
