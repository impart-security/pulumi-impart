// Package main is th tfgen entrypoint for the Pulumi provider.
package main

import (
	"github.com/pulumi/pulumi-terraform-bridge/pf/tfgen"

	impart "github.com/impart-security/pulumi-impart/provider"
)

func main() {
	tfgen.Main("impart", impart.Provider())
}
