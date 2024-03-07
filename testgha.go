package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/awnumar/memguard"
	"github.com/openpubkey/openpubkey/client"
	"github.com/openpubkey/openpubkey/client/providers"
	"github.com/openpubkey/openpubkey/util"
)

func testIdt2LB() {
	bodyb64 := []byte("eyJjb3VudCI6MTg1NywidmFsdWUiOiJleUowZVhBaU9pSktWMVFpTENKaGJHY2lPaUpTVXpJMU5pSXNJbmcxZENJNklraDVjVFJPUVZSQmFuTnVjVU0zYldSeWRFRm9hSEpEVWpKZlVTSXNJbXRwWkNJNklqRkdNa0ZDT0RNME1EUkRNRGhGUXpsRlFUQkNRams1UkVGRlJEQXlNVGcyUWpBNU1VUkNSalFpZlEuZXlKcWRHa2lPaUkwTjJKa01UUTJaaTA1WXpRM0xUUXpZbU10T0RRMk5TMDBZV0ppWTJNMk1UWTNaVEVpTENKemRXSWlPaUp5WlhCdk9rVjBhR0Z1U0dWcGJHMWhiaTluYUdGbGVIQXlPbkpsWmpweVpXWnpMMmhsWVdSekwyMWhhVzRpTENKaGRXUWlPaUpNU2twbVlXaEZOV05ETVVGblFWZHlUV3RWUkV3NE5XUXdiMU5UUW1OUU5rWktWbE4xYkhwdmFtUnpJaXdpY21WbUlqb2ljbVZtY3k5b1pXRmtjeTl0WVdsdUlpd2ljMmhoSWpvaU16VXpOekl5WXpreE4yRXpaamswT1RnNFlqZ3lObUk0TWpRd05XTmhNRFZtWldSa1lqRm1aU0lzSW5KbGNHOXphWFJ2Y25raU9pSkZkR2hoYmtobGFXeHRZVzR2WjJoaFpYaHdNaUlzSW5KbGNHOXphWFJ2Y25sZmIzZHVaWElpT2lKRmRHaGhia2hsYVd4dFlXNGlMQ0p5WlhCdmMybDBiM0o1WDI5M2JtVnlYMmxrSWpvaU1qYzBPREUwSWl3aWNuVnVYMmxrSWpvaU9ERTVNemczTlRJMU5TSXNJbkoxYmw5dWRXMWlaWElpT2lJeE15SXNJbkoxYmw5aGRIUmxiWEIwSWpvaU1TSXNJbkpsY0c5emFYUnZjbmxmZG1semFXSnBiR2wwZVNJNkluQjFZbXhwWXlJc0luSmxjRzl6YVhSdmNubGZhV1FpT2lJM05qZzNORGcwTmpJaUxDSmhZM1J2Y2w5cFpDSTZJakkzTkRneE5DSXNJbUZqZEc5eUlqb2lSWFJvWVc1SVpXbHNiV0Z1SWl3aWQyOXlhMlpzYjNjaU9pSkRTU0lzSW1obFlXUmZjbVZtSWpvaUlpd2lZbUZ6WlY5eVpXWWlPaUlpTENKbGRtVnVkRjl1WVcxbElqb2ljSFZ6YUNJc0luSmxabDl3Y205MFpXTjBaV1FpT2lKbVlXeHpaU0lzSW5KbFpsOTBlWEJsSWpvaVluSmhibU5vSWl3aWQyOXlhMlpzYjNkZmNtVm1Jam9pUlhSb1lXNUlaV2xzYldGdUwyZG9ZV1Y0Y0RJdkxtZHBkR2gxWWk5M2IzSnJabXh2ZDNNdmRHVnpkQzU1Yld4QWNtVm1jeTlvWldGa2N5OXRZV2x1SWl3aWQyOXlhMlpzYjNkZmMyaGhJam9pTXpVek56SXlZemt4TjJFelpqazBPVGc0WWpneU5tSTRNalF3TldOaE1EVm1aV1JrWWpGbVpTSXNJbXB2WWw5M2IzSnJabXh2ZDE5eVpXWWlPaUpGZEdoaGJraGxhV3h0WVc0dloyaGhaWGh3TWk4dVoybDBhSFZpTDNkdmNtdG1iRzkzY3k5MFpYTjBMbmx0YkVCeVpXWnpMMmhsWVdSekwyMWhhVzRpTENKcWIySmZkMjl5YTJac2IzZGZjMmhoSWpvaU16VXpOekl5WXpreE4yRXpaamswT1RnNFlqZ3lObUk0TWpRd05XTmhNRFZtWldSa1lqRm1aU0lzSW5KMWJtNWxjbDlsYm5acGNtOXViV1Z1ZENJNkltZHBkR2gxWWkxb2IzTjBaV1FpTENKcGMzTWlPaUpvZEhSd2N6b3ZMM1J2YTJWdUxtRmpkR2x2Ym5NdVoybDBhSFZpZFhObGNtTnZiblJsYm5RdVkyOXRJaXdpYm1KbUlqb3hOekE1T0RNNU9EWTVMQ0psZUhBaU9qRTNNRGs0TkRBM05qa3NJbWxoZENJNk1UY3dPVGcwTURRMk9YMC5ORkJaZjlwbm02RVVYVUhOWHF2ak9mSWRieW5kRkdpMEVIZWUtc2VFZzByWGlEd1BHbVVOdFJJRG8zYWtTV1g3MFBaVW10NHU1OHhJUFN0c0Y0Zm5Kd2FRbUF0YkpyTjRyYXJaSlpEd2g3Rmwxcm85YTBVeWx5aVduVjZlUUk3RnZGUkxOeGp5VXAxRC1YSDFlUG9MbFFKVzNwblU2MlAwVGx3dDdRUDdadlQ1Y25RSExHY0RSMW5oaEV0aXdjdm1uTkY1ay1fZVVvUUFXTWVpLWhVdXB5bXpqTUhRbVduMTFIQnI1bXZ6RzZfWFA3UjFLd1hRUUs3cGhOOG83QmhBeURGWUR4SHR6WHRnLXNpLXRYU1ZHQkJXTUdyQTFpRmE3X3BLM052XzFxSzUtUkgxZWtPRXJmelFrdVNmNEdMWG0wRW1INGdiUExLaW5TcFJNTm5aZ3cifQ==")
	rawBody, err := util.Base64Decode(bodyb64)
	if err != nil {
		panic(err)
	}

	var jwt struct {
		Value *memguard.LockedBuffer
	}
	err = json.Unmarshal(rawBody, &jwt)
	if err != nil {
		panic(err)
	}
	fmt.Println(jwt.Value)
}

func main() {
	// testIdt2LB()

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

	c, err := client.New(ghaOp, client.WithSignGQ(true))
	if err != nil {
		fmt.Printf("Error creating client")
		panic(err)
	}

	pkt, err := c.Auth(context.TODO())
	if err != nil {
		fmt.Printf("Error getting pkt")
		panic(err)
	}

	// verErr := ghaOp.Verifier().VerifyProvider(context.TODO(), pkt)
	verErr := client.VerifyPKToken(context.TODO(), pkt, ghaOp)
	if verErr != nil {
		fmt.Printf("Error verifying pktoken")
		panic(verErr)
	} else {
		fmt.Printf("Success verifying pktoken")

	}

}
