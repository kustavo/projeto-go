package model

import (
	"github.com/kustavo/benchmark/go/domain"
	"github.com/kustavo/benchmark/go/domain/message"
)

// swagger:model user
type User struct {
	Base
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func NewUser(username string, password string, name string, email string) (*User, error) {
	user := &User{
		Username: username,
		Password: password,
		Name:     name,
		Email:    email,
	}

	err := user.Validate()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) SetId(id uint64) {
	u.ID = id
}

func (u *User) Validate() error {
	var errs []error

	if u.Username == "" {
		errs = append(errs, message.ErrUsernameIsRequired)
	}

	if u.Password == "" {
		errs = append(errs, message.ErrPasswordIsRequired)
	}

	if u.Name == "" {
		errs = append(errs, message.ErrNameIsRequired)
	}

	if u.Email == "" {
		errs = append(errs, message.ErrEmailIsRequired)
	}

	if len(errs) == 0 {
		return nil
	}

	return domain.NewErrorsList(errs)
}
