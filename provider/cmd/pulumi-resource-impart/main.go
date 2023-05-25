package main

import (
	"context"
	_ "embed"

	"github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"

	impart "github.com/impart-security/pulumi-impart/provider"
)

//go:embed schema.json
var schema []byte

func main() {
	meta := tfbridge.ProviderMetadata{PackageSchema: schema}
	tfbridge.Main(context.Background(), "impart", impart.Provider(), meta)
}
