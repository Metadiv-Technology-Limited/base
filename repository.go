package base

import (
	"github.com/Metadiv-Technology-Limited/sql"
	"github.com/Metadiv-Technology-Limited/sql/models"
	"gorm.io/gorm"
)

type Repository[T any] struct{}

func (r *Repository[T]) Save(tx *gorm.DB, entity *T) (*T, error) {
	return sql.Save(tx, entity)
}

func (r *Repository[T]) SaveAll(tx *gorm.DB, entities []T) ([]T, error) {
	return sql.SaveAll(tx, entities)
}

func (r *Repository[T]) Delete(tx *gorm.DB, entity *T) error {
	return sql.Delete(tx, entity)
}

func (r *Repository[T]) DeleteAll(tx *gorm.DB, entities []T) error {
	return sql.DeleteAll(tx, entities)
}

func (r *Repository[T]) DeleteBy(tx *gorm.DB, clause models.IClause) error {
	return sql.DeleteBy[T](tx, clause)
}

func (r *Repository[T]) FindOne(tx *gorm.DB, clause models.IClause) (*T, error) {
	return sql.FindOne[T](tx, clause)
}

func (r *Repository[T]) FindAll(tx *gorm.DB, clause models.IClause) ([]T, error) {
	return sql.FindAll[T](tx, clause)
}

func (r *Repository[T]) FindAllComplex(tx *gorm.DB, clause models.IClause, page models.IPagination, sort models.ISorting) ([]T, models.IPagination, error) {
	return sql.FindAllComplex[T](tx, clause, page, sort)
}

func (r *Repository[T]) Count(tx *gorm.DB, clause models.IClause) (int64, error) {
	return sql.Count[T](tx, clause)
}

func (r *Repository[T]) FindByID(tx *gorm.DB, id uint) (*T, error) {
	return sql.FindOne[T](tx, sql.Eq("id", id))
}
