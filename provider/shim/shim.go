package shim

import (
	framework "github.com/hashicorp/terraform-plugin-framework/provider"
	provider "github.com/impart-security/terraform-provider-impart/internal/provider"
)

func NewProvider() framework.Provider {
	return provider.New("shim")()
}
