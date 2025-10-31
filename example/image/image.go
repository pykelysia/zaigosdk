package main

import (
	"fmt"

	"github.com/pykelysia/zaigosdk"
)

func main() {
	imageClient := zaigosdk.MustDefaultImageModel()
	res, err := imageClient.Chat("画一张可爱的小狗")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
