package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/bernardn38/financefirst/db"
	"github.com/bernardn38/financefirst/models"
	"github.com/bernardn38/financefirst/token"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	authorizationHeaderKey  = "auhtorization"
	authorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

type userResponse struct {
	Username string `json:"username"`
	UserId   int    `json:"user_id"`
}
type loginUserRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword checks if the provided password is correct or not
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginUser(c *gin.Context) {
	var req loginUserRequest
	var user models.User
	db := db.DBConn
	c.ShouldBind(&user)
	c.ShouldBind(&req)

	maker, err := token.NewPasetoMaker("lodrtnufepalofeiwnctvuqmifeasklo")
	if err != nil {
		fmt.Println(err)
	}

	db.Where("username = ?", req.Username).First(&user)

	err = CheckPassword(req.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := maker.CreateToken(user.Username, user.Id, time.Hour, user.IsAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	userResponse := userResponse{
		Username: user.Username,
		UserId:   user.Id,
	}
	c.SetCookie(AuthorizationPayloadKey, token, 60*60*24, "/", "localhost", true, true)
	c.JSON(http.StatusOK, userResponse)
}

func ReadHeaders(c *gin.Context) {
	payload := c.MustGet(AuthorizationPayloadKey)
	data := payload.(*token.Payload)
	fmt.Println(data.Username)
	c.JSON(200, data)
}

func Register(c *gin.Context) {
	var user models.User
	maker, err := token.NewPasetoMaker("lodrtnufepalofeiwnctvuqmifeasklo")
	db := db.DBConn
	c.ShouldBind(&user)

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("error hasing password"))
		return
	}

	user.Password = hashedPassword
	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error.Error())
		return
	}
	token, err := maker.CreateToken(user.Username, user.Id, time.Hour, user.IsAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	userResponse := userResponse{
		Username: user.Username,
		UserId:   user.Id,
	}
	c.SetCookie(AuthorizationPayloadKey, token, 60*60*24, "/", "localhost", false, false)
	c.JSON(201, userResponse)
}
