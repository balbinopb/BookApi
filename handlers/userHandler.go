package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {

	// get request
	var body struct {
		email    string
		password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get body",
		})
		return
	}

	_,err:=bcrypt.GenerateFromPassword([]byte(body.password),12)

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to to hash password",
		})
		return
	}





}



func Login(c *gin.Context)  {
	var request LoginRequest;

	err:=c.ShouldBindJSON(&request)

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "input tidak valid",
		})
	}




	
}
