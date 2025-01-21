package entities

import (
	"fmt"
	"time"
)

type Category struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	CreteAt  string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

func NewCategory(name string) (*Category, error) {
	category := &Category{
		Name:    name,
		CreteAt: time.Now().String(),
		UpdateAt: time.Now().String(),
	}

	err := category.IsValid()

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c *Category) IsValid() error {
	if (len(c.Name) < 5 ) {
		return fmt.Errorf("Name must be greater than 5, got %d", len(c.Name))
	}
	return nil
}