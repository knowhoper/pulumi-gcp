// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package assuredworkloads

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AssuredWorkloads Workload resource
//
// ## Example Usage
// ### Basic_workload
// A basic test of a assuredworkloads api
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/assuredworkloads"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := assuredworkloads.NewWorkload(ctx, "primary", &assuredworkloads.WorkloadArgs{
// 			BillingAccount:   pulumi.String("billingAccounts/000000-0000000-0000000-000000"),
// 			ComplianceRegime: pulumi.String("FEDRAMP_MODERATE"),
// 			DisplayName:      pulumi.String("Workload Example"),
// 			KmsSettings: &assuredworkloads.WorkloadKmsSettingsArgs{
// 				NextRotationTime: pulumi.String("9999-10-02T15:01:23Z"),
// 				RotationPeriod:   pulumi.String("10368000s"),
// 			},
// 			Labels: pulumi.StringMap{
// 				"label-one": pulumi.String("value-one"),
// 			},
// 			Location:                   pulumi.String("us-west1"),
// 			Organization:               pulumi.String("123456789"),
// 			ProvisionedResourcesParent: pulumi.String("folders/519620126891"),
// 			ResourceSettings: assuredworkloads.WorkloadResourceSettingArray{
// 				&assuredworkloads.WorkloadResourceSettingArgs{
// 					ResourceType: pulumi.String("CONSUMER_PROJECT"),
// 				},
// 				&assuredworkloads.WorkloadResourceSettingArgs{
// 					ResourceType: pulumi.String("ENCRYPTION_KEYS_PROJECT"),
// 				},
// 				&assuredworkloads.WorkloadResourceSettingArgs{
// 					ResourceId:   pulumi.String("ring"),
// 					ResourceType: pulumi.String("KEYRING"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Workload can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:assuredworkloads/workload:Workload default organizations/{{organization}}/locations/{{location}}/workloads/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:assuredworkloads/workload:Workload default {{organization}}/{{location}}/{{name}}
// ```
type Workload struct {
	pulumi.CustomResourceState

	// Required. Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, 'billingAccounts/012345-567890-ABCDEF`.
	BillingAccount pulumi.StringOutput `pulumi:"billingAccount"`
	// Required. Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS
	ComplianceRegime pulumi.StringOutput `pulumi:"complianceRegime"`
	// Output only. Immutable. The Workload creation timestamp.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Required. The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Input only. Settings used to create a CMEK crypto key. When set a project with a KMS CMEK key is provisioned. This field is mandatory for a subset of Compliance Regimes.
	KmsSettings WorkloadKmsSettingsPtrOutput `pulumi:"kmsSettings"`
	// Optional. Labels applied to the workload.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// Output only. The resource name of the workload.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization for the resource
	Organization pulumi.StringOutput `pulumi:"organization"`
	// Input only. The parent resource for the resources managed by this Assured Workload. May be either an organization or a folder. Must be the same or a child of the Workload parent. If not specified all resources are created under the Workload parent. Formats: folders/{folder_id}, organizations/{organization_id}
	ProvisionedResourcesParent pulumi.StringPtrOutput `pulumi:"provisionedResourcesParent"`
	// Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
	ResourceSettings WorkloadResourceSettingArrayOutput `pulumi:"resourceSettings"`
	// Output only. The resources associated with this workload. These resources will be created when creating the workload. If
	// any of the projects already exist, the workload creation will fail. Always read only.
	Resources WorkloadResourceArrayOutput `pulumi:"resources"`
}

// NewWorkload registers a new resource with the given unique name, arguments, and options.
func NewWorkload(ctx *pulumi.Context,
	name string, args *WorkloadArgs, opts ...pulumi.ResourceOption) (*Workload, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccount == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccount'")
	}
	if args.ComplianceRegime == nil {
		return nil, errors.New("invalid value for required argument 'ComplianceRegime'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	var resource Workload
	err := ctx.RegisterResource("gcp:assuredworkloads/workload:Workload", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkload gets an existing Workload resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkload(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadState, opts ...pulumi.ResourceOption) (*Workload, error) {
	var resource Workload
	err := ctx.ReadResource("gcp:assuredworkloads/workload:Workload", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workload resources.
type workloadState struct {
	// Required. Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, 'billingAccounts/012345-567890-ABCDEF`.
	BillingAccount *string `pulumi:"billingAccount"`
	// Required. Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS
	ComplianceRegime *string `pulumi:"complianceRegime"`
	// Output only. Immutable. The Workload creation timestamp.
	CreateTime *string `pulumi:"createTime"`
	// Required. The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
	DisplayName *string `pulumi:"displayName"`
	// Input only. Settings used to create a CMEK crypto key. When set a project with a KMS CMEK key is provisioned. This field is mandatory for a subset of Compliance Regimes.
	KmsSettings *WorkloadKmsSettings `pulumi:"kmsSettings"`
	// Optional. Labels applied to the workload.
	Labels map[string]string `pulumi:"labels"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// Output only. The resource name of the workload.
	Name *string `pulumi:"name"`
	// The organization for the resource
	Organization *string `pulumi:"organization"`
	// Input only. The parent resource for the resources managed by this Assured Workload. May be either an organization or a folder. Must be the same or a child of the Workload parent. If not specified all resources are created under the Workload parent. Formats: folders/{folder_id}, organizations/{organization_id}
	ProvisionedResourcesParent *string `pulumi:"provisionedResourcesParent"`
	// Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
	ResourceSettings []WorkloadResourceSetting `pulumi:"resourceSettings"`
	// Output only. The resources associated with this workload. These resources will be created when creating the workload. If
	// any of the projects already exist, the workload creation will fail. Always read only.
	Resources []WorkloadResource `pulumi:"resources"`
}

type WorkloadState struct {
	// Required. Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, 'billingAccounts/012345-567890-ABCDEF`.
	BillingAccount pulumi.StringPtrInput
	// Required. Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS
	ComplianceRegime pulumi.StringPtrInput
	// Output only. Immutable. The Workload creation timestamp.
	CreateTime pulumi.StringPtrInput
	// Required. The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
	DisplayName pulumi.StringPtrInput
	// Input only. Settings used to create a CMEK crypto key. When set a project with a KMS CMEK key is provisioned. This field is mandatory for a subset of Compliance Regimes.
	KmsSettings WorkloadKmsSettingsPtrInput
	// Optional. Labels applied to the workload.
	Labels pulumi.StringMapInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// Output only. The resource name of the workload.
	Name pulumi.StringPtrInput
	// The organization for the resource
	Organization pulumi.StringPtrInput
	// Input only. The parent resource for the resources managed by this Assured Workload. May be either an organization or a folder. Must be the same or a child of the Workload parent. If not specified all resources are created under the Workload parent. Formats: folders/{folder_id}, organizations/{organization_id}
	ProvisionedResourcesParent pulumi.StringPtrInput
	// Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
	ResourceSettings WorkloadResourceSettingArrayInput
	// Output only. The resources associated with this workload. These resources will be created when creating the workload. If
	// any of the projects already exist, the workload creation will fail. Always read only.
	Resources WorkloadResourceArrayInput
}

func (WorkloadState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadState)(nil)).Elem()
}

type workloadArgs struct {
	// Required. Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, 'billingAccounts/012345-567890-ABCDEF`.
	BillingAccount string `pulumi:"billingAccount"`
	// Required. Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS
	ComplianceRegime string `pulumi:"complianceRegime"`
	// Required. The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
	DisplayName string `pulumi:"displayName"`
	// Input only. Settings used to create a CMEK crypto key. When set a project with a KMS CMEK key is provisioned. This field is mandatory for a subset of Compliance Regimes.
	KmsSettings *WorkloadKmsSettings `pulumi:"kmsSettings"`
	// Optional. Labels applied to the workload.
	Labels map[string]string `pulumi:"labels"`
	// The location for the resource
	Location string `pulumi:"location"`
	// The organization for the resource
	Organization string `pulumi:"organization"`
	// Input only. The parent resource for the resources managed by this Assured Workload. May be either an organization or a folder. Must be the same or a child of the Workload parent. If not specified all resources are created under the Workload parent. Formats: folders/{folder_id}, organizations/{organization_id}
	ProvisionedResourcesParent *string `pulumi:"provisionedResourcesParent"`
	// Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
	ResourceSettings []WorkloadResourceSetting `pulumi:"resourceSettings"`
}

// The set of arguments for constructing a Workload resource.
type WorkloadArgs struct {
	// Required. Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, 'billingAccounts/012345-567890-ABCDEF`.
	BillingAccount pulumi.StringInput
	// Required. Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS
	ComplianceRegime pulumi.StringInput
	// Required. The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
	DisplayName pulumi.StringInput
	// Input only. Settings used to create a CMEK crypto key. When set a project with a KMS CMEK key is provisioned. This field is mandatory for a subset of Compliance Regimes.
	KmsSettings WorkloadKmsSettingsPtrInput
	// Optional. Labels applied to the workload.
	Labels pulumi.StringMapInput
	// The location for the resource
	Location pulumi.StringInput
	// The organization for the resource
	Organization pulumi.StringInput
	// Input only. The parent resource for the resources managed by this Assured Workload. May be either an organization or a folder. Must be the same or a child of the Workload parent. If not specified all resources are created under the Workload parent. Formats: folders/{folder_id}, organizations/{organization_id}
	ProvisionedResourcesParent pulumi.StringPtrInput
	// Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
	ResourceSettings WorkloadResourceSettingArrayInput
}

func (WorkloadArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadArgs)(nil)).Elem()
}

type WorkloadInput interface {
	pulumi.Input

	ToWorkloadOutput() WorkloadOutput
	ToWorkloadOutputWithContext(ctx context.Context) WorkloadOutput
}

func (*Workload) ElementType() reflect.Type {
	return reflect.TypeOf((**Workload)(nil)).Elem()
}

func (i *Workload) ToWorkloadOutput() WorkloadOutput {
	return i.ToWorkloadOutputWithContext(context.Background())
}

func (i *Workload) ToWorkloadOutputWithContext(ctx context.Context) WorkloadOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadOutput)
}

// WorkloadArrayInput is an input type that accepts WorkloadArray and WorkloadArrayOutput values.
// You can construct a concrete instance of `WorkloadArrayInput` via:
//
//          WorkloadArray{ WorkloadArgs{...} }
type WorkloadArrayInput interface {
	pulumi.Input

	ToWorkloadArrayOutput() WorkloadArrayOutput
	ToWorkloadArrayOutputWithContext(context.Context) WorkloadArrayOutput
}

type WorkloadArray []WorkloadInput

func (WorkloadArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Workload)(nil)).Elem()
}

func (i WorkloadArray) ToWorkloadArrayOutput() WorkloadArrayOutput {
	return i.ToWorkloadArrayOutputWithContext(context.Background())
}

func (i WorkloadArray) ToWorkloadArrayOutputWithContext(ctx context.Context) WorkloadArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadArrayOutput)
}

// WorkloadMapInput is an input type that accepts WorkloadMap and WorkloadMapOutput values.
// You can construct a concrete instance of `WorkloadMapInput` via:
//
//          WorkloadMap{ "key": WorkloadArgs{...} }
type WorkloadMapInput interface {
	pulumi.Input

	ToWorkloadMapOutput() WorkloadMapOutput
	ToWorkloadMapOutputWithContext(context.Context) WorkloadMapOutput
}

type WorkloadMap map[string]WorkloadInput

func (WorkloadMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Workload)(nil)).Elem()
}

func (i WorkloadMap) ToWorkloadMapOutput() WorkloadMapOutput {
	return i.ToWorkloadMapOutputWithContext(context.Background())
}

func (i WorkloadMap) ToWorkloadMapOutputWithContext(ctx context.Context) WorkloadMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadMapOutput)
}

type WorkloadOutput struct{ *pulumi.OutputState }

func (WorkloadOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workload)(nil)).Elem()
}

func (o WorkloadOutput) ToWorkloadOutput() WorkloadOutput {
	return o
}

func (o WorkloadOutput) ToWorkloadOutputWithContext(ctx context.Context) WorkloadOutput {
	return o
}

type WorkloadArrayOutput struct{ *pulumi.OutputState }

func (WorkloadArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Workload)(nil)).Elem()
}

func (o WorkloadArrayOutput) ToWorkloadArrayOutput() WorkloadArrayOutput {
	return o
}

func (o WorkloadArrayOutput) ToWorkloadArrayOutputWithContext(ctx context.Context) WorkloadArrayOutput {
	return o
}

func (o WorkloadArrayOutput) Index(i pulumi.IntInput) WorkloadOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Workload {
		return vs[0].([]*Workload)[vs[1].(int)]
	}).(WorkloadOutput)
}

type WorkloadMapOutput struct{ *pulumi.OutputState }

func (WorkloadMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Workload)(nil)).Elem()
}

func (o WorkloadMapOutput) ToWorkloadMapOutput() WorkloadMapOutput {
	return o
}

func (o WorkloadMapOutput) ToWorkloadMapOutputWithContext(ctx context.Context) WorkloadMapOutput {
	return o
}

func (o WorkloadMapOutput) MapIndex(k pulumi.StringInput) WorkloadOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Workload {
		return vs[0].(map[string]*Workload)[vs[1].(string)]
	}).(WorkloadOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadInput)(nil)).Elem(), &Workload{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadArrayInput)(nil)).Elem(), WorkloadArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadMapInput)(nil)).Elem(), WorkloadMap{})
	pulumi.RegisterOutputType(WorkloadOutput{})
	pulumi.RegisterOutputType(WorkloadArrayOutput{})
	pulumi.RegisterOutputType(WorkloadMapOutput{})
}