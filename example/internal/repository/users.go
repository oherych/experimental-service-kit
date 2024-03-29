package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/oherych/experimental-service-kit/kit"

	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound = errors.New("user not found in the database")
)

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"column:username"`
	Email    string `gorm:"column:email"`
}

type Users struct {
	db *gorm.DB
}

func NewUsers(con *sql.DB) (Users, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{Conn: con}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return Users{}, err
	}

	return Users{db: db.Unscoped()}, nil
}

// All return a list of Users
func (r Users) All(ctx context.Context, pagination kit.Pagination) ([]User, error) {
	var target []User
	result := r.db.WithContext(ctx).Find(&target)

	return target, result.Error
}

// GetByID return user by ID
// If user not found method will return ErrUserNotFound
func (r Users) GetByID(ctx context.Context, id int) (*User, error) {
	var target User
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&target)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}

	return &target, nil
}

func (r Users) Delete(ctx context.Context, id int) error {
	panic("implement me")
}
