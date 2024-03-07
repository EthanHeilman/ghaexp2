package main

import (
	"context"
	"fmt"

	"github.com/openpubkey/openpubkey/client"
	"github.com/openpubkey/openpubkey/client/providers"
)

func main() {
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
		fmt.Printf("Error creating client")
		panic(err)
	}

	verErr := ghaOp.Verifier().VerifyProvider(context.TODO(), pkt)
	if verErr != nil {
		fmt.Printf("Error verifying pktoken")
		panic(verErr)
	}
}
