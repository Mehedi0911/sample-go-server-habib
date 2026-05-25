package main

import (
	"fmt"
	"sample-server/utils"
)

func main() {
	// cmd.Serve()

	// s := "fascism 2.0 already loaded"

	// bytesStr := []byte(s)

	// enc := base64.URLEncoding
	// enc = enc.WithPadding(base64.NoPadding)

	// encoded := enc.EncodeToString(bytesStr)
	// fmt.Println(encoded)

	jwt, err := utils.CreateJWT("my-secret", utils.Payload{
		Sub:         "1234567890",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "m@gmail.com",
		IsShopOwner: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(jwt)
}
