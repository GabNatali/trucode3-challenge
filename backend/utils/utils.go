package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/GabNatali/trucode3-challenge-final/internal/adult"
	"gorm.io/gorm"
)

func ReadFile(path string, linesChan chan<- []adult.Adult) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()

	var batch []adult.Adult
	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		lineCount++
		if lineCount == 1 {
			continue
		}

		fields := strings.Split(line, ",")
		if len(fields) < 15 {
			continue
		}

		age, _ := strconv.Atoi(cleanText(fields[0]))
		fnlwgt, _ := strconv.Atoi(cleanText(fields[2]))
		educationNum, _ := strconv.Atoi(cleanText(fields[4]))

		capitalGain, _ := strconv.Atoi(cleanText(fields[10]))
		capitalLoss, _ := strconv.Atoi(cleanText(fields[11]))
		hoursPerWeek, _ := strconv.Atoi(cleanText(fields[12]))

		record := adult.Adult{
			Age:           age,
			Workclass:     adult.Workclass(cleanText(fields[1])),
			Fnlwgt:        fnlwgt,
			Education:     adult.Education(cleanText(fields[3])),
			EducationNum:  adult.EducationNum(educationNum),
			MaritalStatus: adult.MaritalStatus(cleanText(fields[5])),
			Occupation:    adult.Occupation(cleanText(fields[6])),
			Relationship:  adult.Relationship(cleanText(fields[7])),
			Race:          adult.Race(cleanText(fields[8])),
			Sex:           adult.Sex(cleanText(fields[9])),
			CapitalGain:   capitalGain,
			CapitalLoss:   capitalLoss,
			HoursPerWeek:  hoursPerWeek,
			NativeCountry: adult.Country(cleanText(fields[13])),
			Income:        adult.Income(cleanText(fields[14])),
		}

		batch = append(batch, record)

		if len(batch) >= 500 {
			linesChan <- batch
			batch = nil
		}
	}

	if len(batch) > 0 {
		linesChan <- batch
	}

	return scanner.Err()
}

func MigrateData(db *gorm.DB) error {

	linesChan := make(chan []adult.Adult, 4)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for batch := range linesChan {
				tx := db.Create(&batch).Error
				if tx != nil {
					return
				}
			}
		}()
	}

	path := "../data/source.data"

	err := ReadFile(path, linesChan)

	if err != nil {
		return err
	}

	close(linesChan)
	wg.Wait()

	return nil

}

func cleanText(value string) string {
	cleaned := strings.TrimSpace(value)
	if cleaned == "?" || cleaned == "" {
		return "NULL"
	}
	return strings.ReplaceAll(cleaned, "  ", " ")
}
