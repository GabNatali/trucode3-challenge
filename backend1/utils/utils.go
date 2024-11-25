package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/GabNatali/trucode3-challenge-final/internal/adult"
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

func cleanText(value string) string {
	cleaned := strings.TrimSpace(value)
	if cleaned == "?" || cleaned == "" {
		return "NULL"
	}
	return strings.ReplaceAll(cleaned, "  ", " ")
}
