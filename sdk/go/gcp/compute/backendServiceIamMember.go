// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BackendServiceIamMember struct {
	pulumi.CustomResourceState

	Condition BackendServiceIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag      pulumi.StringOutput                       `pulumi:"etag"`
	Member    pulumi.StringOutput                       `pulumi:"member"`
	Name      pulumi.StringOutput                       `pulumi:"name"`
	Project   pulumi.StringOutput                       `pulumi:"project"`
	Role      pulumi.StringOutput                       `pulumi:"role"`
}

// NewBackendServiceIamMember registers a new resource with the given unique name, arguments, and options.
func NewBackendServiceIamMember(ctx *pulumi.Context,
	name string, args *BackendServiceIamMemberArgs, opts ...pulumi.ResourceOption) (*BackendServiceIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource BackendServiceIamMember
	err := ctx.RegisterResource("gcp:compute/backendServiceIamMember:BackendServiceIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendServiceIamMember gets an existing BackendServiceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendServiceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendServiceIamMemberState, opts ...pulumi.ResourceOption) (*BackendServiceIamMember, error) {
	var resource BackendServiceIamMember
	err := ctx.ReadResource("gcp:compute/backendServiceIamMember:BackendServiceIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendServiceIamMember resources.
type backendServiceIamMemberState struct {
	Condition *BackendServiceIamMemberCondition `pulumi:"condition"`
	Etag      *string                           `pulumi:"etag"`
	Member    *string                           `pulumi:"member"`
	Name      *string                           `pulumi:"name"`
	Project   *string                           `pulumi:"project"`
	Role      *string                           `pulumi:"role"`
}

type BackendServiceIamMemberState struct {
	Condition BackendServiceIamMemberConditionPtrInput
	Etag      pulumi.StringPtrInput
	Member    pulumi.StringPtrInput
	Name      pulumi.StringPtrInput
	Project   pulumi.StringPtrInput
	Role      pulumi.StringPtrInput
}

func (BackendServiceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendServiceIamMemberState)(nil)).Elem()
}

type backendServiceIamMemberArgs struct {
	Condition *BackendServiceIamMemberCondition `pulumi:"condition"`
	Member    string                            `pulumi:"member"`
	Name      *string                           `pulumi:"name"`
	Project   *string                           `pulumi:"project"`
	Role      string                            `pulumi:"role"`
}

// The set of arguments for constructing a BackendServiceIamMember resource.
type BackendServiceIamMemberArgs struct {
	Condition BackendServiceIamMemberConditionPtrInput
	Member    pulumi.StringInput
	Name      pulumi.StringPtrInput
	Project   pulumi.StringPtrInput
	Role      pulumi.StringInput
}

func (BackendServiceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendServiceIamMemberArgs)(nil)).Elem()
}

type BackendServiceIamMemberInput interface {
	pulumi.Input

	ToBackendServiceIamMemberOutput() BackendServiceIamMemberOutput
	ToBackendServiceIamMemberOutputWithContext(ctx context.Context) BackendServiceIamMemberOutput
}

func (*BackendServiceIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceIamMember)(nil)).Elem()
}

func (i *BackendServiceIamMember) ToBackendServiceIamMemberOutput() BackendServiceIamMemberOutput {
	return i.ToBackendServiceIamMemberOutputWithContext(context.Background())
}

func (i *BackendServiceIamMember) ToBackendServiceIamMemberOutputWithContext(ctx context.Context) BackendServiceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceIamMemberOutput)
}

// BackendServiceIamMemberArrayInput is an input type that accepts BackendServiceIamMemberArray and BackendServiceIamMemberArrayOutput values.
// You can construct a concrete instance of `BackendServiceIamMemberArrayInput` via:
//
//          BackendServiceIamMemberArray{ BackendServiceIamMemberArgs{...} }
type BackendServiceIamMemberArrayInput interface {
	pulumi.Input

	ToBackendServiceIamMemberArrayOutput() BackendServiceIamMemberArrayOutput
	ToBackendServiceIamMemberArrayOutputWithContext(context.Context) BackendServiceIamMemberArrayOutput
}

type BackendServiceIamMemberArray []BackendServiceIamMemberInput

func (BackendServiceIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendServiceIamMember)(nil)).Elem()
}

func (i BackendServiceIamMemberArray) ToBackendServiceIamMemberArrayOutput() BackendServiceIamMemberArrayOutput {
	return i.ToBackendServiceIamMemberArrayOutputWithContext(context.Background())
}

func (i BackendServiceIamMemberArray) ToBackendServiceIamMemberArrayOutputWithContext(ctx context.Context) BackendServiceIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceIamMemberArrayOutput)
}

// BackendServiceIamMemberMapInput is an input type that accepts BackendServiceIamMemberMap and BackendServiceIamMemberMapOutput values.
// You can construct a concrete instance of `BackendServiceIamMemberMapInput` via:
//
//          BackendServiceIamMemberMap{ "key": BackendServiceIamMemberArgs{...} }
type BackendServiceIamMemberMapInput interface {
	pulumi.Input

	ToBackendServiceIamMemberMapOutput() BackendServiceIamMemberMapOutput
	ToBackendServiceIamMemberMapOutputWithContext(context.Context) BackendServiceIamMemberMapOutput
}

type BackendServiceIamMemberMap map[string]BackendServiceIamMemberInput

func (BackendServiceIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendServiceIamMember)(nil)).Elem()
}

func (i BackendServiceIamMemberMap) ToBackendServiceIamMemberMapOutput() BackendServiceIamMemberMapOutput {
	return i.ToBackendServiceIamMemberMapOutputWithContext(context.Background())
}

func (i BackendServiceIamMemberMap) ToBackendServiceIamMemberMapOutputWithContext(ctx context.Context) BackendServiceIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceIamMemberMapOutput)
}

type BackendServiceIamMemberOutput struct{ *pulumi.OutputState }

func (BackendServiceIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceIamMember)(nil)).Elem()
}

func (o BackendServiceIamMemberOutput) ToBackendServiceIamMemberOutput() BackendServiceIamMemberOutput {
	return o
}

func (o BackendServiceIamMemberOutput) ToBackendServiceIamMemberOutputWithContext(ctx context.Context) BackendServiceIamMemberOutput {
	return o
}

type BackendServiceIamMemberArrayOutput struct{ *pulumi.OutputState }

func (BackendServiceIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendServiceIamMember)(nil)).Elem()
}

func (o BackendServiceIamMemberArrayOutput) ToBackendServiceIamMemberArrayOutput() BackendServiceIamMemberArrayOutput {
	return o
}

func (o BackendServiceIamMemberArrayOutput) ToBackendServiceIamMemberArrayOutputWithContext(ctx context.Context) BackendServiceIamMemberArrayOutput {
	return o
}

func (o BackendServiceIamMemberArrayOutput) Index(i pulumi.IntInput) BackendServiceIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackendServiceIamMember {
		return vs[0].([]*BackendServiceIamMember)[vs[1].(int)]
	}).(BackendServiceIamMemberOutput)
}

type BackendServiceIamMemberMapOutput struct{ *pulumi.OutputState }

func (BackendServiceIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendServiceIamMember)(nil)).Elem()
}

func (o BackendServiceIamMemberMapOutput) ToBackendServiceIamMemberMapOutput() BackendServiceIamMemberMapOutput {
	return o
}

func (o BackendServiceIamMemberMapOutput) ToBackendServiceIamMemberMapOutputWithContext(ctx context.Context) BackendServiceIamMemberMapOutput {
	return o
}

func (o BackendServiceIamMemberMapOutput) MapIndex(k pulumi.StringInput) BackendServiceIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackendServiceIamMember {
		return vs[0].(map[string]*BackendServiceIamMember)[vs[1].(string)]
	}).(BackendServiceIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServiceIamMemberInput)(nil)).Elem(), &BackendServiceIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServiceIamMemberArrayInput)(nil)).Elem(), BackendServiceIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServiceIamMemberMapInput)(nil)).Elem(), BackendServiceIamMemberMap{})
	pulumi.RegisterOutputType(BackendServiceIamMemberOutput{})
	pulumi.RegisterOutputType(BackendServiceIamMemberArrayOutput{})
	pulumi.RegisterOutputType(BackendServiceIamMemberMapOutput{})
}