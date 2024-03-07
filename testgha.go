package main

import (
	"context"
	"fmt"

	"github.com/openpubkey/openpubkey/client"
	"github.com/openpubkey/openpubkey/client/providers"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Errorf("%+v", r))
			// fmt.Println(fmt.Sprintf("%+v", r))
		}
	}()

	ghaOp, err := providers.NewGithubOpFromEnvironment()
	if err != nil {
		fmt.Printf("Error creating provider")
		panic(err)
	}

	_ = ghaOp

	c, err := client.New(ghaOp)
	if err != nil {
		fmt.Printf("Error creating client")
		panic(err)
	}

	pkt, err := c.Auth(context.TODO())
	if err != nil {
		fmt.Printf("Error getting pkt")
		panic(err)
	}

	client.VerifyPKToken(context.TODO(), pkt, ghaOp)

	// verErr := ghaOp.Verifier().VerifyProvider(context.TODO(), pkt)
	// if verErr != nil {
	// 	fmt.Printf("Error verifying pktoken")
	// 	panic(verErr)
	// }
}
