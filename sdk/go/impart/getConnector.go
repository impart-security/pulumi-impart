// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package impart

import (
	"context"
	"reflect"

	"github.com/impart-security/pulumi-impart/sdk/go/impart/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage a connector.
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
//			_, err := impart.GetConnector(ctx, &impart.GetConnectorArgs{
//				Id: "<id>",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetConnector(ctx *pulumi.Context, args *GetConnectorArgs, opts ...pulumi.InvokeOption) (*GetConnectorResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetConnectorResult
	err := ctx.Invoke("impart:index/getConnector:GetConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetConnector.
type GetConnectorArgs struct {
	// ID of the connector type (eg. ID for our Slack or Jira connector types).
	ConnectorTypeId *string `pulumi:"connectorTypeId"`
	// Identifier for this connector.
	Id string `pulumi:"id"`
	// Whether or not the connector is authenticated via OAuth2.
	IsConnected *bool `pulumi:"isConnected"`
	// Name for this connector.
	Name *string `pulumi:"name"`
}

// A collection of values returned by GetConnector.
type GetConnectorResult struct {
	// ID of the connector type (eg. ID for our Slack or Jira connector types).
	ConnectorTypeId *string `pulumi:"connectorTypeId"`
	// Identifier for this connector.
	Id string `pulumi:"id"`
	// Whether or not the connector is authenticated via OAuth2.
	IsConnected *bool `pulumi:"isConnected"`
	// Name for this connector.
	Name *string `pulumi:"name"`
}

func GetConnectorOutput(ctx *pulumi.Context, args GetConnectorOutputArgs, opts ...pulumi.InvokeOption) GetConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetConnectorResultOutput, error) {
			args := v.(GetConnectorArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetConnectorResult
			secret, err := ctx.InvokePackageRaw("impart:index/getConnector:GetConnector", args, &rv, "", opts...)
			if err != nil {
				return GetConnectorResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetConnectorResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetConnectorResultOutput), nil
			}
			return output, nil
		}).(GetConnectorResultOutput)
}

// A collection of arguments for invoking GetConnector.
type GetConnectorOutputArgs struct {
	// ID of the connector type (eg. ID for our Slack or Jira connector types).
	ConnectorTypeId pulumi.StringPtrInput `pulumi:"connectorTypeId"`
	// Identifier for this connector.
	Id pulumi.StringInput `pulumi:"id"`
	// Whether or not the connector is authenticated via OAuth2.
	IsConnected pulumi.BoolPtrInput `pulumi:"isConnected"`
	// Name for this connector.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetConnectorArgs)(nil)).Elem()
}

// A collection of values returned by GetConnector.
type GetConnectorResultOutput struct{ *pulumi.OutputState }

func (GetConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetConnectorResult)(nil)).Elem()
}

func (o GetConnectorResultOutput) ToGetConnectorResultOutput() GetConnectorResultOutput {
	return o
}

func (o GetConnectorResultOutput) ToGetConnectorResultOutputWithContext(ctx context.Context) GetConnectorResultOutput {
	return o
}

// ID of the connector type (eg. ID for our Slack or Jira connector types).
func (o GetConnectorResultOutput) ConnectorTypeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetConnectorResult) *string { return v.ConnectorTypeId }).(pulumi.StringPtrOutput)
}

// Identifier for this connector.
func (o GetConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether or not the connector is authenticated via OAuth2.
func (o GetConnectorResultOutput) IsConnected() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetConnectorResult) *bool { return v.IsConnected }).(pulumi.BoolPtrOutput)
}

// Name for this connector.
func (o GetConnectorResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetConnectorResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetConnectorResultOutput{})
}
