package adult

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
)

type AdultService interface {
	Find(params QueryParams) ([]Adult, error)
	ExportCSV(params QueryParams) ([]byte, error)
}
type adultService struct {
	repo AdultRepository
}

func NewAdultService(repo AdultRepository) AdultService {
	return &adultService{
		repo: repo,
	}
}

// FindAll implements AdultService.
func (a *adultService) Find(params QueryParams) ([]Adult, error) {
	return a.repo.Find(params)
}

func (a *adultService) ExportCSV(params QueryParams) ([]byte, error) {
	adults, err := a.repo.ExportCSV(params)

	if err != nil {
		return nil, err
	}

	var csvData bytes.Buffer
	writer := csv.NewWriter(&csvData)

	headers := []string{"Age", "Workclass", "Fnlwgt", "Education", "Education Num", "Marital Status", "Occupation", "Relationship", "Race", "Sex", "Capital Gain", "Capital Loss", "Hours per week", "Native Country", "Income"}

	if err := writer.Write(headers); err != nil {
		return nil, errors.New("failed to write headers to CSV")
	}

	for _, adult := range adults {
		row := []string{
			fmt.Sprintf("%v", adult.Age),
			fmt.Sprintf("%v", adult.Fnlwgt),
			fmt.Sprintf("%v", adult.Sex),
			fmt.Sprintf("%v", adult.Race),
			fmt.Sprintf("%v", adult.CapitalGain),
			fmt.Sprintf("%v", adult.CapitalLoss),
			fmt.Sprintf("%v", adult.HoursPerWeek),
			fmt.Sprintf("%v", adult.NativeCountry),
			fmt.Sprintf("%v", adult.Income),
			fmt.Sprintf("%v", adult.Workclass),
			fmt.Sprintf("%v", adult.Education),
			fmt.Sprintf("%v", adult.EducationNum),
			fmt.Sprintf("%v", adult.MaritalStatus),
			fmt.Sprintf("%v", adult.Occupation),
			fmt.Sprintf("%v", adult.Relationship),
		}
		if err := writer.Write(row); err != nil {
			return nil, fmt.Errorf("failed to write CSV row: %w", err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, fmt.Errorf("failed to flush CSV writer: %w", err)
	}

	return csvData.Bytes(), nil
}
