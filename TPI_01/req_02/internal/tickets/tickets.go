package tickets

import (
	"TPI_01/internal/utils"
	"encoding/csv"
	"fmt"
	"os"
)

type CSVRow struct {
	ID       string
	Name     string
	Email    string
	Country  string
	Time     string
	Quantity string
}

type CountCalculator interface {
	GetCount(rows []CSVRow) (int, error)
}

type MadrugadaCalculator struct{}

func (c MadrugadaCalculator) GetCount(rows []CSVRow) (int, error) {
	count := 0
	for _, row := range rows {
		hour, _, _ := utils.ParseTime(row.Time)
		if hour >= 0 && hour < 6 {
			count++
		}
	}
	return count, nil
}

type MananaCalculator struct{}

func (c MananaCalculator) GetCount(rows []CSVRow) (int, error) {
	count := 0
	for _, row := range rows {
		hour, _, _ := utils.ParseTime(row.Time)
		if hour >= 7 && hour <= 12 {
			count++
		}
	}
	return count, nil
}

type TardeCalculator struct{}

func (c TardeCalculator) GetCount(rows []CSVRow) (int, error) {
	count := 0
	for _, row := range rows {
		hour, _, _ := utils.ParseTime(row.Time)
		if hour >= 13 && hour <= 19 {
			count++
		}
	}
	return count, nil
}

type NocheCalculator struct{}

func (c NocheCalculator) GetCount(rows []CSVRow) (int, error) {
	count := 0
	for _, row := range rows {
		hour, _, _ := utils.ParseTime(row.Time)
		if hour >= 20 && hour <= 23 {
			count++
		}
	}
	return count, nil
}

// Factory
func CreateCountCalculator(period string) CountCalculator {
	switch period {
	case "madrugada":
		return MadrugadaCalculator{}
	case "manana":
		return MananaCalculator{}
	case "tarde":
		return TardeCalculator{}
	case "noche":
		return NocheCalculator{}
	default:
		return nil
	}
}

func GetCountByPeriod(fileName, period string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	reader.FieldsPerRecord = -1
	lines, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}

	var rows []CSVRow
	for _, line := range lines {
		rows = append(rows, CSVRow{
			ID:       line[0],
			Name:     line[1],
			Email:    line[2],
			Country:  line[3],
			Time:     line[4],
			Quantity: line[5],
		})
	}

	calculator := CreateCountCalculator(period)
	if calculator == nil {
		return 0, fmt.Errorf("Período no válido: %s", period)
	}

	count, err := calculator.GetCount(rows)
	if err != nil {
		return 0, err
	}

	return count, nil
}
