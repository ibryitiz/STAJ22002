package main

import "fmt"

type Student struct {
	id         int
	name       string
	grades     []Grade
	totalGrade float64
}

type Grade struct {
	id    int
	name  string
	quiz1 int
	vize  int
	quiz2 int
	work  int
	final int
}

func newGrade(id int, name string, quiz1 int, quiz2 int, vize int, work int, final int) Grade {
	return Grade{id: id, name: name, quiz1: quiz1, quiz2: quiz2, vize: vize, work: work, final: final}
}

// func newStudent(id int, name string, grades []Grade, totalGrade float64) Student {
// 	return Student{id: id, name: name, grades: grades, totalGrade: totalGrade}
// }

func (g Grade) calculateGrade() float64 {
	quiz1 := (g.quiz1 * 10) / 100
	quiz2 := (g.quiz2 * 10) / 100
	vize := (g.vize * 20) / 100
	work := (g.work * 10) / 100
	final := (g.final * 50) / 100

	average := quiz1 + quiz2 + vize + work + final
	return float64(average)
}

func main() {
	maps := map[string]*Student{
		"Öğrenci 1": {1, "İbrahim Yitiz", []Grade{}, 0.00},
		"Öğrenci 2": {2, "Rıdvan Bekleviç", []Grade{}, 0.00},
		"Öğrenci 3": {3, "Ahmet Ağzıgüzel", []Grade{}, 0.00},
		"Öğrenci 4": {4, "Yusuf Duman", []Grade{}, 0.00},
	}

	for {
		var studentControle int
		fmt.Println("\nÖğrenci seçiniz: ")
		fmt.Println("İbrahim için 1")
		fmt.Println("Rıdvan için 2")
		fmt.Println("Ahmet için 3")
		fmt.Println("Yusuf için 4")
		fmt.Println("Çıkış için 0")
		fmt.Scan(&studentControle)

		if studentControle == 0 {
			break
		}

		switch studentControle {
		case 1:
			handleStudent(maps["Öğrenci 1"])
		case 2:
			handleStudent(maps["Öğrenci 2"])
		case 3:
			handleStudent(maps["Öğrenci 3"])
		case 4:
			handleStudent(maps["Öğrenci 4"])
		default:
			fmt.Println("Geçersiz seçim.")
		}
	}
}

func handleStudent(student *Student) {
	fmt.Printf("%s için hangi işlemi yapmak istiyorsunuz?\n", student.name)
	fmt.Println("Ders ekleme : 1")
	fmt.Println("Genel ortalama görüntüleme : 2")
	fmt.Println("Dersleri görüntüle : 3")
	fmt.Println("Ana menüye dön : 0")

	var gradeControle int
	fmt.Scan(&gradeControle)

	switch gradeControle {
	case 1:
		addGrade(student)
	case 2:
		fmt.Printf("%s'in toplam ortalaması : %f\n", student.name, student.totalGrade)
	case 3:
		getFullGrade(student.grades)
	case 0:
		return
	default:
		fmt.Println("Geçersiz seçim.")
	}
}

func getNewGrade() Grade {
	var name string
	fmt.Println("Dersin adını giriniz : ")
	fmt.Scan(&name)
	var quiz1 int
	fmt.Println("Quiz 1 notunuzu giriniz : ")
	fmt.Scan(&quiz1)
	var quiz2 int
	fmt.Println("Quiz 2 notunuzu giriniz : ")
	fmt.Scan(&quiz2)
	var vize int
	fmt.Println("Vize notunuzu giriniz : ")
	fmt.Scan(&vize)
	var work int
	fmt.Println("Ödev notunuzu giriniz : ")
	fmt.Scan(&work)
	var final int
	fmt.Println("Final notunuzu giriniz : ")
	fmt.Scan(&final)
	return newGrade(1, name, quiz1, quiz2, vize, work, final)
}

func addGrade(student *Student) {
	newGrade := getNewGrade()
	student.grades = append(student.grades, newGrade)
	student.totalGrade = calculateTotalGrade(student.grades)
}

func getFullGrade(grades []Grade) {
	for _, value := range grades {
		fmt.Println("Ders id: ", value.id)
		fmt.Println("Ders adı: ", value.name)
		fmt.Println("Ders quiz 1 notu: ", value.quiz1)
		fmt.Println("Ders quiz 2 notu: ", value.quiz2)
		fmt.Println("Ders vize notu: ", value.vize)
		fmt.Println("Ders ödev notu: ", value.work)
		fmt.Println("Ders final notu: ", value.final)
	}
}

func calculateTotalGrade(grades []Grade) float64 {
	var total float64
	for _, v := range grades {
		total += v.calculateGrade()
	}
	if len(grades) > 0 {
		return total / float64(len(grades))
	}
	return 0.0
}
