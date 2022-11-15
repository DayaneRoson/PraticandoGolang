package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {

	r := func() {
		fmt.Println("pode até função, transformamos o go na M do js")
	}

	mapa := map[interface{}]interface{}{
		1:                "String",
		"olha-a-bagunça": float64(300),
		true:             123.4,
	}

	generica("String")
	generica(3247263847)
	generica(r)
	generica(mapa)
}
