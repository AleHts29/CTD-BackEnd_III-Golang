package utils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	Id          string
	Name        string
	Email       string
	Destination string
	FlightTime  string
	Price       string
}

func ReadCSV(filename string) ([]Ticket, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("no se pudo abrir el archivo")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("no se pudo leer el archivo")
	}
	var tickets []Ticket
	for _, line := range lines {
		ticket := Ticket{
			Id:          line[0],
			Name:        line[1],
			Email:       line[2],
			Destination: line[3],
			FlightTime:  line[4],
			Price:       line[5],
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

func ParseTime(timeStr string) (hour, minute int, err error) {
	splitTime := strings.Split(timeStr, ":")
	if len(splitTime) != 2 {
		err = fmt.Errorf("formato de hora no valido: %s", timeStr)
		return
	}

	hour, err = strconv.Atoi(splitTime[0])
	if err != nil {
		return 0, 0, errors.New("no se pudo convertir a INT")
	}
	if hour < 0 && hour > 24 {
		return 0, 0, errors.New("no es un horario valido")
	}

	minute, err = strconv.Atoi(splitTime[1])
	if err != nil {
		return 0, 0, errors.New("no se pudo convertir a INT")
	}
	if hour < 0 && hour > 59 {
		return 0, 0, errors.New("no es un horario valido")
	}
	return
}

func GetAllTickets() (int, error) {
	tickets, err := ReadCSV("tickets.csv")
	if err != nil {
		return 0, err
	}
	return len(tickets), nil
}
