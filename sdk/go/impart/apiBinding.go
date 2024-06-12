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

// Manage an api binding.
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
//			// Create a new api binding
//			_, err := impart.NewApiBinding(ctx, "example", &impart.ApiBindingArgs{
//				Name:     pulumi.String("api_binding_example"),
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
type ApiBinding struct {
	pulumi.CustomResourceState

	// The basePath for this api binding.
	BasePath pulumi.StringOutput `pulumi:"basePath"`
	// The disabled for this api binding.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// The forwardedFor for this api binding.
	ForwardedFors pulumi.StringArrayOutput `pulumi:"forwardedFors"`
	// The forwardedHost for this api binding.
	ForwardedHosts pulumi.StringArrayOutput `pulumi:"forwardedHosts"`
	// The forwardedId for this api binding.
	ForwardedIds pulumi.StringArrayOutput `pulumi:"forwardedIds"`
	// The forwardedProto for this api binding.
	ForwardedProtos pulumi.StringArrayOutput `pulumi:"forwardedProtos"`
	// The hops for this api binding.
	Hops pulumi.IntPtrOutput `pulumi:"hops"`
	// The hostname for this api binding.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The name for this api binding.
	Name pulumi.StringOutput `pulumi:"name"`
	// The port for this api binding.
	Port pulumi.IntOutput `pulumi:"port"`
	// The specification id.
	SpecId pulumi.StringOutput `pulumi:"specId"`
	// The upstreamOrigin for this api binding.
	UpstreamOrigin pulumi.StringPtrOutput `pulumi:"upstreamOrigin"`
	// The useForwarded for this api binding.
	UseForwarded pulumi.BoolPtrOutput `pulumi:"useForwarded"`
}

// NewApiBinding registers a new resource with the given unique name, arguments, and options.
func NewApiBinding(ctx *pulumi.Context,
	name string, args *ApiBindingArgs, opts ...pulumi.ResourceOption) (*ApiBinding, error) {
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
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApiBinding
	err := ctx.RegisterResource("impart:index/apiBinding:ApiBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiBinding gets an existing ApiBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiBindingState, opts ...pulumi.ResourceOption) (*ApiBinding, error) {
	var resource ApiBinding
	err := ctx.ReadResource("impart:index/apiBinding:ApiBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiBinding resources.
type apiBindingState struct {
	// The basePath for this api binding.
	BasePath *string `pulumi:"basePath"`
	// The disabled for this api binding.
	Disabled *bool `pulumi:"disabled"`
	// The forwardedFor for this api binding.
	ForwardedFors []string `pulumi:"forwardedFors"`
	// The forwardedHost for this api binding.
	ForwardedHosts []string `pulumi:"forwardedHosts"`
	// The forwardedId for this api binding.
	ForwardedIds []string `pulumi:"forwardedIds"`
	// The forwardedProto for this api binding.
	ForwardedProtos []string `pulumi:"forwardedProtos"`
	// The hops for this api binding.
	Hops *int `pulumi:"hops"`
	// The hostname for this api binding.
	Hostname *string `pulumi:"hostname"`
	// The name for this api binding.
	Name *string `pulumi:"name"`
	// The port for this api binding.
	Port *int `pulumi:"port"`
	// The specification id.
	SpecId *string `pulumi:"specId"`
	// The upstreamOrigin for this api binding.
	UpstreamOrigin *string `pulumi:"upstreamOrigin"`
	// The useForwarded for this api binding.
	UseForwarded *bool `pulumi:"useForwarded"`
}

type ApiBindingState struct {
	// The basePath for this api binding.
	BasePath pulumi.StringPtrInput
	// The disabled for this api binding.
	Disabled pulumi.BoolPtrInput
	// The forwardedFor for this api binding.
	ForwardedFors pulumi.StringArrayInput
	// The forwardedHost for this api binding.
	ForwardedHosts pulumi.StringArrayInput
	// The forwardedId for this api binding.
	ForwardedIds pulumi.StringArrayInput
	// The forwardedProto for this api binding.
	ForwardedProtos pulumi.StringArrayInput
	// The hops for this api binding.
	Hops pulumi.IntPtrInput
	// The hostname for this api binding.
	Hostname pulumi.StringPtrInput
	// The name for this api binding.
	Name pulumi.StringPtrInput
	// The port for this api binding.
	Port pulumi.IntPtrInput
	// The specification id.
	SpecId pulumi.StringPtrInput
	// The upstreamOrigin for this api binding.
	UpstreamOrigin pulumi.StringPtrInput
	// The useForwarded for this api binding.
	UseForwarded pulumi.BoolPtrInput
}

func (ApiBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiBindingState)(nil)).Elem()
}

type apiBindingArgs struct {
	// The basePath for this api binding.
	BasePath string `pulumi:"basePath"`
	// The disabled for this api binding.
	Disabled *bool `pulumi:"disabled"`
	// The forwardedFor for this api binding.
	ForwardedFors []string `pulumi:"forwardedFors"`
	// The forwardedHost for this api binding.
	ForwardedHosts []string `pulumi:"forwardedHosts"`
	// The forwardedId for this api binding.
	ForwardedIds []string `pulumi:"forwardedIds"`
	// The forwardedProto for this api binding.
	ForwardedProtos []string `pulumi:"forwardedProtos"`
	// The hops for this api binding.
	Hops *int `pulumi:"hops"`
	// The hostname for this api binding.
	Hostname string `pulumi:"hostname"`
	// The name for this api binding.
	Name string `pulumi:"name"`
	// The port for this api binding.
	Port int `pulumi:"port"`
	// The specification id.
	SpecId string `pulumi:"specId"`
	// The upstreamOrigin for this api binding.
	UpstreamOrigin *string `pulumi:"upstreamOrigin"`
	// The useForwarded for this api binding.
	UseForwarded *bool `pulumi:"useForwarded"`
}

// The set of arguments for constructing a ApiBinding resource.
type ApiBindingArgs struct {
	// The basePath for this api binding.
	BasePath pulumi.StringInput
	// The disabled for this api binding.
	Disabled pulumi.BoolPtrInput
	// The forwardedFor for this api binding.
	ForwardedFors pulumi.StringArrayInput
	// The forwardedHost for this api binding.
	ForwardedHosts pulumi.StringArrayInput
	// The forwardedId for this api binding.
	ForwardedIds pulumi.StringArrayInput
	// The forwardedProto for this api binding.
	ForwardedProtos pulumi.StringArrayInput
	// The hops for this api binding.
	Hops pulumi.IntPtrInput
	// The hostname for this api binding.
	Hostname pulumi.StringInput
	// The name for this api binding.
	Name pulumi.StringInput
	// The port for this api binding.
	Port pulumi.IntInput
	// The specification id.
	SpecId pulumi.StringInput
	// The upstreamOrigin for this api binding.
	UpstreamOrigin pulumi.StringPtrInput
	// The useForwarded for this api binding.
	UseForwarded pulumi.BoolPtrInput
}

func (ApiBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiBindingArgs)(nil)).Elem()
}

type ApiBindingInput interface {
	pulumi.Input

	ToApiBindingOutput() ApiBindingOutput
	ToApiBindingOutputWithContext(ctx context.Context) ApiBindingOutput
}

func (*ApiBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiBinding)(nil)).Elem()
}

func (i *ApiBinding) ToApiBindingOutput() ApiBindingOutput {
	return i.ToApiBindingOutputWithContext(context.Background())
}

func (i *ApiBinding) ToApiBindingOutputWithContext(ctx context.Context) ApiBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiBindingOutput)
}

// ApiBindingArrayInput is an input type that accepts ApiBindingArray and ApiBindingArrayOutput values.
// You can construct a concrete instance of `ApiBindingArrayInput` via:
//
//	ApiBindingArray{ ApiBindingArgs{...} }
type ApiBindingArrayInput interface {
	pulumi.Input

	ToApiBindingArrayOutput() ApiBindingArrayOutput
	ToApiBindingArrayOutputWithContext(context.Context) ApiBindingArrayOutput
}

type ApiBindingArray []ApiBindingInput

func (ApiBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiBinding)(nil)).Elem()
}

func (i ApiBindingArray) ToApiBindingArrayOutput() ApiBindingArrayOutput {
	return i.ToApiBindingArrayOutputWithContext(context.Background())
}

func (i ApiBindingArray) ToApiBindingArrayOutputWithContext(ctx context.Context) ApiBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiBindingArrayOutput)
}

// ApiBindingMapInput is an input type that accepts ApiBindingMap and ApiBindingMapOutput values.
// You can construct a concrete instance of `ApiBindingMapInput` via:
//
//	ApiBindingMap{ "key": ApiBindingArgs{...} }
type ApiBindingMapInput interface {
	pulumi.Input

	ToApiBindingMapOutput() ApiBindingMapOutput
	ToApiBindingMapOutputWithContext(context.Context) ApiBindingMapOutput
}

type ApiBindingMap map[string]ApiBindingInput

func (ApiBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiBinding)(nil)).Elem()
}

func (i ApiBindingMap) ToApiBindingMapOutput() ApiBindingMapOutput {
	return i.ToApiBindingMapOutputWithContext(context.Background())
}

func (i ApiBindingMap) ToApiBindingMapOutputWithContext(ctx context.Context) ApiBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiBindingMapOutput)
}

type ApiBindingOutput struct{ *pulumi.OutputState }

func (ApiBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiBinding)(nil)).Elem()
}

func (o ApiBindingOutput) ToApiBindingOutput() ApiBindingOutput {
	return o
}

func (o ApiBindingOutput) ToApiBindingOutputWithContext(ctx context.Context) ApiBindingOutput {
	return o
}

// The basePath for this api binding.
func (o ApiBindingOutput) BasePath() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiBinding) pulumi.StringOutput { return v.BasePath }).(pulumi.StringOutput)
}

// The disabled for this api binding.
func (o ApiBindingOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiBinding) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// The forwardedFor for this api binding.
func (o ApiBindingOutput) ForwardedFors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiBinding) pulumi.StringArrayOutput { return v.ForwardedFors }).(pulumi.StringArrayOutput)
}

// The forwardedHost for this api binding.
func (o ApiBindingOutput) ForwardedHosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiBinding) pulumi.StringArrayOutput { return v.ForwardedHosts }).(pulumi.StringArrayOutput)
}

// The forwardedId for this api binding.
func (o ApiBindingOutput) ForwardedIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiBinding) pulumi.StringArrayOutput { return v.ForwardedIds }).(pulumi.StringArrayOutput)
}

// The forwardedProto for this api binding.
func (o ApiBindingOutput) ForwardedProtos() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiBinding) pulumi.StringArrayOutput { return v.ForwardedProtos }).(pulumi.StringArrayOutput)
}

// The hops for this api binding.
func (o ApiBindingOutput) Hops() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApiBinding) pulumi.IntPtrOutput { return v.Hops }).(pulumi.IntPtrOutput)
}

// The hostname for this api binding.
func (o ApiBindingOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiBinding) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// The name for this api binding.
func (o ApiBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The port for this api binding.
func (o ApiBindingOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *ApiBinding) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// The specification id.
func (o ApiBindingOutput) SpecId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiBinding) pulumi.StringOutput { return v.SpecId }).(pulumi.StringOutput)
}

// The upstreamOrigin for this api binding.
func (o ApiBindingOutput) UpstreamOrigin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiBinding) pulumi.StringPtrOutput { return v.UpstreamOrigin }).(pulumi.StringPtrOutput)
}

// The useForwarded for this api binding.
func (o ApiBindingOutput) UseForwarded() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiBinding) pulumi.BoolPtrOutput { return v.UseForwarded }).(pulumi.BoolPtrOutput)
}

type ApiBindingArrayOutput struct{ *pulumi.OutputState }

func (ApiBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiBinding)(nil)).Elem()
}

func (o ApiBindingArrayOutput) ToApiBindingArrayOutput() ApiBindingArrayOutput {
	return o
}

func (o ApiBindingArrayOutput) ToApiBindingArrayOutputWithContext(ctx context.Context) ApiBindingArrayOutput {
	return o
}

func (o ApiBindingArrayOutput) Index(i pulumi.IntInput) ApiBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApiBinding {
		return vs[0].([]*ApiBinding)[vs[1].(int)]
	}).(ApiBindingOutput)
}

type ApiBindingMapOutput struct{ *pulumi.OutputState }

func (ApiBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiBinding)(nil)).Elem()
}

func (o ApiBindingMapOutput) ToApiBindingMapOutput() ApiBindingMapOutput {
	return o
}

func (o ApiBindingMapOutput) ToApiBindingMapOutputWithContext(ctx context.Context) ApiBindingMapOutput {
	return o
}

func (o ApiBindingMapOutput) MapIndex(k pulumi.StringInput) ApiBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApiBinding {
		return vs[0].(map[string]*ApiBinding)[vs[1].(string)]
	}).(ApiBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiBindingInput)(nil)).Elem(), &ApiBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiBindingArrayInput)(nil)).Elem(), ApiBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiBindingMapInput)(nil)).Elem(), ApiBindingMap{})
	pulumi.RegisterOutputType(ApiBindingOutput{})
	pulumi.RegisterOutputType(ApiBindingArrayOutput{})
	pulumi.RegisterOutputType(ApiBindingMapOutput{})
}
