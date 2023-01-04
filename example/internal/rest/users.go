package rest

import (
	"github.com/oherych/experimental-service-kit/example/internal/locator"
	"github.com/oherych/experimental-service-kit/example/internal/permissions"
	"github.com/oherych/experimental-service-kit/example/internal/repository"
	"github.com/oherych/experimental-service-kit/example/internal/rest/schemas"
	listener "github.com/oherych/experimental-service-kit/pkg/echo-listener"
)

type UsersController struct {
	d *locator.Container
}

func (cc UsersController) List(c listener.Context) error {
	if err := c.ShouldCan(permissions.UserList); err != nil {
		return err
	}

	pagination, err := c.Pagination(100)
	if err != nil {
		return err
	}

	users, err := cc.d.Users.All(c.Context(), pagination)
	if err != nil {
		return err
	}

	result := make([]schemas.User, len(users))
	for i, item := range users {
		result[i] = cc.display(item)
	}

	return c.List(pagination, result)
}

func (cc UsersController) Get(c listener.Context) error {
	if err := c.ShouldCan(permissions.UserGet); err != nil {
		return err
	}

	id := c.ParamString("id")

	user, err := cc.d.Users.GetByID(c.Context(), id)
	if err == repository.ErrUserNotFound {
		return listener.NotFound{Reason: "user"}
	}
	if err != nil {
		return err
	}

	return c.StatusOK(cc.display(*user))
}

func (cc UsersController) Delete(c listener.Context) error {
	panic("implement me")
}

func (cc UsersController) display(in repository.User) schemas.User {
	return schemas.User{
		ID:       in.ID,
		Username: in.Username,
		Email:    in.Username,
	}
}
