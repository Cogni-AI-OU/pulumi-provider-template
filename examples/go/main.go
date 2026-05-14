package main

import (
	template "github.com/Cogni-AI-OU/pulumi-provider-template/sdk/go/pulumi-provider-template"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myRandomResource, err := boilerplate.NewRandom(ctx, "myRandomResource", &boilerplate.RandomArgs{
			Length: pulumi.Int(24),
		})
		if err != nil {
			return err
		}
		_, err = boilerplate.NewRandomComponent(ctx, "myRandomComponent", &boilerplate.RandomComponentArgs{
			Length: pulumi.Int(24),
		})
		if err != nil {
			return err
		}
		ctx.Export("output", pulumi.StringMap{
			"value": myRandomResource.Result,
		})
		return nil
	})
}
