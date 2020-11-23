package repository

import (
	"fmt"

	"github.com/gobuffalo/uuid"
	"github.com/products/models"
)

func FindColors() ([]models.Color, error) {
	var listColor []models.Color
	s := "select c.id as ID , c.\"NAME\" as Name from \"COLORS\" c order by c.\"NAME\";"
	if err := models.DB.RawQuery(s).All(&listColor); err != nil {
		fmt.Errorf(err.Error())
	}
	values := make([]models.Color, 0, len(listColor))
	for _, color := range listColor {
		values = append(values, color)
	}
	return values, nil
}

func FindColorById(ID uuid.UUID) models.Color {
	foundColor := models.Color{
		Name:  "",
	}
	s := "select c.id as ID , c.\"NAME\" as Name from \"COLORS\" c where c.id = '" + ID.String() + "';"
	if err := models.DB.RawQuery(s).First(&foundColor); err != nil {
		fmt.Errorf(err.Error())
	}

	return foundColor
}

func GetColorMap() map[string]uuid.UUID {
	colors, _ := FindColors()
	var mapColors = make(map[string]uuid.UUID)
	for _, color := range colors {
		mapColors[color.Name]=color.ID
	}
	return mapColors
}
