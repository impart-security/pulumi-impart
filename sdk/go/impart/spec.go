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

// Manage a specification.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/impart-security/pulumi-impart/sdk/go/impart"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Create a new specification
//			_, err := impart.NewSpec(ctx, "example", &impart.SpecArgs{
//				Name:       pulumi.String("spec_example"),
//				SourceFile: pulumi.String(fmt.Sprintf("%v/spec.yaml", path.Module)),
//				SourceHash: pulumi.String("<sha256 hash for the source_file content>"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// #!/bin/bash
//
// ```sh
// $ pulumi import impart:index/spec:Spec example "<id>"
// ```
type Spec struct {
	pulumi.CustomResourceState

	// The name for this specification.
	Name pulumi.StringOutput `pulumi:"name"`
	// The specification file.
	SourceFile pulumi.StringOutput `pulumi:"sourceFile"`
	// The specification source hash.
	SourceHash pulumi.StringPtrOutput `pulumi:"sourceHash"`
}

// NewSpec registers a new resource with the given unique name, arguments, and options.
func NewSpec(ctx *pulumi.Context,
	name string, args *SpecArgs, opts ...pulumi.ResourceOption) (*Spec, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.SourceFile == nil {
		return nil, errors.New("invalid value for required argument 'SourceFile'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Spec
	err := ctx.RegisterResource("impart:index/spec:Spec", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpec gets an existing Spec resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpec(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpecState, opts ...pulumi.ResourceOption) (*Spec, error) {
	var resource Spec
	err := ctx.ReadResource("impart:index/spec:Spec", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Spec resources.
type specState struct {
	// The name for this specification.
	Name *string `pulumi:"name"`
	// The specification file.
	SourceFile *string `pulumi:"sourceFile"`
	// The specification source hash.
	SourceHash *string `pulumi:"sourceHash"`
}

type SpecState struct {
	// The name for this specification.
	Name pulumi.StringPtrInput
	// The specification file.
	SourceFile pulumi.StringPtrInput
	// The specification source hash.
	SourceHash pulumi.StringPtrInput
}

func (SpecState) ElementType() reflect.Type {
	return reflect.TypeOf((*specState)(nil)).Elem()
}

type specArgs struct {
	// The name for this specification.
	Name string `pulumi:"name"`
	// The specification file.
	SourceFile string `pulumi:"sourceFile"`
	// The specification source hash.
	SourceHash *string `pulumi:"sourceHash"`
}

// The set of arguments for constructing a Spec resource.
type SpecArgs struct {
	// The name for this specification.
	Name pulumi.StringInput
	// The specification file.
	SourceFile pulumi.StringInput
	// The specification source hash.
	SourceHash pulumi.StringPtrInput
}

func (SpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*specArgs)(nil)).Elem()
}

type SpecInput interface {
	pulumi.Input

	ToSpecOutput() SpecOutput
	ToSpecOutputWithContext(ctx context.Context) SpecOutput
}

func (*Spec) ElementType() reflect.Type {
	return reflect.TypeOf((**Spec)(nil)).Elem()
}

func (i *Spec) ToSpecOutput() SpecOutput {
	return i.ToSpecOutputWithContext(context.Background())
}

func (i *Spec) ToSpecOutputWithContext(ctx context.Context) SpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpecOutput)
}

// SpecArrayInput is an input type that accepts SpecArray and SpecArrayOutput values.
// You can construct a concrete instance of `SpecArrayInput` via:
//
//	SpecArray{ SpecArgs{...} }
type SpecArrayInput interface {
	pulumi.Input

	ToSpecArrayOutput() SpecArrayOutput
	ToSpecArrayOutputWithContext(context.Context) SpecArrayOutput
}

type SpecArray []SpecInput

func (SpecArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Spec)(nil)).Elem()
}

func (i SpecArray) ToSpecArrayOutput() SpecArrayOutput {
	return i.ToSpecArrayOutputWithContext(context.Background())
}

func (i SpecArray) ToSpecArrayOutputWithContext(ctx context.Context) SpecArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpecArrayOutput)
}

// SpecMapInput is an input type that accepts SpecMap and SpecMapOutput values.
// You can construct a concrete instance of `SpecMapInput` via:
//
//	SpecMap{ "key": SpecArgs{...} }
type SpecMapInput interface {
	pulumi.Input

	ToSpecMapOutput() SpecMapOutput
	ToSpecMapOutputWithContext(context.Context) SpecMapOutput
}

type SpecMap map[string]SpecInput

func (SpecMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Spec)(nil)).Elem()
}

func (i SpecMap) ToSpecMapOutput() SpecMapOutput {
	return i.ToSpecMapOutputWithContext(context.Background())
}

func (i SpecMap) ToSpecMapOutputWithContext(ctx context.Context) SpecMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpecMapOutput)
}

type SpecOutput struct{ *pulumi.OutputState }

func (SpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Spec)(nil)).Elem()
}

func (o SpecOutput) ToSpecOutput() SpecOutput {
	return o
}

func (o SpecOutput) ToSpecOutputWithContext(ctx context.Context) SpecOutput {
	return o
}

// The name for this specification.
func (o SpecOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The specification file.
func (o SpecOutput) SourceFile() pulumi.StringOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringOutput { return v.SourceFile }).(pulumi.StringOutput)
}

// The specification source hash.
func (o SpecOutput) SourceHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Spec) pulumi.StringPtrOutput { return v.SourceHash }).(pulumi.StringPtrOutput)
}

type SpecArrayOutput struct{ *pulumi.OutputState }

func (SpecArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Spec)(nil)).Elem()
}

func (o SpecArrayOutput) ToSpecArrayOutput() SpecArrayOutput {
	return o
}

func (o SpecArrayOutput) ToSpecArrayOutputWithContext(ctx context.Context) SpecArrayOutput {
	return o
}

func (o SpecArrayOutput) Index(i pulumi.IntInput) SpecOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Spec {
		return vs[0].([]*Spec)[vs[1].(int)]
	}).(SpecOutput)
}

type SpecMapOutput struct{ *pulumi.OutputState }

func (SpecMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Spec)(nil)).Elem()
}

func (o SpecMapOutput) ToSpecMapOutput() SpecMapOutput {
	return o
}

func (o SpecMapOutput) ToSpecMapOutputWithContext(ctx context.Context) SpecMapOutput {
	return o
}

func (o SpecMapOutput) MapIndex(k pulumi.StringInput) SpecOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Spec {
		return vs[0].(map[string]*Spec)[vs[1].(string)]
	}).(SpecOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpecInput)(nil)).Elem(), &Spec{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpecArrayInput)(nil)).Elem(), SpecArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpecMapInput)(nil)).Elem(), SpecMap{})
	pulumi.RegisterOutputType(SpecOutput{})
	pulumi.RegisterOutputType(SpecArrayOutput{})
	pulumi.RegisterOutputType(SpecMapOutput{})
}
