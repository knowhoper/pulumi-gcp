// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package secretmanager

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Three different resources help you manage your IAM policy for Secret Manager Secret. Each of these resources serves a different use case:
//
// * `secretmanager.SecretIamPolicy`: Authoritative. Sets the IAM policy for the secret and replaces any existing policy already attached.
// * `secretmanager.SecretIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the secret are preserved.
// * `secretmanager.SecretIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the secret are preserved.
//
// > **Note:** `secretmanager.SecretIamPolicy` **cannot** be used in conjunction with `secretmanager.SecretIamBinding` and `secretmanager.SecretIamMember` or they will fight over what your policy should be.
//
// > **Note:** `secretmanager.SecretIamBinding` resources **can be** used in conjunction with `secretmanager.SecretIamMember` resources **only if** they do not grant privilege to the same role.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/secret_manager_secret_iam.html.markdown.
type SecretIamBinding struct {
	pulumi.CustomResourceState

	Condition SecretIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `secretmanager.SecretIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role     pulumi.StringOutput `pulumi:"role"`
	SecretId pulumi.StringOutput `pulumi:"secretId"`
}

// NewSecretIamBinding registers a new resource with the given unique name, arguments, and options.
func NewSecretIamBinding(ctx *pulumi.Context,
	name string, args *SecretIamBindingArgs, opts ...pulumi.ResourceOption) (*SecretIamBinding, error) {
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.SecretId == nil {
		return nil, errors.New("missing required argument 'SecretId'")
	}
	if args == nil {
		args = &SecretIamBindingArgs{}
	}
	var resource SecretIamBinding
	err := ctx.RegisterResource("gcp:secretmanager/secretIamBinding:SecretIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretIamBinding gets an existing SecretIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretIamBindingState, opts ...pulumi.ResourceOption) (*SecretIamBinding, error) {
	var resource SecretIamBinding
	err := ctx.ReadResource("gcp:secretmanager/secretIamBinding:SecretIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretIamBinding resources.
type secretIamBindingState struct {
	Condition *SecretIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `secretmanager.SecretIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role     *string `pulumi:"role"`
	SecretId *string `pulumi:"secretId"`
}

type SecretIamBindingState struct {
	Condition SecretIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `secretmanager.SecretIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role     pulumi.StringPtrInput
	SecretId pulumi.StringPtrInput
}

func (SecretIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretIamBindingState)(nil)).Elem()
}

type secretIamBindingArgs struct {
	Condition *SecretIamBindingCondition `pulumi:"condition"`
	Members   []string                   `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `secretmanager.SecretIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role     string `pulumi:"role"`
	SecretId string `pulumi:"secretId"`
}

// The set of arguments for constructing a SecretIamBinding resource.
type SecretIamBindingArgs struct {
	Condition SecretIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `secretmanager.SecretIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role     pulumi.StringInput
	SecretId pulumi.StringInput
}

func (SecretIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretIamBindingArgs)(nil)).Elem()
}