package rest

import (
	"github.com/oherych/experimental-service-kit/example/internal/locator"
	"github.com/oherych/experimental-service-kit/example/internal/repository"
	"github.com/oherych/experimental-service-kit/example/internal/rest/schemas"
	"github.com/oherych/experimental-service-kit/kit/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UsersController struct {
	d locator.Implementation
}

func (cc UsersController) List(c echo.Context) error {
	users, err := cc.d.Users.All(c.Request().Context())
	if err != nil {
		return err
	}

	result := make([]schemas.User, len(users))
	for i, item := range users {
		result[i] = cc.display(item)
	}

	return c.JSON(http.StatusOK, result)
}

func (cc UsersController) Get(c echo.Context) error {
	id := c.Param("id")

	user, err := cc.d.Users.GetByID(c.Request().Context(), id)
	if err == repository.ErrUserNotFound {
		return responses.NotFound{Reason: "user"}
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, cc.display(*user))
}

func (cc UsersController) display(in repository.User) schemas.User {
	return schemas.User{
		ID:       in.ID,
		Username: in.Username,
		Email:    in.Username,
	}
}
