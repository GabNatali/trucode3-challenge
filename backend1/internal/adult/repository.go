package adult

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type AdultRepository struct {
	db *gorm.DB
}

func NewAdultRepository(db *gorm.DB) AdultRepository {
	return AdultRepository{db: db}
}

func (r *AdultRepository) FindAll(params QueryParams) ([]Adult, error) {
	var adults []Adult

	query := r.buildQuery(params)

	offset := (params.Page - 1) * params.PageSize

	result := query.Offset(offset).Limit(params.PageSize).Find(&adults)

	if result.Error != nil {
		return nil, result.Error
	}
	return adults, nil
}

func (r *AdultRepository) Count(params QueryParams) (int64, error) {

	query := r.buildQuery(params)

	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		return 0, err
	}

	return totalCount, nil

}

func (r *AdultRepository) ExportCSV(params QueryParams) ([]Adult, error) {

	var adults []Adult

	query := r.buildQuery(params)

	result := query.Find(&adults)

	if result.Error != nil {
		return nil, result.Error
	}
	return adults, nil
}

func (r *AdultRepository) buildQuery(params QueryParams) *gorm.DB {
	query := r.db.Model(&Adult{})

	if params.Education != "" {
		educationOptions := strings.Split(params.Education, ",")
		query = query.Where("education IN ?", educationOptions)
	}

	if params.MaritalStatus != "" {
		options := strings.Split(params.MaritalStatus, ",")
		query = query.Where("marital_status IN ?", options)
	}

	if params.Occupation != "" {
		options := strings.Split(params.Occupation, ",")
		query = query.Where("occupation IN ?", options)
	}

	if params.MinAge > 0 && params.MaxAge > 0 {
		query = query.Where("age BETWEEN ? AND ?", params.MinAge, params.MaxAge)
	}

	if params.Income != "" {
		query = query.Where("income= ?", params.Income)
	}

	if params.OrderBy != "" && (params.OrderDir == "asc" || params.OrderDir == "desc") {
		query = query.Order(fmt.Sprintf("%s %s", params.OrderBy, params.OrderDir))
	}

	fmt.Printf("query: %v\n", query)
	return query
}
