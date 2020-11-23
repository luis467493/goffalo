package repository

import (
	"fmt"

	"github.com/gobuffalo/uuid"
	"github.com/products/models"
)

func FindProducts() ([]models.Product, error) {
	var listProduct []models.Product
	s := "select prd.id as ID, prd.\"NAME\" as Name, prd.color as Color, prd.\"TYPE\" as Type from \"PRODUCTS\" prd order by prd.\"TYPE\", prd.color, prd.\"NAME\";"
	if err := models.DB.RawQuery(s).All(&listProduct); err != nil {
		fmt.Errorf(err.Error())
	}
	values := make([]models.Product, 0, len(listProduct))
	for _, product := range listProduct {
		values = append(values, product)
	}
	return values, nil
}

func FindProductById(ID uuid.UUID) models.Product {
	foundProduct := models.Product{
		Name:  "",
	}

	s := "select prd.id as ID, prd.\"NAME\" as Name, prd.color as Color, prd.\"TYPE\" as Type from \"PRODUCTS\" prd where prd.id = '" + ID.String() + "';"
	if err := models.DB.RawQuery(s).First(&foundProduct); err != nil {
		fmt.Errorf(err.Error())
	}

	return foundProduct
}

func SaveProduct(product *models.Product) {
	s := "INSERT INTO public.\"PRODUCTS\"(\"NAME\", color, \"TYPE\")VALUES('" + product.Name + "', '" + product.Color.String() + "', '" + product.Type.String() + "');"
	if err := models.DB.RawQuery(s).Exec(); err != nil {
		fmt.Errorf(err.Error())
	}
}

func UpdateProduct(product *models.Product) {
	s := "UPDATE public.\"PRODUCTS\" SET \"NAME\"='" + product.Name + "', color='" + product.Color.String() + "', \"TYPE\"='" + product.Type.String() + "' WHERE id='" + product.ID.String() + "';"
	if err := models.DB.RawQuery(s).Exec(); err != nil {
		fmt.Errorf(err.Error())
	}
}
