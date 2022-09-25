package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Product struct {
	id uint8
	name string
	price int
}

type Products []Product

func main() {
	//List of product
	productList := Products{
		{1, "Benih Lele", 50000},
		{2, "Pakan Lele Cap Menara", 25000},
		{3, "Probiotik A", 75000},
		{4, "Probiotik Nila B", 10000},
		{5, "Pakan Nila", 20000},
		{6, "Benih Nila", 20000},
		{7, "Cupang", 5000},
		{8, "Benih Nila", 30000},
		{9, "Benih Cupang", 10000},
		{10, "Probiotik B", 10000},
	}
	//Find the highes price product (hpp) on the list 
	_, hpp, _ := productList.findByPrice("max")
	fmt.Println("=====================================================")
	fmt.Printf("Product with the highest price: \n %s - Rp. %d \n\n", hpp[0].name, hpp[0].price)
	
	//Find the lowest price product (lpp) on the list
	_, lpp, _ := productList.findByPrice("min")
	fmt.Println("=====================================================")
	fmt.Printf("Product with the lowest price: \n %s - Rp. %d \n\n", lpp[0].name,lpp[0].price)

	// Find product with price of specified pricetag (fpp)
	priceToFind := 10000 //edit to change priceTag
	_, fpp, err := productList.findByPrice(priceToFind)
	if len(fpp) != 0 {
		fmt.Println("=====================================================")
		fmt.Printf("Product list with the price of Rp. %d :\n", priceToFind)
		for _, product := range fpp {
			fmt.Printf("%s - Rp. %d \n", product.name, product.price)
		}
		fmt.Printf("Total products: %d Product(s) \n\n", len(fpp))
	} else {
		fmt.Println("=====================================================")
		fmt.Printf("%s \n\n", err)
	}

	//buyOptimizer to get maximum products with specified amount of money
	money := 200000
	moneyChanges, totalPrice, cart := productList.buyOptimizer(money)
	fmt.Println("=====================================================")
	fmt.Printf("List of Product bought with Rp. %d \n", money)
	for _, product := range cart {
		fmt.Printf("%s - Rp. %d \n", product.name, product.price)
	}
	fmt.Printf("Total Products: %d Product(s) \n Payment: Rp. %d \n Total Price: Rp. %d \n Changes: Rp. %d \n", len(cart), money, totalPrice, moneyChanges)
}

func (pl Products) findByPrice(method any) (string, Products, string){
	var products Products
	var status, err string
	if _, isString  := method.(string); isString {
		switch s := strings.ToLower(method.(string)); s {
		case "max":
			matchProduct := pl[0]
			for i := 1; i < len(pl); i++ {
				if pl[i].price > matchProduct.price {
					matchProduct = pl[i]
				}
			}
			status = "success"
			products = append(products, matchProduct)
		case "min":
			matchProduct := pl[0]
			for i := 1; i < len(pl); i++ {
				if pl[i].price < matchProduct.price {
					matchProduct = pl[i]
				}
			}
			status = "success"
			products = append(products, matchProduct)
		default:
			status = "fail"
			err = "Method invalid. Should use max, min, or integer"
		}
	}

	if _, isInteger := method.(int); isInteger {
		matchProduct := Products{}
		// find the product with defined price
		for _, product := range pl {
			if product.price == method {
				matchProduct = append(matchProduct, product)
			}
		}

		//check if the matchProduct is empty or not
		if len(matchProduct) != 0 {
			status = "success"
			products = matchProduct
		} else {
			status = "fail"
			err = "Product with price of Rp." + strconv.Itoa(method.(int)) + " not found"
		}
	}
	return status, products, err
}

func (pl Products) buyOptimizer(money int) (int,int, Products){
	var cart Products
	var totalPrice int
	sortedProducts := pl
	sort.Slice(sortedProducts, func(i, j int) bool {
		return pl[i].price < pl[j].price
	})
	for _, product := range sortedProducts{
		if money > 0 {
			if product.price < money {
				cart = append(cart, product)
				totalPrice = totalPrice + product.price
				money = money - product.price
			} else {
				continue
			}
		} 
	}
	return money, totalPrice, cart
}