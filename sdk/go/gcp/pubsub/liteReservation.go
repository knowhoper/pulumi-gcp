// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A named resource representing a shared pool of capacity.
//
// To get more information about Reservation, see:
//
// * [API documentation](https://cloud.google.com/pubsub/lite/docs/reference/rest/v1/admin.projects.locations.reservations)
// * How-to Guides
//     * [Managing Reservations](https://cloud.google.com/pubsub/lite/docs/reservations)
//
// ## Example Usage
// ### Pubsub Lite Reservation Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		project, err := organizations.LookupProject(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pubsub.NewLiteReservation(ctx, "example", &pubsub.LiteReservationArgs{
// 			Project:            pulumi.String(project.Number),
// 			ThroughputCapacity: pulumi.Int(2),
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
// Reservation can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:pubsub/liteReservation:LiteReservation default projects/{{project}}/locations/{{region}}/reservations/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:pubsub/liteReservation:LiteReservation default {{project}}/{{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:pubsub/liteReservation:LiteReservation default {{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:pubsub/liteReservation:LiteReservation default {{name}}
// ```
type LiteReservation struct {
	pulumi.CustomResourceState

	// Name of the reservation.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region of the pubsub lite reservation.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// The reserved throughput capacity. Every unit of throughput capacity is
	// equivalent to 1 MiB/s of published messages or 2 MiB/s of subscribed
	// messages.
	ThroughputCapacity pulumi.IntOutput `pulumi:"throughputCapacity"`
}

// NewLiteReservation registers a new resource with the given unique name, arguments, and options.
func NewLiteReservation(ctx *pulumi.Context,
	name string, args *LiteReservationArgs, opts ...pulumi.ResourceOption) (*LiteReservation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ThroughputCapacity == nil {
		return nil, errors.New("invalid value for required argument 'ThroughputCapacity'")
	}
	var resource LiteReservation
	err := ctx.RegisterResource("gcp:pubsub/liteReservation:LiteReservation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLiteReservation gets an existing LiteReservation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLiteReservation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LiteReservationState, opts ...pulumi.ResourceOption) (*LiteReservation, error) {
	var resource LiteReservation
	err := ctx.ReadResource("gcp:pubsub/liteReservation:LiteReservation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LiteReservation resources.
type liteReservationState struct {
	// Name of the reservation.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the pubsub lite reservation.
	Region *string `pulumi:"region"`
	// The reserved throughput capacity. Every unit of throughput capacity is
	// equivalent to 1 MiB/s of published messages or 2 MiB/s of subscribed
	// messages.
	ThroughputCapacity *int `pulumi:"throughputCapacity"`
}

type LiteReservationState struct {
	// Name of the reservation.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the pubsub lite reservation.
	Region pulumi.StringPtrInput
	// The reserved throughput capacity. Every unit of throughput capacity is
	// equivalent to 1 MiB/s of published messages or 2 MiB/s of subscribed
	// messages.
	ThroughputCapacity pulumi.IntPtrInput
}

func (LiteReservationState) ElementType() reflect.Type {
	return reflect.TypeOf((*liteReservationState)(nil)).Elem()
}

type liteReservationArgs struct {
	// Name of the reservation.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the pubsub lite reservation.
	Region *string `pulumi:"region"`
	// The reserved throughput capacity. Every unit of throughput capacity is
	// equivalent to 1 MiB/s of published messages or 2 MiB/s of subscribed
	// messages.
	ThroughputCapacity int `pulumi:"throughputCapacity"`
}

// The set of arguments for constructing a LiteReservation resource.
type LiteReservationArgs struct {
	// Name of the reservation.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the pubsub lite reservation.
	Region pulumi.StringPtrInput
	// The reserved throughput capacity. Every unit of throughput capacity is
	// equivalent to 1 MiB/s of published messages or 2 MiB/s of subscribed
	// messages.
	ThroughputCapacity pulumi.IntInput
}

func (LiteReservationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*liteReservationArgs)(nil)).Elem()
}

type LiteReservationInput interface {
	pulumi.Input

	ToLiteReservationOutput() LiteReservationOutput
	ToLiteReservationOutputWithContext(ctx context.Context) LiteReservationOutput
}

func (*LiteReservation) ElementType() reflect.Type {
	return reflect.TypeOf((*LiteReservation)(nil))
}

func (i *LiteReservation) ToLiteReservationOutput() LiteReservationOutput {
	return i.ToLiteReservationOutputWithContext(context.Background())
}

func (i *LiteReservation) ToLiteReservationOutputWithContext(ctx context.Context) LiteReservationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiteReservationOutput)
}

func (i *LiteReservation) ToLiteReservationPtrOutput() LiteReservationPtrOutput {
	return i.ToLiteReservationPtrOutputWithContext(context.Background())
}

func (i *LiteReservation) ToLiteReservationPtrOutputWithContext(ctx context.Context) LiteReservationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiteReservationPtrOutput)
}

type LiteReservationPtrInput interface {
	pulumi.Input

	ToLiteReservationPtrOutput() LiteReservationPtrOutput
	ToLiteReservationPtrOutputWithContext(ctx context.Context) LiteReservationPtrOutput
}

type liteReservationPtrType LiteReservationArgs

func (*liteReservationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiteReservation)(nil))
}

func (i *liteReservationPtrType) ToLiteReservationPtrOutput() LiteReservationPtrOutput {
	return i.ToLiteReservationPtrOutputWithContext(context.Background())
}

func (i *liteReservationPtrType) ToLiteReservationPtrOutputWithContext(ctx context.Context) LiteReservationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiteReservationPtrOutput)
}

// LiteReservationArrayInput is an input type that accepts LiteReservationArray and LiteReservationArrayOutput values.
// You can construct a concrete instance of `LiteReservationArrayInput` via:
//
//          LiteReservationArray{ LiteReservationArgs{...} }
type LiteReservationArrayInput interface {
	pulumi.Input

	ToLiteReservationArrayOutput() LiteReservationArrayOutput
	ToLiteReservationArrayOutputWithContext(context.Context) LiteReservationArrayOutput
}

type LiteReservationArray []LiteReservationInput

func (LiteReservationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LiteReservation)(nil)).Elem()
}

func (i LiteReservationArray) ToLiteReservationArrayOutput() LiteReservationArrayOutput {
	return i.ToLiteReservationArrayOutputWithContext(context.Background())
}

func (i LiteReservationArray) ToLiteReservationArrayOutputWithContext(ctx context.Context) LiteReservationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiteReservationArrayOutput)
}

// LiteReservationMapInput is an input type that accepts LiteReservationMap and LiteReservationMapOutput values.
// You can construct a concrete instance of `LiteReservationMapInput` via:
//
//          LiteReservationMap{ "key": LiteReservationArgs{...} }
type LiteReservationMapInput interface {
	pulumi.Input

	ToLiteReservationMapOutput() LiteReservationMapOutput
	ToLiteReservationMapOutputWithContext(context.Context) LiteReservationMapOutput
}

type LiteReservationMap map[string]LiteReservationInput

func (LiteReservationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LiteReservation)(nil)).Elem()
}

func (i LiteReservationMap) ToLiteReservationMapOutput() LiteReservationMapOutput {
	return i.ToLiteReservationMapOutputWithContext(context.Background())
}

func (i LiteReservationMap) ToLiteReservationMapOutputWithContext(ctx context.Context) LiteReservationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiteReservationMapOutput)
}

type LiteReservationOutput struct{ *pulumi.OutputState }

func (LiteReservationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiteReservation)(nil))
}

func (o LiteReservationOutput) ToLiteReservationOutput() LiteReservationOutput {
	return o
}

func (o LiteReservationOutput) ToLiteReservationOutputWithContext(ctx context.Context) LiteReservationOutput {
	return o
}

func (o LiteReservationOutput) ToLiteReservationPtrOutput() LiteReservationPtrOutput {
	return o.ToLiteReservationPtrOutputWithContext(context.Background())
}

func (o LiteReservationOutput) ToLiteReservationPtrOutputWithContext(ctx context.Context) LiteReservationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiteReservation) *LiteReservation {
		return &v
	}).(LiteReservationPtrOutput)
}

type LiteReservationPtrOutput struct{ *pulumi.OutputState }

func (LiteReservationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiteReservation)(nil))
}

func (o LiteReservationPtrOutput) ToLiteReservationPtrOutput() LiteReservationPtrOutput {
	return o
}

func (o LiteReservationPtrOutput) ToLiteReservationPtrOutputWithContext(ctx context.Context) LiteReservationPtrOutput {
	return o
}

func (o LiteReservationPtrOutput) Elem() LiteReservationOutput {
	return o.ApplyT(func(v *LiteReservation) LiteReservation {
		if v != nil {
			return *v
		}
		var ret LiteReservation
		return ret
	}).(LiteReservationOutput)
}

type LiteReservationArrayOutput struct{ *pulumi.OutputState }

func (LiteReservationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiteReservation)(nil))
}

func (o LiteReservationArrayOutput) ToLiteReservationArrayOutput() LiteReservationArrayOutput {
	return o
}

func (o LiteReservationArrayOutput) ToLiteReservationArrayOutputWithContext(ctx context.Context) LiteReservationArrayOutput {
	return o
}

func (o LiteReservationArrayOutput) Index(i pulumi.IntInput) LiteReservationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LiteReservation {
		return vs[0].([]LiteReservation)[vs[1].(int)]
	}).(LiteReservationOutput)
}

type LiteReservationMapOutput struct{ *pulumi.OutputState }

func (LiteReservationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LiteReservation)(nil))
}

func (o LiteReservationMapOutput) ToLiteReservationMapOutput() LiteReservationMapOutput {
	return o
}

func (o LiteReservationMapOutput) ToLiteReservationMapOutputWithContext(ctx context.Context) LiteReservationMapOutput {
	return o
}

func (o LiteReservationMapOutput) MapIndex(k pulumi.StringInput) LiteReservationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LiteReservation {
		return vs[0].(map[string]LiteReservation)[vs[1].(string)]
	}).(LiteReservationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LiteReservationInput)(nil)).Elem(), &LiteReservation{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiteReservationPtrInput)(nil)).Elem(), &LiteReservation{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiteReservationArrayInput)(nil)).Elem(), LiteReservationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiteReservationMapInput)(nil)).Elem(), LiteReservationMap{})
	pulumi.RegisterOutputType(LiteReservationOutput{})
	pulumi.RegisterOutputType(LiteReservationPtrOutput{})
	pulumi.RegisterOutputType(LiteReservationArrayOutput{})
	pulumi.RegisterOutputType(LiteReservationMapOutput{})
}