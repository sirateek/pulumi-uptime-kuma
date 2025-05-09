package main

import (
	"github.com/sirateek/pulumi-uptime-kuma/sdk/go/uptime-kuma"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myRandomResource, err := uptime-kuma.NewRandom(ctx, "myRandomResource", &uptime-kuma.RandomArgs{
			Length: pulumi.Int(24),
		})
		if err != nil {
			return err
		}
		ctx.Export("output", map[string]interface{}{
			"value": myRandomResource.Result,
		})
		return nil
	})
}
