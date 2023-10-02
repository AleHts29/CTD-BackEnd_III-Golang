package models

type Service struct {
	Name        string
	Price       float64
	CantMinutes int
}

func SumServices(services []Service) float64 {
	total := 0.0
	for _, product := range services {
		total += product.Price * float64(product.CantMinutes)
	}
	return total
}
