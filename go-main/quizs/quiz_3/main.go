package main

import "fmt"

type Product struct {
	name     string
	price    int
	quantity int
}

func newProduct(name string, price int, quantity int) Product {
	return Product{name, price, quantity}
}

func buyProduct(maps map[string]Product, productNumber int) {
	var productKey string
	switch productNumber {
	case 1:
		productKey = "laptop"
	case 2:
		productKey = "phone"
	case 3:
		productKey = "mouse"
	case 4:
		productKey = "keyboard"
	default:
		fmt.Println("maalesef bu ürün mağazamızda bulunmamaktadır.")
		return
	}

	product := maps[productKey]
	if product.quantity == 0 {
		fmt.Println("maalesef stokta ürün kalmadı")
		return
	}

	product.quantity -= 1
	maps[productKey] = product
	fmt.Printf("%s için kalan stok: %d\n", product.name, product.quantity)
}

func totalPrice(maps map[string]Product) int {
	total := 0
	for _, product := range maps {
		productPrice := product.price * product.quantity
		total += productPrice
	}
	return total
}

func main() {
	products := make(map[string]Product)

	products["laptop"] = newProduct("Macbook Pro", 45000, 3)
	products["phone"] = newProduct("iPhone", 65000, 2)
	products["mouse"] = newProduct("Logitech", 800, 10)
	products["keyboard"] = newProduct("Monster", 2500, 5)

	buyProduct(products, 2)
	buyProduct(products, 2)
	buyProduct(products, 2)

	buyProduct(products, 3)
	buyProduct(products, 3)

	buyProduct(products, 4)
	buyProduct(products, 4)

	fmt.Println("----------------------")
	fmt.Println("Mağazadaki ürünlerin toplam fiyatı: ")
	result := totalPrice(products)
	fmt.Println(result)
}
