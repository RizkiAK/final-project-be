package admin

import (
	"blog-mandalika/domain"
	"blog-mandalika/helpers"
	"context"
	"net/http"
	"time"

	"github.com/Yureka-Teknologi-Cipta/yureka/response"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (u *adminUsecase) Login(ctx context.Context, options map[string]interface{}) response.Base {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	payload := options["payload"].(domain.LoginAdminRequest)

	errValidation := make(map[string]string)
	// validating request
	if payload.Email == "" {
		errValidation["email"] = "email field is required"
	}

	if payload.Password == "" {
		errValidation["password"] = "password field is required"
	}

	if len(errValidation) > 0 {
		return response.ErrorValidation(errValidation, "error validation")
	}

	// check the db
	admin, err := u.mysqlRepo.FetchOneAdmin(ctx, map[string]interface{}{
		"email": payload.Email,
	})
	if err != nil {
		return response.Error(http.StatusInternalServerError, err.Error())
	}

	// check password
	if err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(payload.Password)); err != nil {
		return response.Error(http.StatusBadRequest, "Wrong password")
	}

	// generate token
	tokenString, err := helpers.GenerateJWTTokenAdmin(domain.JWTClaimAdmin{
		UserID: admin.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:       uuid.NewString(),
			Issuer:   "admin",
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	})
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error())
	}

	return response.Success(map[string]interface{}{
		"user":  admin,
		"token": tokenString,
	})
}

func (u *adminUsecase) CreateBlog(ctx context.Context, options map[string]interface{}) response.Base {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	payload := options["payload"].(domain.CreateBlogRequest)

	errValidation := make(map[string]string)
	// validating request
	if payload.Title == "" {
		errValidation["title"] = "title field is required"
	}

	if payload.Content == "" {
		errValidation["content"] = "content field is required"
	}

	if payload.Category == "" {
		errValidation["category"] = "category field is required"
	}

	if len(errValidation) > 0 {
		return response.ErrorValidation(errValidation, "error validation")
	}

	// check the db
	err := u.mysqlRepo.CreateBlog(ctx, domain.Blog{
		ID:        uuid.NewString(),
		Title:     payload.Title,
		Content:   payload.Content,
		MediaUrl:  "",
		Category:  payload.Category,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return response.Error(http.StatusInternalServerError, err.Error())
	}

	return response.Success(map[string]interface{}{
		"message": "success create blog",
	})
}
