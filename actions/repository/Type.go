package repository

import (
	"fmt"

	"github.com/gobuffalo/uuid"
	"github.com/products/models"
)

func FindTypes() ([]models.Type, error) {
	var listType []models.Type
	s := "select c.id as ID , c.\"NAME\" as Name from \"TYPES\" c order by c.\"NAME\";"
	if err := models.DB.RawQuery(s).All(&listType); err != nil {
		fmt.Errorf(err.Error())
	}
	values := make([]models.Type, 0, len(listType))
	for _, varType := range listType {
		values = append(values, varType)
	}
	return values, nil
}

func FindTypeById(ID uuid.UUID) models.Color {
	foundColor := models.Color{
		Name:  "",
	}
	s := "select c.id as ID , c.\"NAME\" as Name from \"TYPES\" c where c.id = '" + ID.String() + "';"
	if err := models.DB.RawQuery(s).First(&foundColor); err != nil {
		fmt.Errorf(err.Error())
	}

	return foundColor
}

func GetTypeMap() map[string]uuid.UUID {
	types, _ := FindTypes()
	var mapTypes = make(map[string]uuid.UUID)
	for _, varType := range types {
		mapTypes[varType.Name]=varType.ID
	}
	return mapTypes
}

