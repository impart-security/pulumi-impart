// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package impart

import (
	"context"
	"reflect"

	"errors"
	"github.com/impart-security/pulumi-impart/sdk/go/impart/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage rule script dependencies. There should only ever be one instance of this resource in a workspace at once, because it manages rule script dependencies at an organization level.
type RuleScriptDependencies struct {
	pulumi.CustomResourceState

	// An array of rule scripts and the other ids of the rules they depend on before executing.
	Dependencies RuleScriptDependenciesDependencyArrayOutput `pulumi:"dependencies"`
}

// NewRuleScriptDependencies registers a new resource with the given unique name, arguments, and options.
func NewRuleScriptDependencies(ctx *pulumi.Context,
	name string, args *RuleScriptDependenciesArgs, opts ...pulumi.ResourceOption) (*RuleScriptDependencies, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dependencies == nil {
		return nil, errors.New("invalid value for required argument 'Dependencies'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RuleScriptDependencies
	err := ctx.RegisterResource("impart:index/ruleScriptDependencies:RuleScriptDependencies", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuleScriptDependencies gets an existing RuleScriptDependencies resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuleScriptDependencies(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleScriptDependenciesState, opts ...pulumi.ResourceOption) (*RuleScriptDependencies, error) {
	var resource RuleScriptDependencies
	err := ctx.ReadResource("impart:index/ruleScriptDependencies:RuleScriptDependencies", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuleScriptDependencies resources.
type ruleScriptDependenciesState struct {
	// An array of rule scripts and the other ids of the rules they depend on before executing.
	Dependencies []RuleScriptDependenciesDependency `pulumi:"dependencies"`
}

type RuleScriptDependenciesState struct {
	// An array of rule scripts and the other ids of the rules they depend on before executing.
	Dependencies RuleScriptDependenciesDependencyArrayInput
}

func (RuleScriptDependenciesState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleScriptDependenciesState)(nil)).Elem()
}

type ruleScriptDependenciesArgs struct {
	// An array of rule scripts and the other ids of the rules they depend on before executing.
	Dependencies []RuleScriptDependenciesDependency `pulumi:"dependencies"`
}

// The set of arguments for constructing a RuleScriptDependencies resource.
type RuleScriptDependenciesArgs struct {
	// An array of rule scripts and the other ids of the rules they depend on before executing.
	Dependencies RuleScriptDependenciesDependencyArrayInput
}

func (RuleScriptDependenciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleScriptDependenciesArgs)(nil)).Elem()
}

type RuleScriptDependenciesInput interface {
	pulumi.Input

	ToRuleScriptDependenciesOutput() RuleScriptDependenciesOutput
	ToRuleScriptDependenciesOutputWithContext(ctx context.Context) RuleScriptDependenciesOutput
}

func (*RuleScriptDependencies) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleScriptDependencies)(nil)).Elem()
}

func (i *RuleScriptDependencies) ToRuleScriptDependenciesOutput() RuleScriptDependenciesOutput {
	return i.ToRuleScriptDependenciesOutputWithContext(context.Background())
}

func (i *RuleScriptDependencies) ToRuleScriptDependenciesOutputWithContext(ctx context.Context) RuleScriptDependenciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleScriptDependenciesOutput)
}

// RuleScriptDependenciesArrayInput is an input type that accepts RuleScriptDependenciesArray and RuleScriptDependenciesArrayOutput values.
// You can construct a concrete instance of `RuleScriptDependenciesArrayInput` via:
//
//	RuleScriptDependenciesArray{ RuleScriptDependenciesArgs{...} }
type RuleScriptDependenciesArrayInput interface {
	pulumi.Input

	ToRuleScriptDependenciesArrayOutput() RuleScriptDependenciesArrayOutput
	ToRuleScriptDependenciesArrayOutputWithContext(context.Context) RuleScriptDependenciesArrayOutput
}

type RuleScriptDependenciesArray []RuleScriptDependenciesInput

func (RuleScriptDependenciesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuleScriptDependencies)(nil)).Elem()
}

func (i RuleScriptDependenciesArray) ToRuleScriptDependenciesArrayOutput() RuleScriptDependenciesArrayOutput {
	return i.ToRuleScriptDependenciesArrayOutputWithContext(context.Background())
}

func (i RuleScriptDependenciesArray) ToRuleScriptDependenciesArrayOutputWithContext(ctx context.Context) RuleScriptDependenciesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleScriptDependenciesArrayOutput)
}

// RuleScriptDependenciesMapInput is an input type that accepts RuleScriptDependenciesMap and RuleScriptDependenciesMapOutput values.
// You can construct a concrete instance of `RuleScriptDependenciesMapInput` via:
//
//	RuleScriptDependenciesMap{ "key": RuleScriptDependenciesArgs{...} }
type RuleScriptDependenciesMapInput interface {
	pulumi.Input

	ToRuleScriptDependenciesMapOutput() RuleScriptDependenciesMapOutput
	ToRuleScriptDependenciesMapOutputWithContext(context.Context) RuleScriptDependenciesMapOutput
}

type RuleScriptDependenciesMap map[string]RuleScriptDependenciesInput

func (RuleScriptDependenciesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuleScriptDependencies)(nil)).Elem()
}

func (i RuleScriptDependenciesMap) ToRuleScriptDependenciesMapOutput() RuleScriptDependenciesMapOutput {
	return i.ToRuleScriptDependenciesMapOutputWithContext(context.Background())
}

func (i RuleScriptDependenciesMap) ToRuleScriptDependenciesMapOutputWithContext(ctx context.Context) RuleScriptDependenciesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleScriptDependenciesMapOutput)
}

type RuleScriptDependenciesOutput struct{ *pulumi.OutputState }

func (RuleScriptDependenciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleScriptDependencies)(nil)).Elem()
}

func (o RuleScriptDependenciesOutput) ToRuleScriptDependenciesOutput() RuleScriptDependenciesOutput {
	return o
}

func (o RuleScriptDependenciesOutput) ToRuleScriptDependenciesOutputWithContext(ctx context.Context) RuleScriptDependenciesOutput {
	return o
}

// An array of rule scripts and the other ids of the rules they depend on before executing.
func (o RuleScriptDependenciesOutput) Dependencies() RuleScriptDependenciesDependencyArrayOutput {
	return o.ApplyT(func(v *RuleScriptDependencies) RuleScriptDependenciesDependencyArrayOutput { return v.Dependencies }).(RuleScriptDependenciesDependencyArrayOutput)
}

type RuleScriptDependenciesArrayOutput struct{ *pulumi.OutputState }

func (RuleScriptDependenciesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuleScriptDependencies)(nil)).Elem()
}

func (o RuleScriptDependenciesArrayOutput) ToRuleScriptDependenciesArrayOutput() RuleScriptDependenciesArrayOutput {
	return o
}

func (o RuleScriptDependenciesArrayOutput) ToRuleScriptDependenciesArrayOutputWithContext(ctx context.Context) RuleScriptDependenciesArrayOutput {
	return o
}

func (o RuleScriptDependenciesArrayOutput) Index(i pulumi.IntInput) RuleScriptDependenciesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RuleScriptDependencies {
		return vs[0].([]*RuleScriptDependencies)[vs[1].(int)]
	}).(RuleScriptDependenciesOutput)
}

type RuleScriptDependenciesMapOutput struct{ *pulumi.OutputState }

func (RuleScriptDependenciesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuleScriptDependencies)(nil)).Elem()
}

func (o RuleScriptDependenciesMapOutput) ToRuleScriptDependenciesMapOutput() RuleScriptDependenciesMapOutput {
	return o
}

func (o RuleScriptDependenciesMapOutput) ToRuleScriptDependenciesMapOutputWithContext(ctx context.Context) RuleScriptDependenciesMapOutput {
	return o
}

func (o RuleScriptDependenciesMapOutput) MapIndex(k pulumi.StringInput) RuleScriptDependenciesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RuleScriptDependencies {
		return vs[0].(map[string]*RuleScriptDependencies)[vs[1].(string)]
	}).(RuleScriptDependenciesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleScriptDependenciesInput)(nil)).Elem(), &RuleScriptDependencies{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleScriptDependenciesArrayInput)(nil)).Elem(), RuleScriptDependenciesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleScriptDependenciesMapInput)(nil)).Elem(), RuleScriptDependenciesMap{})
	pulumi.RegisterOutputType(RuleScriptDependenciesOutput{})
	pulumi.RegisterOutputType(RuleScriptDependenciesArrayOutput{})
	pulumi.RegisterOutputType(RuleScriptDependenciesMapOutput{})
}
