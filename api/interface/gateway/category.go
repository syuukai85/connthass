package gateway

import (
	"fmt"

	"github.com/connthass/connthass/api/entity"
	"github.com/connthass/connthass/api/infrastructure/orm/model"
)

func categoriesToEntities(categories []model.Category) []*entity.Category {
	entityCategories := make([]*entity.Category, 0)

	for _, category := range categories {
		entityCategoriy := &entity.Category{
			ID:   entity.CategoryID(fmt.Sprint(category.ID)),
			Name: category.Name,
		}
		entityCategories = append(entityCategories, entityCategoriy)
	}

	return entityCategories
}

func categoriesToModels(categories []*entity.Category) []model.Category {
	modelCategories := make([]model.Category, 0)

	for _, category := range categories {
		modelCategory := model.Category{
			Name: category.Name,
		}
		modelCategories = append(modelCategories, modelCategory)
	}

	return modelCategories
}
