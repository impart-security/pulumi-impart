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
type RuleTestCase struct {
	pulumi.CustomResourceState

	// The assertions of the test case.
	Assertions RuleTestCaseAssertionArrayOutput `pulumi:"assertions"`
	// The description of the test case.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The applied labels.
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// The messages of the test case.
	Messages RuleTestCaseMessageArrayOutput `pulumi:"messages"`
	// The name of the test case.
	Name pulumi.StringOutput `pulumi:"name"`
	// Sets if test case required to pass on update.
	Required pulumi.BoolPtrOutput `pulumi:"required"`
}

// NewRuleTestCase registers a new resource with the given unique name, arguments, and options.
func NewRuleTestCase(ctx *pulumi.Context,
	name string, args *RuleTestCaseArgs, opts ...pulumi.ResourceOption) (*RuleTestCase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Messages == nil {
		return nil, errors.New("invalid value for required argument 'Messages'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RuleTestCase
	err := ctx.RegisterResource("impart:index/ruleTestCase:RuleTestCase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuleTestCase gets an existing RuleTestCase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuleTestCase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleTestCaseState, opts ...pulumi.ResourceOption) (*RuleTestCase, error) {
	var resource RuleTestCase
	err := ctx.ReadResource("impart:index/ruleTestCase:RuleTestCase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuleTestCase resources.
type ruleTestCaseState struct {
	// The assertions of the test case.
	Assertions []RuleTestCaseAssertion `pulumi:"assertions"`
	// The description of the test case.
	Description *string `pulumi:"description"`
	// The applied labels.
	Labels []string `pulumi:"labels"`
	// The messages of the test case.
	Messages []RuleTestCaseMessage `pulumi:"messages"`
	// The name of the test case.
	Name *string `pulumi:"name"`
	// Sets if test case required to pass on update.
	Required *bool `pulumi:"required"`
}

type RuleTestCaseState struct {
	// The assertions of the test case.
	Assertions RuleTestCaseAssertionArrayInput
	// The description of the test case.
	Description pulumi.StringPtrInput
	// The applied labels.
	Labels pulumi.StringArrayInput
	// The messages of the test case.
	Messages RuleTestCaseMessageArrayInput
	// The name of the test case.
	Name pulumi.StringPtrInput
	// Sets if test case required to pass on update.
	Required pulumi.BoolPtrInput
}

func (RuleTestCaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleTestCaseState)(nil)).Elem()
}

type ruleTestCaseArgs struct {
	// The assertions of the test case.
	Assertions []RuleTestCaseAssertion `pulumi:"assertions"`
	// The description of the test case.
	Description *string `pulumi:"description"`
	// The applied labels.
	Labels []string `pulumi:"labels"`
	// The messages of the test case.
	Messages []RuleTestCaseMessage `pulumi:"messages"`
	// The name of the test case.
	Name string `pulumi:"name"`
	// Sets if test case required to pass on update.
	Required *bool `pulumi:"required"`
}

// The set of arguments for constructing a RuleTestCase resource.
type RuleTestCaseArgs struct {
	// The assertions of the test case.
	Assertions RuleTestCaseAssertionArrayInput
	// The description of the test case.
	Description pulumi.StringPtrInput
	// The applied labels.
	Labels pulumi.StringArrayInput
	// The messages of the test case.
	Messages RuleTestCaseMessageArrayInput
	// The name of the test case.
	Name pulumi.StringInput
	// Sets if test case required to pass on update.
	Required pulumi.BoolPtrInput
}

func (RuleTestCaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleTestCaseArgs)(nil)).Elem()
}

type RuleTestCaseInput interface {
	pulumi.Input

	ToRuleTestCaseOutput() RuleTestCaseOutput
	ToRuleTestCaseOutputWithContext(ctx context.Context) RuleTestCaseOutput
}

func (*RuleTestCase) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleTestCase)(nil)).Elem()
}

func (i *RuleTestCase) ToRuleTestCaseOutput() RuleTestCaseOutput {
	return i.ToRuleTestCaseOutputWithContext(context.Background())
}

func (i *RuleTestCase) ToRuleTestCaseOutputWithContext(ctx context.Context) RuleTestCaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleTestCaseOutput)
}

// RuleTestCaseArrayInput is an input type that accepts RuleTestCaseArray and RuleTestCaseArrayOutput values.
// You can construct a concrete instance of `RuleTestCaseArrayInput` via:
//
//	RuleTestCaseArray{ RuleTestCaseArgs{...} }
type RuleTestCaseArrayInput interface {
	pulumi.Input

	ToRuleTestCaseArrayOutput() RuleTestCaseArrayOutput
	ToRuleTestCaseArrayOutputWithContext(context.Context) RuleTestCaseArrayOutput
}

type RuleTestCaseArray []RuleTestCaseInput

func (RuleTestCaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuleTestCase)(nil)).Elem()
}

func (i RuleTestCaseArray) ToRuleTestCaseArrayOutput() RuleTestCaseArrayOutput {
	return i.ToRuleTestCaseArrayOutputWithContext(context.Background())
}

func (i RuleTestCaseArray) ToRuleTestCaseArrayOutputWithContext(ctx context.Context) RuleTestCaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleTestCaseArrayOutput)
}

// RuleTestCaseMapInput is an input type that accepts RuleTestCaseMap and RuleTestCaseMapOutput values.
// You can construct a concrete instance of `RuleTestCaseMapInput` via:
//
//	RuleTestCaseMap{ "key": RuleTestCaseArgs{...} }
type RuleTestCaseMapInput interface {
	pulumi.Input

	ToRuleTestCaseMapOutput() RuleTestCaseMapOutput
	ToRuleTestCaseMapOutputWithContext(context.Context) RuleTestCaseMapOutput
}

type RuleTestCaseMap map[string]RuleTestCaseInput

func (RuleTestCaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuleTestCase)(nil)).Elem()
}

func (i RuleTestCaseMap) ToRuleTestCaseMapOutput() RuleTestCaseMapOutput {
	return i.ToRuleTestCaseMapOutputWithContext(context.Background())
}

func (i RuleTestCaseMap) ToRuleTestCaseMapOutputWithContext(ctx context.Context) RuleTestCaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleTestCaseMapOutput)
}

type RuleTestCaseOutput struct{ *pulumi.OutputState }

func (RuleTestCaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleTestCase)(nil)).Elem()
}

func (o RuleTestCaseOutput) ToRuleTestCaseOutput() RuleTestCaseOutput {
	return o
}

func (o RuleTestCaseOutput) ToRuleTestCaseOutputWithContext(ctx context.Context) RuleTestCaseOutput {
	return o
}

// The assertions of the test case.
func (o RuleTestCaseOutput) Assertions() RuleTestCaseAssertionArrayOutput {
	return o.ApplyT(func(v *RuleTestCase) RuleTestCaseAssertionArrayOutput { return v.Assertions }).(RuleTestCaseAssertionArrayOutput)
}

// The description of the test case.
func (o RuleTestCaseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuleTestCase) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The applied labels.
func (o RuleTestCaseOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RuleTestCase) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

// The messages of the test case.
func (o RuleTestCaseOutput) Messages() RuleTestCaseMessageArrayOutput {
	return o.ApplyT(func(v *RuleTestCase) RuleTestCaseMessageArrayOutput { return v.Messages }).(RuleTestCaseMessageArrayOutput)
}

// The name of the test case.
func (o RuleTestCaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleTestCase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Sets if test case required to pass on update.
func (o RuleTestCaseOutput) Required() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RuleTestCase) pulumi.BoolPtrOutput { return v.Required }).(pulumi.BoolPtrOutput)
}

type RuleTestCaseArrayOutput struct{ *pulumi.OutputState }

func (RuleTestCaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuleTestCase)(nil)).Elem()
}

func (o RuleTestCaseArrayOutput) ToRuleTestCaseArrayOutput() RuleTestCaseArrayOutput {
	return o
}

func (o RuleTestCaseArrayOutput) ToRuleTestCaseArrayOutputWithContext(ctx context.Context) RuleTestCaseArrayOutput {
	return o
}

func (o RuleTestCaseArrayOutput) Index(i pulumi.IntInput) RuleTestCaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RuleTestCase {
		return vs[0].([]*RuleTestCase)[vs[1].(int)]
	}).(RuleTestCaseOutput)
}

type RuleTestCaseMapOutput struct{ *pulumi.OutputState }

func (RuleTestCaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuleTestCase)(nil)).Elem()
}

func (o RuleTestCaseMapOutput) ToRuleTestCaseMapOutput() RuleTestCaseMapOutput {
	return o
}

func (o RuleTestCaseMapOutput) ToRuleTestCaseMapOutputWithContext(ctx context.Context) RuleTestCaseMapOutput {
	return o
}

func (o RuleTestCaseMapOutput) MapIndex(k pulumi.StringInput) RuleTestCaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RuleTestCase {
		return vs[0].(map[string]*RuleTestCase)[vs[1].(string)]
	}).(RuleTestCaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleTestCaseInput)(nil)).Elem(), &RuleTestCase{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleTestCaseArrayInput)(nil)).Elem(), RuleTestCaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleTestCaseMapInput)(nil)).Elem(), RuleTestCaseMap{})
	pulumi.RegisterOutputType(RuleTestCaseOutput{})
	pulumi.RegisterOutputType(RuleTestCaseArrayOutput{})
	pulumi.RegisterOutputType(RuleTestCaseMapOutput{})
}
