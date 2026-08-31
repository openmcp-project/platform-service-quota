package main

import (
	"context"
	"fmt"
	"os"

	"github.com/openmcp-project/platform-service-quota/cmd/platform-service-quota/app"

	"github.com/openmcp-project/controller-utils/pkg/fips"
)

func main() {
	fips.Verify(context.Background())

	cmd := app.NewPlatformServiceQuotaCommand()

	if err := cmd.Execute(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
