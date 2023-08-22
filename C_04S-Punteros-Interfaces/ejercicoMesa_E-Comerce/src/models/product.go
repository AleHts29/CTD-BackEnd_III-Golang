package models

type SmallProduct struct {
	PriceP float64
}
type MediumProduct struct {
	PriceP float64
}
type BigProduct struct {
	PriceP float64
}

func (s SmallProduct) Price() float64 {
	return s.PriceP
}

func (m MediumProduct) Price() float64 {
	return m.PriceP * 1.03
}

func (b BigProduct) Price() float64 {
	return b.PriceP*1.06 + 2500
}
