package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	Product struct{
		Id int `json:"id"`
		Name string `json:"name"`
		Quantity int `json:"quantity"`
		Price int `json:"price"`
	}

	Products []Product
) 

var (
	products = Products{}
	productId = 1
)

func homeAPI(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
}

func createProduct(c echo.Context) error {
	dataQuery := c.QueryParam("data")
	if dataQuery == "bulk" {
		p := &Products{}
		if err := c.Bind(p); err != nil {
			return err
		}
		for _, product := range *p {
			product.Id = productId
			products = append(products, product)
			productId++
		}
		return c.String(http.StatusCreated, "Bulk product created Success")
	} else {
		p := &Product{
			Id: productId,
		}
		if err := c.Bind(p); err != nil {
			return err
		}
		products = append(products, *p)
		productId++
		return c.String(http.StatusCreated, "Product created successfully")
	}
}

func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func getProductById(c echo.Context) error {
	product := products.findProductById(c.Param("id"))
	if product.Id != 0 {
		return c.JSON(http.StatusOK, product)
	} else {
		return c.String(http.StatusNotFound, "Product not found or doesnt exist")
	}
}

func updateProductById(c echo.Context) error {
	strid := c.Param("id")
	intid,_ := strconv.Atoi(strid)
	pd := new(Product)
	if err := c.Bind(pd); err != nil {
		return err
	}
	pIndex := indexOfProduct(intid, products)
	if pIndex != -1 {
		pd.Id = products[pIndex].Id
		products[pIndex] = *pd
		return c.String(http.StatusOK, "Update Success")
	}
	return c.String(http.StatusNotFound, "Update Failed. Product not found or does not exist")
}

func patchProductById(c echo.Context) error {
	p := &Product{}
	if err := c.Bind(p); err != nil {
			return err
		}
	intid,_ := strconv.Atoi(c.Param("id"))
	pIndex := indexOfProduct(intid, products)
	if pIndex != -1 {
		switch m := c.QueryParam("method"); m{
		case "add-stock":
			products[pIndex].Quantity = products[pIndex].Quantity + p.Quantity
			return c.String(http.StatusOK, "Stock added successfully")
		case "reduce-stock":
			products[pIndex].Quantity = products[pIndex].Quantity - p.Quantity
			return c.String(http.StatusOK, "Stock reduced successfully")
		case "update-price":
			products[pIndex].Price = p.Price
			return c.String(http.StatusOK, "Product price updated successfully")
		}
	}
	return c.String(http.StatusNotFound, "Patch update Failed. Product not found or does not exist")
}

func deleteProductById(c echo.Context) error {
	strid := c.Param("id")
	intid, _ := strconv.Atoi(strid)
	pIndex := indexOfProduct(intid, products)
	if pIndex != -1 {
		copy(products[pIndex:], products[pIndex+1:])
		products = products[:len(products)-1]
		return c.String(http.StatusOK, "Delete Success")
	} 
	return c.String(http.StatusNotFound, "Delete Failed. Product id not found or does not exist")
}

func (p Products) findProductById(strid string) Product {
	id, _ := strconv.Atoi(strid)
	for _, product := range p {
		if product.Id == id {
			return product
			break
		}
	}
	return Product{}
}

func indexOfProduct(s any, data Products) int {
	for k, v := range data {
		if v.Id == s {
			return k
		}
	}
	return -1
}

func main() {
	//init echo instances
	e := echo.New()

	// set Routing
	// Home API
	e.GET("/", homeAPI)
	//Get all products
	e.GET("/products", getProducts)
	//Add new product to products collection
	e.POST("/products", createProduct)
	//Get product by ID
	e.GET("/products/:id", getProductById)
	//Update product by ID
	e.PUT("/products/:id", updateProductById)
	//Update product by ID on selected key
	e.PATCH("/products/:id", patchProductById)
	//Delete product by ID
	e.DELETE("/products/:id", deleteProductById)

	//Start Server
	e.Logger.Fatal(e.Start(":1323"))
}