package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	result, err := evenNum(11)

	if err != nil {

		fmt.Println(err)

	} else {

		fmt.Println("Girdiğiniz sayı : ", result)
	}

	result2, err2 := karekokAlma(-10)

	if err2 == nil {
		fmt.Println("Karekökü : ", result2)
	} else {
		fmt.Println(err2)
	}
}

func evenNum(num int) (int, error) {
	if num%2 != 0 {

		return 0, errors.New("HATA: Girdiğiniz sayı çift değil!")

	}

	return num, nil

}

func karekokAlma(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("Hata: Sıfırdan küçük sayı girilemez!")
	}
	return math.Sqrt(num), nil
}
