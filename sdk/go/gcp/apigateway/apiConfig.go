// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// ApiConfig can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:apigateway/apiConfig:ApiConfig default projects/{{project}}/locations/global/apis/{{api}}/configs/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:apigateway/apiConfig:ApiConfig default {{project}}/{{api}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:apigateway/apiConfig:ApiConfig default {{api}}/{{name}}
// ```
type ApiConfig struct {
	pulumi.CustomResourceState

	// The API to attach the config to.
	Api pulumi.StringOutput `pulumi:"api"`
	// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
	ApiConfigId pulumi.StringOutput `pulumi:"apiConfigId"`
	// Creates a unique name beginning with the
	// specified prefix. If this and apiConfigId are unspecified, a random value is chosen for the name.
	ApiConfigIdPrefix pulumi.StringOutput `pulumi:"apiConfigIdPrefix"`
	// A user-visible name for the API.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Immutable. Gateway specific configuration.
	// If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
	// Structure is documented below.
	GatewayConfig ApiConfigGatewayConfigPtrOutput `pulumi:"gatewayConfig"`
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource name of the API Config.
	Name pulumi.StringOutput `pulumi:"name"`
	// An OpenAPI Specification Document describing an API.
	// Structure is documented below.
	OpenapiDocuments ApiConfigOpenapiDocumentArrayOutput `pulumi:"openapiDocuments"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
	ServiceConfigId pulumi.StringOutput `pulumi:"serviceConfigId"`
}

// NewApiConfig registers a new resource with the given unique name, arguments, and options.
func NewApiConfig(ctx *pulumi.Context,
	name string, args *ApiConfigArgs, opts ...pulumi.ResourceOption) (*ApiConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Api == nil {
		return nil, errors.New("invalid value for required argument 'Api'")
	}
	if args.OpenapiDocuments == nil {
		return nil, errors.New("invalid value for required argument 'OpenapiDocuments'")
	}
	var resource ApiConfig
	err := ctx.RegisterResource("gcp:apigateway/apiConfig:ApiConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiConfig gets an existing ApiConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiConfigState, opts ...pulumi.ResourceOption) (*ApiConfig, error) {
	var resource ApiConfig
	err := ctx.ReadResource("gcp:apigateway/apiConfig:ApiConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiConfig resources.
type apiConfigState struct {
	// The API to attach the config to.
	Api *string `pulumi:"api"`
	// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
	ApiConfigId *string `pulumi:"apiConfigId"`
	// Creates a unique name beginning with the
	// specified prefix. If this and apiConfigId are unspecified, a random value is chosen for the name.
	ApiConfigIdPrefix *string `pulumi:"apiConfigIdPrefix"`
	// A user-visible name for the API.
	DisplayName *string `pulumi:"displayName"`
	// Immutable. Gateway specific configuration.
	// If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
	// Structure is documented below.
	GatewayConfig *ApiConfigGatewayConfig `pulumi:"gatewayConfig"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// The resource name of the API Config.
	Name *string `pulumi:"name"`
	// An OpenAPI Specification Document describing an API.
	// Structure is documented below.
	OpenapiDocuments []ApiConfigOpenapiDocument `pulumi:"openapiDocuments"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
	ServiceConfigId *string `pulumi:"serviceConfigId"`
}

type ApiConfigState struct {
	// The API to attach the config to.
	Api pulumi.StringPtrInput
	// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
	ApiConfigId pulumi.StringPtrInput
	// Creates a unique name beginning with the
	// specified prefix. If this and apiConfigId are unspecified, a random value is chosen for the name.
	ApiConfigIdPrefix pulumi.StringPtrInput
	// A user-visible name for the API.
	DisplayName pulumi.StringPtrInput
	// Immutable. Gateway specific configuration.
	// If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
	// Structure is documented below.
	GatewayConfig ApiConfigGatewayConfigPtrInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// The resource name of the API Config.
	Name pulumi.StringPtrInput
	// An OpenAPI Specification Document describing an API.
	// Structure is documented below.
	OpenapiDocuments ApiConfigOpenapiDocumentArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
	ServiceConfigId pulumi.StringPtrInput
}

func (ApiConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiConfigState)(nil)).Elem()
}

type apiConfigArgs struct {
	// The API to attach the config to.
	Api string `pulumi:"api"`
	// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
	ApiConfigId *string `pulumi:"apiConfigId"`
	// Creates a unique name beginning with the
	// specified prefix. If this and apiConfigId are unspecified, a random value is chosen for the name.
	ApiConfigIdPrefix *string `pulumi:"apiConfigIdPrefix"`
	// A user-visible name for the API.
	DisplayName *string `pulumi:"displayName"`
	// Immutable. Gateway specific configuration.
	// If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
	// Structure is documented below.
	GatewayConfig *ApiConfigGatewayConfig `pulumi:"gatewayConfig"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// An OpenAPI Specification Document describing an API.
	// Structure is documented below.
	OpenapiDocuments []ApiConfigOpenapiDocument `pulumi:"openapiDocuments"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a ApiConfig resource.
type ApiConfigArgs struct {
	// The API to attach the config to.
	Api pulumi.StringInput
	// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
	ApiConfigId pulumi.StringPtrInput
	// Creates a unique name beginning with the
	// specified prefix. If this and apiConfigId are unspecified, a random value is chosen for the name.
	ApiConfigIdPrefix pulumi.StringPtrInput
	// A user-visible name for the API.
	DisplayName pulumi.StringPtrInput
	// Immutable. Gateway specific configuration.
	// If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
	// Structure is documented below.
	GatewayConfig ApiConfigGatewayConfigPtrInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// An OpenAPI Specification Document describing an API.
	// Structure is documented below.
	OpenapiDocuments ApiConfigOpenapiDocumentArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ApiConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiConfigArgs)(nil)).Elem()
}

type ApiConfigInput interface {
	pulumi.Input

	ToApiConfigOutput() ApiConfigOutput
	ToApiConfigOutputWithContext(ctx context.Context) ApiConfigOutput
}

func (ApiConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiConfig)(nil)).Elem()
}

func (i ApiConfig) ToApiConfigOutput() ApiConfigOutput {
	return i.ToApiConfigOutputWithContext(context.Background())
}

func (i ApiConfig) ToApiConfigOutputWithContext(ctx context.Context) ApiConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConfigOutput)
}

type ApiConfigOutput struct {
	*pulumi.OutputState
}

func (ApiConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiConfigOutput)(nil)).Elem()
}

func (o ApiConfigOutput) ToApiConfigOutput() ApiConfigOutput {
	return o
}

func (o ApiConfigOutput) ToApiConfigOutputWithContext(ctx context.Context) ApiConfigOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApiConfigOutput{})
}