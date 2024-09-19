// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {

// 	result, err := divide(10, 0)

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(result)
// 	}

// }

// func divide(x int, y int) (int, error) {
// 	if y == 0 {
// 		return 0, errors.New("Payda 0 olamaz!")
// 	} else {
// 		return x / y, nil
// 	}
// }

package main

import (
	"fmt"
)

type Product struct {
	id    int
	name  string
	price int
}

type ProductService struct {
}

func (productService *ProductService) Add(product Product) error {
	if len(product.name) == 0 {
		//return errors.New("Ürün adı boş olamaz!")
		// ! YERİNE
		return ValidationError{text: "Ürün ismi boş olamaz!", code: "1001"}
	}
	if product.price < 0 {
		//return errors.New("Ürün fiyatı 0 dan büyük olmalıdır!")
		// ! YERİNE
		return ValidationError{text: "Ürün fiyatı 0 dan büyük olmalıdır!", code: "1002"}
	}

	fmt.Println("Ürün eklendi.")
	return nil
}

func main() {

	productService := ProductService{}

	err := productService.Add(Product{id: 1, name: "Elma", price: -10})

	if err != nil {
		fmt.Println(err)
	}
}

type ValidationError struct {
	code string
	text string
}

func (validationError ValidationError) Error() string {
	return fmt.Sprintf("Hata : %s , Kod : %s", validationError.text, validationError.code)
}
