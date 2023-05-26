// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package impart

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage a binding.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/impart-security/pulumi-impart/sdk/go/impart"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := impart.NewBinding(ctx, "example", &impart.BindingArgs{
//				Name:     pulumi.String("binding_example"),
//				Port:     pulumi.Int(443),
//				SpecId:   pulumi.Any(resource.Impart_spec.Example.Id),
//				Hostname: pulumi.String("example.com"),
//				BasePath: pulumi.String("/"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Binding struct {
	pulumi.CustomResourceState

	// The basePath for this binding.
	BasePath pulumi.StringOutput `pulumi:"basePath"`
	// The hostname for this binding.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The name for this binding.
	Name pulumi.StringOutput `pulumi:"name"`
	// The port for this binding.
	Port pulumi.IntOutput `pulumi:"port"`
	// The specification id.
	SpecId pulumi.StringOutput `pulumi:"specId"`
	// The upstreamOrigin for this binding.
	UpstreamOrigin pulumi.StringPtrOutput `pulumi:"upstreamOrigin"`
}

// NewBinding registers a new resource with the given unique name, arguments, and options.
func NewBinding(ctx *pulumi.Context,
	name string, args *BindingArgs, opts ...pulumi.ResourceOption) (*Binding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BasePath == nil {
		return nil, errors.New("invalid value for required argument 'BasePath'")
	}
	if args.Hostname == nil {
		return nil, errors.New("invalid value for required argument 'Hostname'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.SpecId == nil {
		return nil, errors.New("invalid value for required argument 'SpecId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Binding
	err := ctx.RegisterResource("impart:index/binding:Binding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBinding gets an existing Binding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BindingState, opts ...pulumi.ResourceOption) (*Binding, error) {
	var resource Binding
	err := ctx.ReadResource("impart:index/binding:Binding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Binding resources.
type bindingState struct {
	// The basePath for this binding.
	BasePath *string `pulumi:"basePath"`
	// The hostname for this binding.
	Hostname *string `pulumi:"hostname"`
	// The name for this binding.
	Name *string `pulumi:"name"`
	// The port for this binding.
	Port *int `pulumi:"port"`
	// The specification id.
	SpecId *string `pulumi:"specId"`
	// The upstreamOrigin for this binding.
	UpstreamOrigin *string `pulumi:"upstreamOrigin"`
}

type BindingState struct {
	// The basePath for this binding.
	BasePath pulumi.StringPtrInput
	// The hostname for this binding.
	Hostname pulumi.StringPtrInput
	// The name for this binding.
	Name pulumi.StringPtrInput
	// The port for this binding.
	Port pulumi.IntPtrInput
	// The specification id.
	SpecId pulumi.StringPtrInput
	// The upstreamOrigin for this binding.
	UpstreamOrigin pulumi.StringPtrInput
}

func (BindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*bindingState)(nil)).Elem()
}

type bindingArgs struct {
	// The basePath for this binding.
	BasePath string `pulumi:"basePath"`
	// The hostname for this binding.
	Hostname string `pulumi:"hostname"`
	// The name for this binding.
	Name string `pulumi:"name"`
	// The port for this binding.
	Port int `pulumi:"port"`
	// The specification id.
	SpecId string `pulumi:"specId"`
	// The upstreamOrigin for this binding.
	UpstreamOrigin *string `pulumi:"upstreamOrigin"`
}

// The set of arguments for constructing a Binding resource.
type BindingArgs struct {
	// The basePath for this binding.
	BasePath pulumi.StringInput
	// The hostname for this binding.
	Hostname pulumi.StringInput
	// The name for this binding.
	Name pulumi.StringInput
	// The port for this binding.
	Port pulumi.IntInput
	// The specification id.
	SpecId pulumi.StringInput
	// The upstreamOrigin for this binding.
	UpstreamOrigin pulumi.StringPtrInput
}

func (BindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bindingArgs)(nil)).Elem()
}

type BindingInput interface {
	pulumi.Input

	ToBindingOutput() BindingOutput
	ToBindingOutputWithContext(ctx context.Context) BindingOutput
}

func (*Binding) ElementType() reflect.Type {
	return reflect.TypeOf((**Binding)(nil)).Elem()
}

func (i *Binding) ToBindingOutput() BindingOutput {
	return i.ToBindingOutputWithContext(context.Background())
}

func (i *Binding) ToBindingOutputWithContext(ctx context.Context) BindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingOutput)
}

// BindingArrayInput is an input type that accepts BindingArray and BindingArrayOutput values.
// You can construct a concrete instance of `BindingArrayInput` via:
//
//	BindingArray{ BindingArgs{...} }
type BindingArrayInput interface {
	pulumi.Input

	ToBindingArrayOutput() BindingArrayOutput
	ToBindingArrayOutputWithContext(context.Context) BindingArrayOutput
}

type BindingArray []BindingInput

func (BindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Binding)(nil)).Elem()
}

func (i BindingArray) ToBindingArrayOutput() BindingArrayOutput {
	return i.ToBindingArrayOutputWithContext(context.Background())
}

func (i BindingArray) ToBindingArrayOutputWithContext(ctx context.Context) BindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingArrayOutput)
}

// BindingMapInput is an input type that accepts BindingMap and BindingMapOutput values.
// You can construct a concrete instance of `BindingMapInput` via:
//
//	BindingMap{ "key": BindingArgs{...} }
type BindingMapInput interface {
	pulumi.Input

	ToBindingMapOutput() BindingMapOutput
	ToBindingMapOutputWithContext(context.Context) BindingMapOutput
}

type BindingMap map[string]BindingInput

func (BindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Binding)(nil)).Elem()
}

func (i BindingMap) ToBindingMapOutput() BindingMapOutput {
	return i.ToBindingMapOutputWithContext(context.Background())
}

func (i BindingMap) ToBindingMapOutputWithContext(ctx context.Context) BindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingMapOutput)
}

type BindingOutput struct{ *pulumi.OutputState }

func (BindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Binding)(nil)).Elem()
}

func (o BindingOutput) ToBindingOutput() BindingOutput {
	return o
}

func (o BindingOutput) ToBindingOutputWithContext(ctx context.Context) BindingOutput {
	return o
}

// The basePath for this binding.
func (o BindingOutput) BasePath() pulumi.StringOutput {
	return o.ApplyT(func(v *Binding) pulumi.StringOutput { return v.BasePath }).(pulumi.StringOutput)
}

// The hostname for this binding.
func (o BindingOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *Binding) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// The name for this binding.
func (o BindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Binding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The port for this binding.
func (o BindingOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Binding) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// The specification id.
func (o BindingOutput) SpecId() pulumi.StringOutput {
	return o.ApplyT(func(v *Binding) pulumi.StringOutput { return v.SpecId }).(pulumi.StringOutput)
}

// The upstreamOrigin for this binding.
func (o BindingOutput) UpstreamOrigin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Binding) pulumi.StringPtrOutput { return v.UpstreamOrigin }).(pulumi.StringPtrOutput)
}

type BindingArrayOutput struct{ *pulumi.OutputState }

func (BindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Binding)(nil)).Elem()
}

func (o BindingArrayOutput) ToBindingArrayOutput() BindingArrayOutput {
	return o
}

func (o BindingArrayOutput) ToBindingArrayOutputWithContext(ctx context.Context) BindingArrayOutput {
	return o
}

func (o BindingArrayOutput) Index(i pulumi.IntInput) BindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Binding {
		return vs[0].([]*Binding)[vs[1].(int)]
	}).(BindingOutput)
}

type BindingMapOutput struct{ *pulumi.OutputState }

func (BindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Binding)(nil)).Elem()
}

func (o BindingMapOutput) ToBindingMapOutput() BindingMapOutput {
	return o
}

func (o BindingMapOutput) ToBindingMapOutputWithContext(ctx context.Context) BindingMapOutput {
	return o
}

func (o BindingMapOutput) MapIndex(k pulumi.StringInput) BindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Binding {
		return vs[0].(map[string]*Binding)[vs[1].(string)]
	}).(BindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BindingInput)(nil)).Elem(), &Binding{})
	pulumi.RegisterInputType(reflect.TypeOf((*BindingArrayInput)(nil)).Elem(), BindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BindingMapInput)(nil)).Elem(), BindingMap{})
	pulumi.RegisterOutputType(BindingOutput{})
	pulumi.RegisterOutputType(BindingArrayOutput{})
	pulumi.RegisterOutputType(BindingMapOutput{})
}