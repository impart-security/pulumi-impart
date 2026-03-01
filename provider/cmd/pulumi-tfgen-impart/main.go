// Package main is th tfgen entrypoint for the Pulumi provider.
package main

import (
	impart "github.com/impart-security/pulumi-impart/provider"
	"github.com/pulumi/pulumi-terraform-bridge/pf/tfgen"
)

func main() {
	tfgen.Main("impart", impart.Provider())
}
