package repository

import "github.com/lesson-golang-unit-test-pzn/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
