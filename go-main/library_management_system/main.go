package main

import "fmt"

type Book struct {
	bookID   int
	title    string
	author   string
	quantity int
}

type Member struct {
	name     string
	memberID int
	books    []Book
}

func newBook(bookID int, title string, author string, quantity int) Book {
	return Book{bookID, title, author, quantity}
}

func newMember(name string, memberID int) Member {
	return Member{name, memberID, []Book{}}
}

func borrowBook(members map[int]Member, books map[int]Book, memberID int, bookID int) {
	member, memberExists := members[memberID]
	book, bookExists := books[bookID]

	if !memberExists {
		fmt.Println("Üye bulunamadı.")
		return
	}

	if !bookExists {
		fmt.Println("Kitap bulunamadı.")
		return
	}

	if book.quantity == 0 {
		fmt.Println("Kitap stokta yok.")
		return
	}

	// Kitap stok miktarını azaltın
	book.quantity -= 1
	books[bookID] = book

	// Üyenin kitaplar listesine ekleyin
	member.books = append(member.books, book)
	members[memberID] = member

	fmt.Printf("%s kitabı %s üyesine başarıyla kiralandı.\n", book.title, member.name)
}
func returnBook(members map[int]Member, books map[int]Book, memberID int, bookID int) {
	member, memberExists := members[memberID]
	book, bookExists := books[bookID]

	if !memberExists {
		fmt.Println("Üye bulunamadı.")
		return
	}

	if !bookExists {
		fmt.Println("Kitap bulunamadı.")
		return
	}

	// Üyenin kiralanan kitaplar listesinden kitabı çıkarın
	for i, borrowedBook := range member.books {
		if borrowedBook.bookID == bookID {
			member.books = append(member.books[:i], member.books[i+1:]...)
			break
		}
	}

	// Kitap stok miktarını artırın
	book.quantity += 1
	books[bookID] = book

	members[memberID] = member

	fmt.Printf("%s kitabı %s üyesi tarafından iade edildi.\n", book.title, member.name)
}

func listBooks(books map[int]Book) {
	fmt.Println("Kütüphanedeki Kitaplar:")
	for _, book := range books {
		fmt.Printf("Kitap: %s, Yazar: %s, Stok: %d\n", book.title, book.author, book.quantity)
	}
}

func listMembers(members map[int]Member) {
	fmt.Println("Üyeler ve Kiraladıkları Kitaplar:")
	for _, member := range members {
		fmt.Printf("Üye: %s\n", member.name)
		for _, book := range member.books {
			fmt.Printf("\tKitap: %s, Yazar: %s\n", book.title, book.author)
		}
	}
}

func main() {
	books := make(map[int]Book)

	books[1] = newBook(1, "Şeker Portakalı", "Jose Mauro De Vasconcelos", 3)
	books[2] = newBook(2, "Hayvan Çifliği", "George Orwell", 4)
	books[3] = newBook(3, "1984", "George Orwell", 2)

	members := make(map[int]Member)

	members[1] = newMember("İbrahim", 1)
	members[2] = newMember("Ali", 2)

	for {
		fmt.Println("1. Kitap Kirala")
		fmt.Println("2. Kitap İade Et")
		fmt.Println("3. Kitapları Listele")
		fmt.Println("4. Üyeleri Listele")
		fmt.Println("5. Çıkış")
		fmt.Print("Seçiminizi yapın: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var memberID int
			var bookID int
			fmt.Print("Üye ID'sini girin: ")
			fmt.Scan(&memberID)
			fmt.Print("Kitap ID'sini girin: ")
			fmt.Scan(&bookID)
			borrowBook(members, books, memberID, bookID)
		case 2:
			var memberID int
			var bookID int
			fmt.Print("Üye ID'sini girin: ")
			fmt.Scan(&memberID)
			fmt.Print("Kitap ID'sini girin: ")
			fmt.Scan(&bookID)
			returnBook(members, books, memberID, bookID)
		case 3:
			listBooks(books)
		case 4:
			listMembers(members)
		case 5:
			return
		default:
			fmt.Println("Geçersiz seçim. Lütfen tekrar deneyin.")
		}
	}

	//! yukardaki for yapılmadan önce kontroller
	//* Kitap kiralama işlemleri
	//borrowBook(members, books, 1, 1)
	//borrowBook(members, books, 2, 2)
	//borrowBook(members, books, 1, 3)
	//borrowBook(members, books, 1, 3) // Bu işlem stokta kitap olmadığı için başarısız olacak
	//borrowBook(members, books, 1, 3) // Bu işlem stokta kitap olmadığı için başarısız olacak

	//* Kitap iade işlemleri
	//returnBook(members, books, 1, 1)
	//returnBook(members, books, 2, 2)
	//returnBook(members, books, 1, 3)

	//* Kitapları listeleme
	//listBooks(books)

	//* Üyeleri ve kiraladıkları kitapları listeleme
	//listMembers(members)
}
