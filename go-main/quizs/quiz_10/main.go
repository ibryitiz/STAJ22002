//! 1. SORU

//* Kendi veri tipimizi tanımlama
//* ve buna 10 değerini atama işlemi

/*package main

import "fmt"

type myType int

func main() {

	var x myType
	x = 10
	fmt.Printf("%T , %v", x, x)

}*/

//! 2. SORU

//* x = 20 değerini oluştur ve y değişkenine
//* x 'in adresini ver sonra y üzerinden x 'i 40 yap

/*package main

import "fmt"

func main() {
	x := 20

	y := &x

	*y = 40

	fmt.Println(x)
}*/

//! 3. SORU

//* Rectangle struct oluştur
//* ve bunun gösterme , alan , çevre
//* fonksiyonlarını yaz

/*
package main

import "fmt"

type rectangle struct {
	tall, short int
}

func (r rectangle) display() {
	fmt.Printf("Uzun Kenar: %d - ", r.tall)
	fmt.Printf("Kısa Kenar: %d", r.short)

}

func (r rectangle) area() {
	area := r.short * r.tall
	fmt.Printf("Alan = %d", area)
}

func (r rectangle) circumference() {
	circum := (r.short + r.tall) * 2
	fmt.Printf("Çevre = %d", circum)
}

func main() {

	r1 := rectangle{5, 10}
	fmt.Println()

	r1.display()
	fmt.Println()

	r1.area()
	fmt.Println()

	r1.circumference()

}
*/

//! Example

package main

import "fmt"

type family struct {
	name      string
	age       int
	isMarried bool
}

func showFamily() []family {
	baba := family{
		name:      "İlhami",
		age:       45,
		isMarried: true,
	}

	anne := family{
		name:      "Neşe",
		age:       40,
		isMarried: true,
	}

	abi := family{
		name: "İbrahim",
		age:  23,
	}

	kücükAbi := family{
		"Murselat",
		21,
		false,
	}

	kardes := family{
		"Ali",
		17,
		false,
	}

	var enKucuk family

	enKucuk.name = "Zeynep"
	enKucuk.age = 10
	enKucuk.isMarried = false

	return []family{baba, anne, abi, kücükAbi, kardes, enKucuk}
}

func main() {

	families := showFamily()

	for _, value := range families {
		fmt.Println(value)
	}

	fmt.Println("------  YA DA  ------")

	//?  YA DA

	for i := 0; i < len(families); i++ {
		fmt.Printf("%v , %T\n", families[i], families[i])
	}
}
