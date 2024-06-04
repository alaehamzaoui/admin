package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/bo-mathadventure/admin/config"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/user"
	"github.com/bo-mathadventure/admin/utils"
	email "github.com/cameronnewman/go-emailvalidation/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// NewUserHandler initialize routes for the given router
func NewUserHandler(ctx context.Context, app fiber.Router, db *ent.Client) {
	app.Get("/", getMe(ctx, db))
	app.Put("/", updateUser(ctx, db))
	app.Get("/token", getTokenLogin(ctx, db))
}

type userResponse struct {
	UUID        string        `json:"uuid"`
	Email       string        `json:"email"`
	Username    string        `json:"username"`
	Permissions []string      `json:"permissions"`
	Tags        []string      `json:"tags"`
	LastLogin   time.Time     `json:"lastLogin" validate:"omitempty"`
	CreatedAt   time.Time     `json:"createdAt"`
	Config      config.Config `json:"config"`
}

func responseUserResponse(thisUser *ent.User) *userResponse {
	return &userResponse{
		UUID:        thisUser.UUID,
		Email:       thisUser.Email,
		Username:    thisUser.Username,
		Permissions: utils.CombinePermissions(thisUser),
		Tags:        utils.CombineTags(thisUser),
		LastLogin:   thisUser.LastLogin,
		CreatedAt:   thisUser.CreatedAt,
		Config:      config.GetConfig(),
	}
}

// getMe godoc
//
//	@Summary		User Info
//	@Description	Get user details of logged-in user
//	@Security		ApiKeyAuth
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	userResponse
//	@Failure		401	{object}	APIResponse
//	@Failure		500	{object}	APIResponse
//	@Router			/system/user [get]
func getMe(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userID := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userID)).First(ctx)
		if err != nil {
			return HandleInternalError(c, err)
		}

		return c.JSON(responseUserResponse(thisUser))
	}
}

type updateUserRequest struct {
	EMail                    string `json:"email" example:"bob@exameple.com" format:"email" validate:"omitempty,email"`
	Username                 string `json:"username" example:"Bob" validate:"omitempty,alphaunicode,min=3,max=16"`
	ClearTextPassword        string `json:"newPassword" example:"my$ecur3P4$$word" validate:"omitempty,min=8"`
	ClearTextPasswordConfirm string `json:"confirmPassword" example:"my$ecur3P4$$word" validate:"omitempty,min=8,eqcsfield=ClearTextPassword"`
	ClearTextCurrentPassword string `json:"password" example:"my$ecur3P4$$word" validate:"omitempty"`
}

// updateUser godoc
//
//	@Summary		Update User
//	@Description	Update details of the logged-in user
//	@Security		ApiKeyAuth
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			params	body		updateUserRequest	true	"-"
//	@Success		200		{object}	userResponse
//	@Failure		400		{object}	APIResponse
//	@Failure		401		{object}	APIResponse
//	@Failure		404		{object}	APIResponse
//	@Failure		500		{object}	APIResponse
//	@Router			/system/user [put]
func updateUser(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userID := int(claims["id"].(float64))

		req := new(updateUserRequest)
		if err := c.BodyParser(req); err != nil {
			return HandleBodyParseError(c, err)
		}

		if valid, err := ValidateStruct(c, req); !valid {
			return err
		}

		foundUser, err := db.User.Query().WithGroups().Where(user.ID(userID)).First(ctx)

		if req.EMail != "" || req.ClearTextPassword != "" {
			if err != nil || foundUser == nil || !utils.CheckPasswordHash(req.ClearTextCurrentPassword, foundUser.Password) {
				return HandleError(c, "ERR_INVALID_PASSWORD")
			}
		}

		update := foundUser.Update()

		if req.EMail != "" {
			err := email.Validate(req.EMail)
			if err != nil {
				return HandleError(c, "ERR_EMAIL_INVALID")
			}

			existingUser, err := db.User.Query().Where(user.Email(email.Normalize(req.EMail))).Where(user.IDNotIn(userID)).Count(ctx)
			if existingUser > 0 || err != nil {
				return HandleError(c, "ERR_USER_EXISTS")
			}

			update = update.SetEmail(email.Normalize(req.EMail))
		}

		if req.Username != "" {
			update = update.SetUsername(req.Username)
		}

		if req.ClearTextPassword != "" && req.ClearTextPasswordConfirm == "" || req.ClearTextPassword == "" && req.ClearTextPasswordConfirm != "" {
			return HandleError(c, "ERR_INVALID_REQUEST")
		}
		if req.ClearTextPassword != "" && req.ClearTextPasswordConfirm != "" {
			if req.ClearTextPassword != req.ClearTextPasswordConfirm {
				return HandleError(c, "ERR_PASSWORD_EQUAL")
			}

			hashedPassword, err := utils.HashPassword(req.ClearTextPassword)
			if err != nil {
				return HandleInternalError(c, err)
			}
			update = update.SetPassword(hashedPassword)
		}

		updatedUser, err := update.Save(ctx)
		if err != nil {
			return HandleInternalError(c, err)
		}

		return c.JSON(responseUserResponse(updatedUser))
	}
}

// getTokenLogin godoc
//
//	@Summary		Workadventure token
//	@Description	Generate JWT Token of logged-in user and directly redirect user to Workadventure
//	@Security		ApiKeyAuth
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		302	{object}	nil
//	@Failure		401	{object}	APIResponse
//	@Failure		500	{object}	APIResponse
//	@Router			/system/user/token [get]
func getTokenLogin(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userID := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userID)).First(ctx)
		if err != nil {
			return HandleInternalError(c, err)
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"identifier":  thisUser.UUID, // fixme may be also email
			"accessToken": nil,
			"username":    thisUser.Username,
		})

		tokenString, err := token.SignedString([]byte(config.GetConfig().WorkadventureSecretKey))
		if err != nil {
			return HandleInternalError(c, err)
		}
		url := fmt.Sprintf("%s?token=%s", config.GetConfig().WorkadventureURL, tokenString)
		return c.JSON(fiber.Map{
			"url": url,
		})
	}
}
