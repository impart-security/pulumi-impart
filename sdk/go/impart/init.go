// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package impart

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/impart-security/pulumi-impart/sdk/go/impart/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "impart:index/apiBinding:ApiBinding":
		r = &ApiBinding{}
	case "impart:index/logBinding:LogBinding":
		r = &LogBinding{}
	case "impart:index/monitor:Monitor":
		r = &Monitor{}
	case "impart:index/notificationTemplate:NotificationTemplate":
		r = &NotificationTemplate{}
	case "impart:index/ruleScript:RuleScript":
		r = &RuleScript{}
	case "impart:index/ruleScriptDependencies:RuleScriptDependencies":
		r = &RuleScriptDependencies{}
	case "impart:index/spec:Spec":
		r = &Spec{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:impart" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"impart",
		"index/apiBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"impart",
		"index/logBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"impart",
		"index/monitor",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"impart",
		"index/notificationTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"impart",
		"index/ruleScript",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"impart",
		"index/ruleScriptDependencies",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"impart",
		"index/spec",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"impart",
		&pkg{version},
	)
}
