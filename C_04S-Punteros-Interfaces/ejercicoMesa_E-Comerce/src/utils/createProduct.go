package utils

import "c_04s_e-comerce/src/models"

const (
	Large  = "large"
	Medium = "medium"
	Small  = "small"
)

// CrearProducto - Función factory que crea un Producto en función del tipo y el precio proporcionados
func CrearProducto(tipo string, precio float64) models.Product {
	switch tipo {
	case Small:
		return models.SmallProduct{precio}
	case Medium:
		return models.MediumProduct{precio}
	case Large:
		return models.BigProduct{precio}
	default:
		// Manejar un tipo de producto desconocido, puedes devolver un error o un producto genérico aquí.
		return nil
	}
}
