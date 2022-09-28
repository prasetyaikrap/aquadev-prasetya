package main

import (
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type (
	IProduct interface{
		findById(id int) Product
	}

	Product struct{
		Id int `json:"id"`
		Name string `json:"name"`
		Quantity int `json:"quantity"`
		Price int `json:"price"`
	}

	Products []Product

	SuccessResponse struct {
		Code int
		Message string
		Data interface{}
	}

	ErrorResponse struct {
		Code int
		Message string
		Error string
	}
)

var (
	products = Products{}
	productId = 1
)

func homeAPI(c echo.Context) error {
		return c.JSON(http.StatusOK, SuccessResponse{
			Code: http.StatusOK,
			Message: "API is Active",
			Data: "Home API",
		})
}

func createProduct(c echo.Context) error {
	dataQuery := c.QueryParam("data")
	if dataQuery == "bulk" {
		p := &Products{}
		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Code: http.StatusBadRequest,
				Message: "Invalid request body",
				Error: err.Error(),
			})
		}
		for _, product := range *p {
			product.Id = productId
			products = append(products, product)
			productId++
		}
		return c.JSON(http.StatusCreated, SuccessResponse{
			Code: http.StatusCreated,
			Message: "Bulk products created successfully",
			Data: p,
		})
	} else {
		p := &Product{
			Id: productId,
		}
		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Code: http.StatusBadRequest,
				Message: "Invalid request body",
				Error: err.Error(),
			})
		}
		products = append(products, *p)
		productId++
		return c.JSON(http.StatusCreated, SuccessResponse{
			Code: http.StatusCreated,
			Message: "Product created successfully",
			Data: p,
		})
	}
}

func getProducts(c echo.Context) error {
	if len(products) <= 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{
				Code: http.StatusNotFound,
				Message: "Get products failed.",
				Error: "Products is empty",
			})
	}
	return c.JSON(http.StatusOK, SuccessResponse{
		Code: http.StatusOK,
		Message: "Products found",
		Data: products,
	})
}

func getProductById(c echo.Context) error {
	intId, _ := strconv.Atoi(c.Param("id"))
	product := products.findById(intId)
	if product.Id != 0 {
		return c.JSON(http.StatusOK, SuccessResponse{
			Code: http.StatusOK,
			Message: "Product found",
			Data: product,
		})
	} else {
		return c.JSON(http.StatusNotFound, ErrorResponse{
			Code: http.StatusNotFound,
			Message: "Get product failed",
			Error: "Product not found or does not exist",
		})
	}
}

func updateProductById(c echo.Context) error {
	strid := c.Param("id")
	intid,_ := strconv.Atoi(strid)
	pd := new(Product)
	if err := c.Bind(pd); err != nil {
		return err
	}
	pIndex := indexOf(intid, products)
	if pIndex != -1 {
		pd.Id = products[pIndex].Id
		products[pIndex] = *pd
		return c.JSON(http.StatusOK, SuccessResponse{
			Code: http.StatusOK,
			Message: "Product updated successfully",
			Data: products[pIndex],
		})
	}
	return c.JSON(http.StatusNotFound, ErrorResponse{
		Code: http.StatusNotFound,
		Message: "Update product failed",
		Error: "Product not found",
	})
}

func patchProductById(c echo.Context) error {
	p := &Product{}
	if err := c.Bind(p); err != nil {
			return err
		}
	intid,_ := strconv.Atoi(c.Param("id"))
	pIndex := indexOf(intid, products)
	if pIndex != -1 {
		switch m := c.QueryParam("method"); m{
		case "add-stock":
			products[pIndex].Quantity = products[pIndex].Quantity + p.Quantity
			return c.JSON(http.StatusOK, SuccessResponse{
				Code: http.StatusOK,
				Message: "Product stock added successfully",
				Data: products[pIndex],
			})
		case "reduce-stock":
			products[pIndex].Quantity = products[pIndex].Quantity - p.Quantity
			return c.JSON(http.StatusOK, SuccessResponse{
				Code: http.StatusOK,
				Message: "Product stock reduced successfully",
				Data: products[pIndex],
			})
		case "update-price":
			products[pIndex].Price = p.Price
			return c.JSON(http.StatusOK, SuccessResponse{
				Code: http.StatusOK,
				Message: "Product price updated successfully",
				Data: products[pIndex],
			})
		}
	}
	return c.JSON(http.StatusNotFound, ErrorResponse{
		Code: http.StatusNotFound,
		Message: "Patch update product failed",
		Error: "Product not found",
	})
}

func deleteProductById(c echo.Context) error {
	strid := c.Param("id")
	intid, _ := strconv.Atoi(strid)
	pIndex := indexOf(intid, products)
	if pIndex != -1 {
		copy(products[pIndex:], products[pIndex+1:])
		products = products[:len(products)-1]
		return c.JSON(http.StatusOK, SuccessResponse{
			Code: http.StatusOK,
			Message: "Product Deleted successfully",
			Data: "Deleted",
		})
	} 
	return c.JSON(http.StatusNotFound, c.JSON(http.StatusNotFound, ErrorResponse{
		Code: http.StatusNotFound,
		Message: "Delete product failed",
		Error: "Product not found or does not exist",
	}))
}

func (p Products) findById(id int) Product {
	for _, product := range p {
		if product.Id == id {
			return product
		}
	}
	return Product{}
}

func indexOf(s any, data Products) int {
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