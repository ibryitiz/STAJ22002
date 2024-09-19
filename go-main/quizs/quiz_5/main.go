package main

import "fmt"

func main() {
	mathProgram(15, 5)
	toplama, çıkarma, çarpma := mathProgram2(15, 5)

	fmt.Printf("Toplama : %d , Çıkarma : %d , Çarpma : %d", toplama, çıkarma, çarpma)

	var kısaSınav1, vize, kısaSınav2, ödev, final int
	fmt.Println(("NOT HESAPLAMA UYGULAMASINA HOŞ GELDİNİZ"))
	fmt.Println("Kısa sınav 1 notunuzu giriniz:")
	fmt.Scan(&kısaSınav1)
	fmt.Println("Vize notunuzu giriniz:")
	fmt.Scan(&vize)
	fmt.Println("Kısa sınav 1 notunuzu giriniz:")
	fmt.Scan(&kısaSınav2)
	fmt.Println("Ödev notunuzu giriniz:")
	fmt.Scan(&ödev)
	fmt.Println("Final notunuzu giriniz:")
	fmt.Scan(&final)

	result := ((kısaSınav1 * 5) / 100) + ((vize * 20) / 100) + ((kısaSınav2 * 8) / 100) + ((ödev * 17) / 100) + (final*50)/100

	if result < 50 {
		fmt.Printf("Ortalamanız : %d  , Maalesef , Kaldınız...", result)
	} else {
		fmt.Printf("Ortalamanız : %d , Geçtiniz , Tebrikler...", result)
	}

}

func mathProgram(num1 int, num2 int) {

	fmt.Printf("Toplama : %d , Çıkarma : %d , Çarpma : %d", num1+num2, num1-num2, num1*num2)

}

func mathProgram2(num1 int, num2 int) (int, int, int) {
	return num1 + num2, num1 - num2, num1 * num2
}
