package impart

//github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge.ShimProvider

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/google/uuid"
	pf "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/impart-security/pulumi-impart/provider/pkg/version"
	"github.com/impart-security/terraform-provider-impart/shim"
)

// all of the impart token components used below.
const (
	mainPkg = "impart"
	mainMod = "index"
)

//go:embed cmd/pulumi-resource-impart/bridge-metadata.json
var bridgeMetadata []byte

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
	return tfbridge.ProviderInfo{
		P:                 pf.ShimProvider(shim.NewProvider()),
		Name:              "impart",
		DisplayName:       "Impart Security",
		Keywords:          []string{"pulumi", "impart", "category/infrastructure"},
		Version:           version.Version,
		License:           "Apache-2.0",
		Homepage:          "https://www.impart.security/",
		GitHubOrg:         "impart-security",
		LogoURL:           "https://console.impartsecurity.net/logo-blue-black.svg",
		Repository:        "https://github.com/impart-security/pulumi-impart",
		PluginDownloadURL: "github://api.github.com/impart-security",
		Publisher:         "Impart Security",
		Description:       "A Pulumi package for creating and managing Impart resources.",
		Config: map[string]*tfbridge.SchemaInfo{
			"endpoint": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"IMPART_ENDPOINT"},
				},
			},
			"token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"IMPART_TOKEN"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"impart_spec":                           {Tok: makeResource(mainMod, "Spec")},
			"impart_api_binding":                    {Tok: makeResource(mainMod, "ApiBinding")},
			"impart_log_binding":                    {Tok: makeResource(mainMod, "LogBinding")},
			"impart_list":                           {Tok: makeResource(mainMod, "List")},
			"impart_rule_script":                    {Tok: makeResource(mainMod, "RuleScript")},
			"impart_rule":                           {Tok: makeResource(mainMod, "Rule")},
			"impart_rule_test_case":                 {Tok: makeResource(mainMod, "RuleTestCase")},
			"impart_notification_template":          {Tok: makeResource(mainMod, "NotificationTemplate")},
			"impart_monitor":                        {Tok: makeResource(mainMod, "Monitor")},
			"impart_rule_client_identifier":         {Tok: makeResource(mainMod, "RuleClientIdentifier")},
			"impart_rule_client_identifier_storage": {Tok: makeResource(mainMod, "RuleClientIdentifierStorage")},
			"impart_core_rule_config":               {Tok: makeResource(mainMod, "CoreRuleConfig")},
			"impart_label": {
				Tok: makeResource(mainMod, "Label"),
				ComputeID: func(_ context.Context, state resource.PropertyMap) (resource.ID, error) {
					if id, ok := state["slug"]; ok {
						return resource.ID(id.StringValue()), nil
					}
					return "", errors.New("slug attribute not found")
				},
			},
			"impart_tag_metadata": {
				Tok: makeResource(mainMod, "TagMetadata"),
				ComputeID: func(_ context.Context, state resource.PropertyMap) (resource.ID, error) {
					if id, ok := state["name"]; ok {
						return resource.ID(id.StringValue()), nil
					}
					return "", errors.New("name attribute not found")
				},
			},
			"impart_rule_script_dependencies": {
				Tok:       makeResource(mainMod, "RuleScriptDependencies"),
				ComputeID: computeRuleScriptDependenciesID,
			},
			"impart_rule_dependencies": {
				Tok:       makeResource(mainMod, "RuleDependencies"),
				ComputeID: computeRuleScriptDependenciesID,
			},
			"impart_external_link": {Tok: makeResource(mainMod, "ExternalLink")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"impart_spec":      {Tok: makeDataSource(mainMod, "GetSpec")},
			"impart_connector": {Tok: makeDataSource(mainMod, "GetConnector")},
		},
		MetadataInfo: tfbridge.NewProviderMetadata(bridgeMetadata),
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/impart-security/pulumi-%s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			PackageName: "@impart-security/pulumi-impart",
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			// Overlay: &tfbridge.OverlayInfo{},
			TypeScriptVersion: "^5.4.3",
		},
	}
}

func computeRuleScriptDependenciesID(_ context.Context, state resource.PropertyMap) (resource.ID, error) {
	return resource.ID(uuid.New().String()), nil
}
