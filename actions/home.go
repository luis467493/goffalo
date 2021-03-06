package actions

import (
	"encoding/base64"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gofrs/uuid"
	"github.com/products/actions/repository"
	"github.com/products/models"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("index.html"))
}

func encodeBase64(ID uuid.UUID) string {
	return base64.StdEncoding.EncodeToString([]byte(ID.String()))
}

func decodeBase64(ID string) uuid.UUID{
	decoded, _ := base64.StdEncoding.DecodeString(ID)
	decodedId, _ := uuid.FromString(string(decoded))
	return decodedId
}

func findProducts(c buffalo.Context) error {
	products, _ := repository.FindProducts()
	listProducts := make([]models.ProductDto, 0, len(products))
	for _, product := range products {
		listProducts = append(listProducts, models.ProductDto{
			ID:    encodeBase64(product.ID),
			Name:  product.Name,
			Color: repository.FindColorById(product.Color).Name,
			Type:  repository.FindTypeById(product.Type).Name,
		})
	}
	c.Set("products", listProducts)
	return c.Render(http.StatusOK, r.HTML("products.html"))
}

func newProduct(c buffalo.Context) error {
	c.Set("colors", repository.GetColorMap())
	c.Set("types", repository.GetTypeMap())

	product := &models.Product{}
	c.Set("product", product)

	return c.Render(http.StatusOK, r.HTML("productNew.html"))
}

func editProduct(c buffalo.Context) error {
	c.Set("colors", repository.GetColorMap())
	c.Set("types", repository.GetTypeMap())

	id := decodeBase64(c.Param("product_id"))

	product := repository.FindProductById(id)
	productDto := models.ProductDto{
		ID:    encodeBase64(product.ID),
		Name:  product.Name,
		Color: product.Color.String(),
		Type:  product.Type.String(),
	}

	c.Set("product", productDto)
	return c.Render(http.StatusOK, r.HTML("productEdit.html"))
}

func saveProduct(c buffalo.Context) error {
	id := decodeBase64(c.Param("product_id"))
	color, _ := uuid.FromString(c.Param("Color"))
	varType, _ := uuid.FromString(c.Param("Type"))

	repository.SaveProduct(&models.Product{
		ID:    id,
		Name:  c.Param("Name"),
		Color: color,
		Type:  varType,
	})

	return findProducts(c)
}

func updateProduct(c buffalo.Context) error {
	id := decodeBase64(c.Param("product_id"))
	color, _ := uuid.FromString(c.Param("Color"))
	varType, _ := uuid.FromString(c.Param("Type"))

	repository.UpdateProduct(&models.Product{
		ID:    id,
		Name:  c.Param("Name"),
		Color: color,
		Type:  varType,
	})

	return findProducts(c)
}

func removeProduct(c buffalo.Context) error {
	id := decodeBase64(c.Param("product_id"))
	repository.RemoveProduct(id)

	return findProducts(c)
}
