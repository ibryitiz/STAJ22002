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

func newStudent(id int, name string, grades []Grade, totalGrade float64) Student {
	return Student{id: id, name: name, grades: grades, totalGrade: totalGrade}
}

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
	var maps = map[string]Student{
		"Öğrenci 1": newStudent(1, "İbrahim Yitiz", []Grade{}, 0.00),
		"Öğrenci 2": newStudent(2, "Rıdvan Bekleviç", []Grade{}, 0.00),
		"Öğrenci 3": newStudent(3, "Ahmet Ağzıgüzel", []Grade{}, 0.00),
		"Öğrenci 4": newStudent(4, "Yusuf Duman", []Grade{}, 0.00),
	}

	var studentControle int
	fmt.Println("Öğrenci seçiniz: ")
	fmt.Println("İbrahim için 1: ")
	fmt.Println("Rıdvan için 2: ")
	fmt.Println("Ahmet için 3: ")
	fmt.Println("Yusuf için 4: ")
	fmt.Scan(&studentControle)

	switch studentControle {
	case 1:
		fmt.Println("İbrahim için hangi işlemi yapmak istiyorsunuz?")
		fmt.Println("Ders ekleme : 1")
		fmt.Println("Genel ortalama görüntüleme : 2")
		fmt.Println("Dersleri görüntüle : 3")
		var gradeControle int
		fmt.Scan(&gradeControle)
		switch gradeControle {
		case 1:
			addGrade(maps["Öğrenci 1"].grades)
			calculateTotalGrade(maps["Öğrenci 1"])
		case 2:
			totalGrade := maps["Öğrenci 1"].totalGrade
			fmt.Printf("İbrahim Yitiz 'in toplam ortalaması : %f", totalGrade)
		case 3:
			getFullGrade(maps["Öğrenci 1"].grades)
		}
	case 2:
		fmt.Println("Rıdvan için hangi işlemi yapmak istiyorsunuz?")
		fmt.Println("Ders ekleme : 1")
		fmt.Println("Genel ortalama görüntüleme : 2")
		fmt.Println("Dersleri görüntüle : 3")
		var gradeControle int
		fmt.Scan(&gradeControle)
		switch gradeControle {
		case 1:
			addGrade(maps["Öğrenci 2"].grades)
			calculateTotalGrade(maps["Öğrenci 2"])
		case 2:
			totalGrade := maps["Öğrenci 2"].totalGrade
			fmt.Printf("Rıdvan Bekleviç'in toplam ortalaması : %f", totalGrade)
		case 3:
			getFullGrade(maps["Öğrenci 2"].grades)
		}

	case 3:
		fmt.Println("Ahmet için hangi işlemi yapmak istiyorsunuz?")
		fmt.Println("Ders ekleme : 1")
		fmt.Println("Genel ortalama görüntüleme : 2")
		fmt.Println("Dersleri görüntüle : 3")
		var gradeControle int
		fmt.Scan(&gradeControle)
		switch gradeControle {
		case 1:
			addGrade(maps["Öğrenci 3"].grades)
			calculateTotalGrade(maps["Öğrenci 3"])
		case 2:
			totalGrade := maps["Öğrenci 3"].totalGrade
			fmt.Printf("Ahmet Ağzıgüzel'in toplam ortalaması : %f", totalGrade)
		case 3:
			getFullGrade(maps["Öğrenci 3"].grades)
		}
	case 4:
		fmt.Println("Yusuf için hangi işlemi yapmak istiyorsunuz?")
		fmt.Println("Ders ekleme : 1")
		fmt.Println("Genel ortalama görüntüleme : 2")
		fmt.Println("Dersleri görüntüle : 3")
		var gradeControle int
		fmt.Scan(&gradeControle)
		switch gradeControle {
		case 1:
			addGrade(maps["Öğrenci 4"].grades)
			calculateTotalGrade(maps["Öğrenci 4"])
		case 2:
			totalGrade := maps["Öğrenci 4"].totalGrade
			fmt.Printf("Yusuf Duman'in toplam ortalaması : %f", totalGrade)
		case 3:
			getFullGrade(maps["Öğrenci 4"].grades)
		}

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

func addGrade(grades []Grade) {
	newGrade := getNewGrade()
	newGrades := append(grades, newGrade)
	fmt.Println(newGrades)
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

func calculateTotalGrade(student Student) {
	for _, v := range student.grades {
		student.totalGrade += v.calculateGrade()
	}
}
