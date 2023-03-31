package main

import "fmt"

func main() {
	var idMap map[int]string
	println(idMap)
	fmt.Println(idMap)

	idMap = make(map[int]string)
	println(idMap)
	fmt.Println(idMap)

	tickers := map[string]string{
		"GOOG" : "Google",
		"MSFT" : "MicroSoft",
		"FB" : "FaceBook",
	}
	println(tickers)
	fmt.Println(tickers)

	tickers["SMSG"] = "Samsung"
	println(tickers)
	fmt.Println(tickers)

	noData := tickers["LG"]
	println(noData)
	fmt.Println(noData)

	

	
}