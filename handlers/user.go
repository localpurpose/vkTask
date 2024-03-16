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

func UserSignUp(w http.ResponseWriter, r *http.Request) {

	//TODO check if users username is already exists

	var user models.User

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("some error while reading body", err)
		return
	}

	err = json.Unmarshal(b, &user)
	if err != nil {
		log.Println("some error while unmarshalling json", err)
		return
	}

	user.Password, err = hash.HashPassword(user.Password)
	if err != nil {
		log.Println("error hashing user password", err)
		return
	}

	if err = postgres.DB.DB.Create(&user).Error; err != nil {
		log.Println("some error while creating user profile")
		return
	}

	w.Write([]byte("user created\n"))
	w.Write(b)
}

func UserSignIn(w http.ResponseWriter, r *http.Request) {

	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error while reading body", err)
		return
	}

	var input LoginInput

	err = json.Unmarshal(b, &input)
	if err != nil {
		log.Println("some error while unmarshalling body", err)
		return
	}

	userModel, err := getUserByUsername(input.Username)
	if err != nil {
		log.Println("some error while getting user by username", err)
		return
	} else if userModel == nil {
		log.Println("userModel is nil", err)
		return
	}

	if !hash.CheckPwdHash(input.Password, userModel.Password) {
		w.Write([]byte("Invalid username or password"))
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

	w.Write([]byte("Successful signing-in.\n"))
	w.Write([]byte("" +
		"{\njwt:" + t + "\n}",
	))
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
