package adult

import (
	"fmt"

	"gorm.io/gorm"
)

type AdultRepository struct {
	db *gorm.DB
}

func NewAdultRepository(db *gorm.DB) AdultRepository {
	return AdultRepository{db: db}
}

func (r *AdultRepository) Find(params QueryParams) ([]Adult, error) {
	var adults []Adult

	query := r.buildQuery(params)

	offset := (params.Page - 1) * params.PageSize
	query = query.Limit(params.PageSize).Offset(offset)

	result := query.Find(&adults)

	if result.Error != nil {
		return nil, result.Error
	}

	return adults, nil

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
		query = query.Where("education = ?", params.Education)
	}

	if params.MaritalStatus != "" {
		query = query.Where("marital_status = ?", params.MaritalStatus)
	}

	if params.Occupation != "" {
		query = query.Where("occupation = ?", params.Occupation)
	}

	if params.MinAge > 0 && params.MaxAge > 0 {
		query = query.Where("age BETWEEN ? AND ?", params.MinAge, params.MaxAge)
	}

	if params.IncomeLevel != "" {
		query = query.Where("income_level = ?", params.IncomeLevel)
	}

	if params.OrderBy != "" && (params.OrderDir == "asc" || params.OrderDir == "desc") {
		query = query.Order(fmt.Sprintf("%s %s", params.OrderBy, params.OrderDir))
	}

	return query
}
