package main

import "fmt"

func main() {

	//* Soru 1
	// Kullanıcıdan iki tamsayı (int) girişi alarak
	// bu sayıları toplayan ve sonucu ekrana yazdıran
	// bir Go programı yazın.

	// var a, b int
	// fmt.Println("Birinci sayıyı girin:")
	// fmt.Scan(&a)
	// fmt.Println("İkinci sayıyı girin:")
	// fmt.Scan(&b)

	// toplam := a + b
	// fmt.Println("Toplam:", toplam)

	//* Soru 2
	//Bir dizi (array) ve slice kullanarak kullanıcıdan
	//5 tamsayı girişi alıp bu sayıları
	//ters sırayla ekrana yazdıran bir Go programı yazın.

	// var numbers []int

	// var a, b, c int
	// fmt.Println("Birinci sayıyı giriniz:")
	// fmt.Scan(&a)
	// fmt.Println("İkinci sayıyı giriniz:")
	// fmt.Scan(&b)
	// fmt.Println("Üçüncü sayıyı giriniz:")
	// fmt.Scan(&c)

	// numbers = append(numbers, a, b, c)
	// fmt.Println(numbers)

	// var derece float64
	// derece = 38.6

	// var f float64 = (derece*1.8 + 32)

	// fmt.Printf("%f derece : %f fahrenheit", derece, f)

	// var numbers = []int{14, 18, 16, 17, 54, 25, 84, 64, 2}
	// fmt.Println(numbers)
	// toplam := 0
	// for index := 0; index < len(numbers); index++ {
	// 	fmt.Println(numbers[index])
	// 	toplam = toplam + numbers[index]
	// }
	// fmt.Printf("numberların toplamı : %d", toplam)

	// var tamsayi int
	// tamsayi = 1245

	// fmt.Println(tamsayi)
	// if tamsayi%2 == 0 {
	// 	fmt.Printf("%d sayısı çiftdir", tamsayi)
	// } else {
	// 	fmt.Printf("%d sayısı tekdir", tamsayi)
	// }

	// Öğrenci isimleri ve notları
	// studentGrades := make(map[string]int)
	// studentGrades["Ali"] = 85
	// studentGrades["Ayşe"] = 90
	// studentGrades["Mehmet"] = 45
	// studentGrades["Fatma"] = 72

	// // Notları slice olarak saklama
	// grades := []int{}
	// for _, grade := range studentGrades {
	// 	grades = append(grades, grade)
	// }

	// // Not ortalaması hesaplama
	// total := 0
	// for _, grade := range grades {
	// 	total += grade
	// }
	// average := float64(total) / float64(len(grades))

	// // Sonuçları ekrana yazdırma
	// fmt.Printf("Not Ortalaması: %.2f\n", average)
	// for student, grade := range studentGrades {
	// 	status := "Kaldı"
	// 	if grade >= 50 {
	// 		status = "Geçti"
	// 	}
	// 	fmt.Printf("%s: %d (%s)\n", student, grade, status)
	// }

	var studentAndNot = map[string]int{
		"İbrahim": 95,
		"Rıdvan":  78,
		"Ahmet":   56,
		"Yusuf":   15,
		"Zeynep":  47,
		"Rabiye":  56,
	}

	notlar := []int{}

	for _, not := range studentAndNot {
		notlar = append(notlar, not)
	}

	toplam := 0
	for i := 0; i < len(notlar); i++ {
		toplam = toplam + notlar[i]
	}

	ort := toplam / len(notlar)

	fmt.Printf("Öğrencilerin not ortalamsı : %d", ort)

	gecti := "Geçti"
	kaldi := "Kaldi"
	for _, not := range studentAndNot {
		if not >= 50 {
			fmt.Println(gecti)
		} else {
			fmt.Println(kaldi)
		}
	}

}
