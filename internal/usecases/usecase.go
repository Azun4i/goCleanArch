package usecases

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"goCleanArch/internal/model"
	"goCleanArch/internal/repository"
	"strconv"
)

type UseCase struct {
	store repository.Repository
	idcnt int
}

type UseCaseLogic interface {
	Create(u *model.User) error
	Delete(id string) error
	Edit(u *model.User) error
	FindById(id string) (*model.User, error)
	Validation(u *model.User) error //// узнать
}

//NewUseCase return interfase of UseCaseLogic
func NewUseCase(store repository.Repository) UseCaseLogic {
	return &UseCase{
		store: store,
		idcnt: 1,
	}
}

func (c *UseCase) Create(u *model.User) error {
	if err := c.Validation(u); err != nil {
		return err
	}
	u.ID = strconv.Itoa(c.idcnt)
	if err := c.store.Create(u); err != nil {
		return err
	}
	c.idcnt++
	return nil
}

func (c *UseCase) Delete(id string) error {

	if err := c.store.Delete(id); err != nil {
		return err
	}
	return nil
}

func (c *UseCase) Edit(u *model.User) error {
	if err := c.Validation(u); err != nil {
		return err
	}

	if err := c.store.Edit(u); err != nil {
		return err
	}
	return nil
}

func (c *UseCase) FindById(id string) (*model.User, error) {

	u, err := c.store.FindById(id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

//validation get validation validationAge and validationEmail  return err
func (c *UseCase) Validation(u *model.User) error {

	if err := c.validationAge(u.Age); err != nil {
		return err
	}

	if err := c.validationEmail(u); err != nil {
		return err
	}

	return nil
}

func (c *UseCase) validationAge(age string) error {
	i, _ := strconv.Atoi(age)
	if i < 18 {
		return errors.New("too young")
	}
	return nil
}

func (c *UseCase) validationEmail(u *model.User) error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
	)
}
