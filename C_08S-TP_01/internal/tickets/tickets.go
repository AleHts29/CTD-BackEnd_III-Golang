package tickets

import (
	"errors"
	"fmt"
	"github.com/AleHts29/parcial_backend_III_g6/internal/utils"
)

const (
	EARLYMORNING = "madrugada"
	MORNIG       = "mañana"
	AFTERNOON    = "tarde"
	NIGHT        = "noche"
)

/*
Requerimiento 1
Una función que calcule cuántas personas viajan a un país determinado.
*/
func GetTotalTickets(destination string) (int, error) {
	auxCout := 0
	tickets, err := utils.ReadCSV("tickets.csv")
	if err != nil {
		return 0, err
	}
	for _, ticket := range tickets {
		if ticket.Destination == destination {
			auxCout++
		}
	}
	return auxCout, nil
}

/*
Requerimiento 2:
Una o varias funciones que calculen cuántas personas viajan en madrugada (0 → 6),
mañana (7 → 12), tarde (13 → 19), y noche (20 → 23).
*/

// CountCalculator - Definimos una interfaz para las funciones de cálculo
type CountCalculator interface {
	GetCount(rows []utils.Ticket) (int, error)
}

// MadrugadaCalculator - Estructura para calcular el recuento en el período de madrugada
type MadrugadaCalculator struct{}

func (c MadrugadaCalculator) GetCount(rows []utils.Ticket) (int, error) {
	count := 0
	for _, row := range rows {
		hour, _, err := utils.ParseTime(row.FlightTime)
		if err != nil {
			return 0, err
		}
		if hour >= 0 && hour < 6 {
			count++
		}
	}
	return count, nil
}

// MananaCalculator - Estructura para calcular el recuento en el período de mañana
type MananaCalculator struct{}

func (c MananaCalculator) GetCount(rows []utils.Ticket) (int, error) {
	count := 0
	for _, row := range rows {
		hour, _, err := utils.ParseTime(row.FlightTime)
		if err != nil {
			return 0, err
		}
		if hour >= 7 && hour <= 12 {
			count++
		}
	}
	return count, nil
}

// TardeCalculator - Estructura para calcular el recuento en el período de tarde
type TardeCalculator struct{}

func (c TardeCalculator) GetCount(rows []utils.Ticket) (int, error) {
	count := 0
	for _, row := range rows {
		hour, _, err := utils.ParseTime(row.FlightTime)
		if err != nil {
			return 0, err
		}
		if hour >= 13 && hour <= 19 {
			count++
		}
	}
	return count, nil
}

// NocheCalculator - Estructura para calcular el recuento en el período de noche
type NocheCalculator struct{}

func (c NocheCalculator) GetCount(rows []utils.Ticket) (int, error) {
	count := 0
	for _, row := range rows {
		hour, _, err := utils.ParseTime(row.FlightTime)
		if err != nil {
			return 0, err
		}
		if hour >= 20 && hour <= 23 {
			count++
		}
	}
	return count, nil
}

// CreateCountCalculator - Factory
func CreateCountCalculator(period string) CountCalculator {
	switch period {
	case EARLYMORNING:
		return MadrugadaCalculator{}
	case MORNIG:
		return MananaCalculator{}
	case AFTERNOON:
		return TardeCalculator{}
	case NIGHT:
		return NocheCalculator{}
	default:
		return nil
	}
}

func GetCountByPeriod(period string) (int, error) {
	tickets, _ := utils.ReadCSV("tickets.csv")

	calculator := CreateCountCalculator(period)
	if calculator == nil {
		return 0, fmt.Errorf("el periodo %s no es valido", period)
	}
	count, err := calculator.GetCount(tickets)
	if err != nil {
		return 0, err
	}
	return count, nil
}

/*
Requerimiento 3:
Calcular el porcentaje de personas que viajan a un país determinado en un día.
*/

func PercentageDestination(destination string, total int) (float64, error) {
	countDestination, err := GetTotalTickets(destination)
	if err != nil {
		return 0, err
	}
	if total == 0 {
		return 0, errors.New("el total de vuelos no puede ser 0")
	}
	percentage := (float64(countDestination) * 100) / float64(total)
	return percentage, nil
}
