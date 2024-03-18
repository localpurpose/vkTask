package handlers

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/localpurpose/vk-filmoteka/models"
	"github.com/localpurpose/vk-filmoteka/pkg/database/postgres"
	"github.com/localpurpose/vk-filmoteka/pkg/hash"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// UserSignUp godoc
//
// @Security ApiKeyAuth
// @Summary	Create user endpoint
// @Tags		users
// @Accept		json
// @Produce	json
// @Success	200
// @Router		/user/sign-up/ [post]
func UserSignUp(w http.ResponseWriter, r *http.Request) {

	var user models.User

	b, err := io.ReadAll(r.Body)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.Unmarshal(b, &user)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	user.Password, err = hash.HashPassword(user.Password)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = postgres.DB.DB.Create(&user).Error; err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newJsonResponse(w, http.StatusOK, map[string]interface{}{
		"ID":       user.ID,
		"Username": user.Username,
		"Password": user.Password,
		"Role":     user.Role,
	})
}

// UserSignIn godoc
//
// @Security ApiKeyAuth
// @Summary	Signing In user endpoint
// @Tags		users
// @Accept		json
// @Produce	json
// @Success	200
// @Router		/user/sign-in/ [post]
func UserSignIn(w http.ResponseWriter, r *http.Request) {

	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	var input LoginInput

	err = json.Unmarshal(b, &input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	userModel, err := getUserByUsername(input.Username)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	} else if userModel == nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !hash.CheckPwdHash(input.Password, userModel.Password) {
		newErrorResponse(w, http.StatusBadRequest, "Invalid username or password")
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userModel.Username
	claims["user_id"] = userModel.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["role"] = userModel.Role

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println("some error while signing JWT string")
		return
	}

	newJsonResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Successfully authorized!",
		"token":   t,
	})
}

func getUserByUsername(u string) (*models.User, error) {
	user := new(models.User)

	if err := postgres.DB.DB.Where(&models.User{Username: u}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
