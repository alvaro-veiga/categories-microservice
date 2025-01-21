package usecases

import (
	"log"

	"github.com/alvaro-veiga/categories-microservice/internal/entities"
)

type CreateCategoryUseCase struct {
}

func NewCreateCategoryUseCase() *CreateCategoryUseCase {
	return &CreateCategoryUseCase{}
}

func (u *CreateCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}


	// TODO: persist entity to database
	log.Println(category)

	return nil
}