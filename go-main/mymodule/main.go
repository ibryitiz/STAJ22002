package main

import (
	"fmt"
	"mymodule/helper"
	"mymodule/helper/rest"
)

func main() {
	fmt.Println("Hello world!")
	helper.Helper1()
	rest.Rest1()
	// ! aşağıdaki gibi fonksiyonlar küçük harfle başlarsa ulaşılamaz olur
	// ! büyük harf olursa ulaşılır buna package scope denir
	//* rest.rest2()
}
