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
//			// Create a new external link
//			_, err := impart.NewExternalLink(ctx, "externalLink1", &impart.ExternalLinkArgs{
//				Description:     pulumi.String("A link to Datadog dashboard for client IP address"),
//				Entity:          pulumi.String("request"),
//				JsonPathElement: pulumi.String("$.client_ip.address"),
//				Name:            pulumi.String("Datadog client IP address"),
//				Url:             pulumi.String("https://app.datadoghq.com/dashboard/3tm-mpc-863?tpl_var_ClientIp=9.37.130.233"),
//				Vendor:          pulumi.String("Datadog"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ExternalLink struct {
	pulumi.CustomResourceState

	// The description of the external link.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The entity to which the links should be applied.
	Entity pulumi.StringOutput `pulumi:"entity"`
	// A JSONPath to the element for which this link should apply (e.g. $.client_ip.address).
	JsonPathElement pulumi.StringOutput `pulumi:"jsonPathElement"`
	// The name of the external link.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of spec IDs this external link applies to (empty means all).
	SpecIds pulumi.StringArrayOutput `pulumi:"specIds"`
	// The external URL template with JSONPath element variables.
	Url pulumi.StringOutput `pulumi:"url"`
	// The vendor for the external link.
	Vendor pulumi.StringOutput `pulumi:"vendor"`
}

// NewExternalLink registers a new resource with the given unique name, arguments, and options.
func NewExternalLink(ctx *pulumi.Context,
	name string, args *ExternalLinkArgs, opts ...pulumi.ResourceOption) (*ExternalLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Entity == nil {
		return nil, errors.New("invalid value for required argument 'Entity'")
	}
	if args.JsonPathElement == nil {
		return nil, errors.New("invalid value for required argument 'JsonPathElement'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.Vendor == nil {
		return nil, errors.New("invalid value for required argument 'Vendor'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ExternalLink
	err := ctx.RegisterResource("impart:index/externalLink:ExternalLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExternalLink gets an existing ExternalLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExternalLinkState, opts ...pulumi.ResourceOption) (*ExternalLink, error) {
	var resource ExternalLink
	err := ctx.ReadResource("impart:index/externalLink:ExternalLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExternalLink resources.
type externalLinkState struct {
	// The description of the external link.
	Description *string `pulumi:"description"`
	// The entity to which the links should be applied.
	Entity *string `pulumi:"entity"`
	// A JSONPath to the element for which this link should apply (e.g. $.client_ip.address).
	JsonPathElement *string `pulumi:"jsonPathElement"`
	// The name of the external link.
	Name *string `pulumi:"name"`
	// A list of spec IDs this external link applies to (empty means all).
	SpecIds []string `pulumi:"specIds"`
	// The external URL template with JSONPath element variables.
	Url *string `pulumi:"url"`
	// The vendor for the external link.
	Vendor *string `pulumi:"vendor"`
}

type ExternalLinkState struct {
	// The description of the external link.
	Description pulumi.StringPtrInput
	// The entity to which the links should be applied.
	Entity pulumi.StringPtrInput
	// A JSONPath to the element for which this link should apply (e.g. $.client_ip.address).
	JsonPathElement pulumi.StringPtrInput
	// The name of the external link.
	Name pulumi.StringPtrInput
	// A list of spec IDs this external link applies to (empty means all).
	SpecIds pulumi.StringArrayInput
	// The external URL template with JSONPath element variables.
	Url pulumi.StringPtrInput
	// The vendor for the external link.
	Vendor pulumi.StringPtrInput
}

func (ExternalLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*externalLinkState)(nil)).Elem()
}

type externalLinkArgs struct {
	// The description of the external link.
	Description *string `pulumi:"description"`
	// The entity to which the links should be applied.
	Entity string `pulumi:"entity"`
	// A JSONPath to the element for which this link should apply (e.g. $.client_ip.address).
	JsonPathElement string `pulumi:"jsonPathElement"`
	// The name of the external link.
	Name string `pulumi:"name"`
	// A list of spec IDs this external link applies to (empty means all).
	SpecIds []string `pulumi:"specIds"`
	// The external URL template with JSONPath element variables.
	Url string `pulumi:"url"`
	// The vendor for the external link.
	Vendor string `pulumi:"vendor"`
}

// The set of arguments for constructing a ExternalLink resource.
type ExternalLinkArgs struct {
	// The description of the external link.
	Description pulumi.StringPtrInput
	// The entity to which the links should be applied.
	Entity pulumi.StringInput
	// A JSONPath to the element for which this link should apply (e.g. $.client_ip.address).
	JsonPathElement pulumi.StringInput
	// The name of the external link.
	Name pulumi.StringInput
	// A list of spec IDs this external link applies to (empty means all).
	SpecIds pulumi.StringArrayInput
	// The external URL template with JSONPath element variables.
	Url pulumi.StringInput
	// The vendor for the external link.
	Vendor pulumi.StringInput
}

func (ExternalLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalLinkArgs)(nil)).Elem()
}

type ExternalLinkInput interface {
	pulumi.Input

	ToExternalLinkOutput() ExternalLinkOutput
	ToExternalLinkOutputWithContext(ctx context.Context) ExternalLinkOutput
}

func (*ExternalLink) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalLink)(nil)).Elem()
}

func (i *ExternalLink) ToExternalLinkOutput() ExternalLinkOutput {
	return i.ToExternalLinkOutputWithContext(context.Background())
}

func (i *ExternalLink) ToExternalLinkOutputWithContext(ctx context.Context) ExternalLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalLinkOutput)
}

// ExternalLinkArrayInput is an input type that accepts ExternalLinkArray and ExternalLinkArrayOutput values.
// You can construct a concrete instance of `ExternalLinkArrayInput` via:
//
//	ExternalLinkArray{ ExternalLinkArgs{...} }
type ExternalLinkArrayInput interface {
	pulumi.Input

	ToExternalLinkArrayOutput() ExternalLinkArrayOutput
	ToExternalLinkArrayOutputWithContext(context.Context) ExternalLinkArrayOutput
}

type ExternalLinkArray []ExternalLinkInput

func (ExternalLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalLink)(nil)).Elem()
}

func (i ExternalLinkArray) ToExternalLinkArrayOutput() ExternalLinkArrayOutput {
	return i.ToExternalLinkArrayOutputWithContext(context.Background())
}

func (i ExternalLinkArray) ToExternalLinkArrayOutputWithContext(ctx context.Context) ExternalLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalLinkArrayOutput)
}

// ExternalLinkMapInput is an input type that accepts ExternalLinkMap and ExternalLinkMapOutput values.
// You can construct a concrete instance of `ExternalLinkMapInput` via:
//
//	ExternalLinkMap{ "key": ExternalLinkArgs{...} }
type ExternalLinkMapInput interface {
	pulumi.Input

	ToExternalLinkMapOutput() ExternalLinkMapOutput
	ToExternalLinkMapOutputWithContext(context.Context) ExternalLinkMapOutput
}

type ExternalLinkMap map[string]ExternalLinkInput

func (ExternalLinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalLink)(nil)).Elem()
}

func (i ExternalLinkMap) ToExternalLinkMapOutput() ExternalLinkMapOutput {
	return i.ToExternalLinkMapOutputWithContext(context.Background())
}

func (i ExternalLinkMap) ToExternalLinkMapOutputWithContext(ctx context.Context) ExternalLinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalLinkMapOutput)
}

type ExternalLinkOutput struct{ *pulumi.OutputState }

func (ExternalLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalLink)(nil)).Elem()
}

func (o ExternalLinkOutput) ToExternalLinkOutput() ExternalLinkOutput {
	return o
}

func (o ExternalLinkOutput) ToExternalLinkOutputWithContext(ctx context.Context) ExternalLinkOutput {
	return o
}

// The description of the external link.
func (o ExternalLinkOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExternalLink) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The entity to which the links should be applied.
func (o ExternalLinkOutput) Entity() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalLink) pulumi.StringOutput { return v.Entity }).(pulumi.StringOutput)
}

// A JSONPath to the element for which this link should apply (e.g. $.client_ip.address).
func (o ExternalLinkOutput) JsonPathElement() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalLink) pulumi.StringOutput { return v.JsonPathElement }).(pulumi.StringOutput)
}

// The name of the external link.
func (o ExternalLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of spec IDs this external link applies to (empty means all).
func (o ExternalLinkOutput) SpecIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ExternalLink) pulumi.StringArrayOutput { return v.SpecIds }).(pulumi.StringArrayOutput)
}

// The external URL template with JSONPath element variables.
func (o ExternalLinkOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalLink) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The vendor for the external link.
func (o ExternalLinkOutput) Vendor() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalLink) pulumi.StringOutput { return v.Vendor }).(pulumi.StringOutput)
}

type ExternalLinkArrayOutput struct{ *pulumi.OutputState }

func (ExternalLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalLink)(nil)).Elem()
}

func (o ExternalLinkArrayOutput) ToExternalLinkArrayOutput() ExternalLinkArrayOutput {
	return o
}

func (o ExternalLinkArrayOutput) ToExternalLinkArrayOutputWithContext(ctx context.Context) ExternalLinkArrayOutput {
	return o
}

func (o ExternalLinkArrayOutput) Index(i pulumi.IntInput) ExternalLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExternalLink {
		return vs[0].([]*ExternalLink)[vs[1].(int)]
	}).(ExternalLinkOutput)
}

type ExternalLinkMapOutput struct{ *pulumi.OutputState }

func (ExternalLinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalLink)(nil)).Elem()
}

func (o ExternalLinkMapOutput) ToExternalLinkMapOutput() ExternalLinkMapOutput {
	return o
}

func (o ExternalLinkMapOutput) ToExternalLinkMapOutputWithContext(ctx context.Context) ExternalLinkMapOutput {
	return o
}

func (o ExternalLinkMapOutput) MapIndex(k pulumi.StringInput) ExternalLinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExternalLink {
		return vs[0].(map[string]*ExternalLink)[vs[1].(string)]
	}).(ExternalLinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalLinkInput)(nil)).Elem(), &ExternalLink{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalLinkArrayInput)(nil)).Elem(), ExternalLinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalLinkMapInput)(nil)).Elem(), ExternalLinkMap{})
	pulumi.RegisterOutputType(ExternalLinkOutput{})
	pulumi.RegisterOutputType(ExternalLinkArrayOutput{})
	pulumi.RegisterOutputType(ExternalLinkMapOutput{})
}
