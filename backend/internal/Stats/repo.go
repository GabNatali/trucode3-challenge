package stats

import (
	"fmt"

	"gorm.io/gorm"
)

type StatsRepository interface {
	GetStats(option string) ([]IncomeStatsByEducation, error)
}

type repo struct {
	db *gorm.DB
}

type IncomeStatsByEducation struct {
	Option          string `json:"option"`
	TotalCount      int    `json:"total_count"`
	HighIncomeCount int    `json:"high_income_count"`
	LowIncomeCount  int    `json:"low_income_count"`
}

func NewRepository(db *gorm.DB) StatsRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) GetStats(option string) ([]IncomeStatsByEducation, error) {

	var stats []IncomeStatsByEducation
	query := fmt.Sprintf(`
		SELECT 
			%s AS option,
			COUNT(*) AS total_count,
			SUM(CASE WHEN income = '<=50K' THEN 1 ELSE 0 END) AS high_income_count,
			SUM(CASE WHEN income = '>50K' THEN 1 ELSE 0 END) AS low_income_count
		FROM 
			adults
		GROUP BY 
			%s
		ORDER BY 
			%s
	`, option, option, option)

	if err := r.db.Raw(query).Scan(&stats).Error; err != nil {
		return []IncomeStatsByEducation{}, err
	}

	fmt.Println(stats)
	return stats, nil
}
