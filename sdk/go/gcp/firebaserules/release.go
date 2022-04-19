// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package firebaserules

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Firebaserules Release resource
//
// ## Example Usage
// ### Basic_release
// Creates a basic Firebase Rules Release
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/firebaserules"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		basic, err := firebaserules.NewRuleset(ctx, "basic", &firebaserules.RulesetArgs{
// 			Project: pulumi.String("my-project-name"),
// 			Source: &firebaserules.RulesetSourceArgs{
// 				Files: firebaserules.RulesetSourceFileArray{
// 					&firebaserules.RulesetSourceFileArgs{
// 						Content:     pulumi.String("service cloud.firestore {match /databases/{database}/documents { match /{document=**} { allow read, write: if false; } } }"),
// 						Fingerprint: pulumi.String(""),
// 						Name:        pulumi.String("firestore.rules"),
// 					},
// 				},
// 				Language: pulumi.String(""),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = firebaserules.NewRelease(ctx, "primary", &firebaserules.ReleaseArgs{
// 			Project: pulumi.String("my-project-name"),
// 			RulesetName: basic.Name.ApplyT(func(name string) (string, error) {
// 				return fmt.Sprintf("%v%v", "projects/my-project-name/rulesets/", name), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = firebaserules.NewRuleset(ctx, "minimal", &firebaserules.RulesetArgs{
// 			Project: pulumi.String("my-project-name"),
// 			Source: &firebaserules.RulesetSourceArgs{
// 				Files: firebaserules.RulesetSourceFileArray{
// 					&firebaserules.RulesetSourceFileArgs{
// 						Content: pulumi.String("service cloud.firestore {match /databases/{database}/documents { match /{document=**} { allow read, write: if false; } } }"),
// 						Name:    pulumi.String("firestore.rules"),
// 					},
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
// ### Minimal_release
// Creates a minimal Firebase Rules Release
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/firebaserules"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		minimal, err := firebaserules.NewRuleset(ctx, "minimal", &firebaserules.RulesetArgs{
// 			Project: pulumi.String("my-project-name"),
// 			Source: &firebaserules.RulesetSourceArgs{
// 				Files: firebaserules.RulesetSourceFileArray{
// 					&firebaserules.RulesetSourceFileArgs{
// 						Content: pulumi.String("service cloud.firestore {match /databases/{database}/documents { match /{document=**} { allow read, write: if false; } } }"),
// 						Name:    pulumi.String("firestore.rules"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = firebaserules.NewRelease(ctx, "primary", &firebaserules.ReleaseArgs{
// 			Project: pulumi.String("my-project-name"),
// 			RulesetName: minimal.Name.ApplyT(func(name string) (string, error) {
// 				return fmt.Sprintf("%v%v", "projects/my-project-name/rulesets/", name), nil
// 			}).(pulumi.StringOutput),
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
// Release can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:firebaserules/release:Release default projects/{{project}}/releases/{{name}}
// ```
type Release struct {
	pulumi.CustomResourceState

	// Output only. Time the release was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Disable the release to keep it from being served. The response code of NOT_FOUND will be given for executables generated
	// from this Release.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// Format: `projects/{project_id}/releases/{release_id}`
	Name pulumi.StringOutput `pulumi:"name"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist the `Release` to be created.
	RulesetName pulumi.StringOutput `pulumi:"rulesetName"`
	// Output only. Time the release was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewRelease registers a new resource with the given unique name, arguments, and options.
func NewRelease(ctx *pulumi.Context,
	name string, args *ReleaseArgs, opts ...pulumi.ResourceOption) (*Release, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RulesetName == nil {
		return nil, errors.New("invalid value for required argument 'RulesetName'")
	}
	var resource Release
	err := ctx.RegisterResource("gcp:firebaserules/release:Release", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRelease gets an existing Release resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRelease(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReleaseState, opts ...pulumi.ResourceOption) (*Release, error) {
	var resource Release
	err := ctx.ReadResource("gcp:firebaserules/release:Release", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Release resources.
type releaseState struct {
	// Output only. Time the release was created.
	CreateTime *string `pulumi:"createTime"`
	// Disable the release to keep it from being served. The response code of NOT_FOUND will be given for executables generated
	// from this Release.
	Disabled *bool `pulumi:"disabled"`
	// Format: `projects/{project_id}/releases/{release_id}`
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist the `Release` to be created.
	RulesetName *string `pulumi:"rulesetName"`
	// Output only. Time the release was updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type ReleaseState struct {
	// Output only. Time the release was created.
	CreateTime pulumi.StringPtrInput
	// Disable the release to keep it from being served. The response code of NOT_FOUND will be given for executables generated
	// from this Release.
	Disabled pulumi.BoolPtrInput
	// Format: `projects/{project_id}/releases/{release_id}`
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist the `Release` to be created.
	RulesetName pulumi.StringPtrInput
	// Output only. Time the release was updated.
	UpdateTime pulumi.StringPtrInput
}

func (ReleaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseState)(nil)).Elem()
}

type releaseArgs struct {
	// Format: `projects/{project_id}/releases/{release_id}`
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist the `Release` to be created.
	RulesetName string `pulumi:"rulesetName"`
}

// The set of arguments for constructing a Release resource.
type ReleaseArgs struct {
	// Format: `projects/{project_id}/releases/{release_id}`
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist the `Release` to be created.
	RulesetName pulumi.StringInput
}

func (ReleaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseArgs)(nil)).Elem()
}

type ReleaseInput interface {
	pulumi.Input

	ToReleaseOutput() ReleaseOutput
	ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput
}

func (*Release) ElementType() reflect.Type {
	return reflect.TypeOf((**Release)(nil)).Elem()
}

func (i *Release) ToReleaseOutput() ReleaseOutput {
	return i.ToReleaseOutputWithContext(context.Background())
}

func (i *Release) ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseOutput)
}

// ReleaseArrayInput is an input type that accepts ReleaseArray and ReleaseArrayOutput values.
// You can construct a concrete instance of `ReleaseArrayInput` via:
//
//          ReleaseArray{ ReleaseArgs{...} }
type ReleaseArrayInput interface {
	pulumi.Input

	ToReleaseArrayOutput() ReleaseArrayOutput
	ToReleaseArrayOutputWithContext(context.Context) ReleaseArrayOutput
}

type ReleaseArray []ReleaseInput

func (ReleaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Release)(nil)).Elem()
}

func (i ReleaseArray) ToReleaseArrayOutput() ReleaseArrayOutput {
	return i.ToReleaseArrayOutputWithContext(context.Background())
}

func (i ReleaseArray) ToReleaseArrayOutputWithContext(ctx context.Context) ReleaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseArrayOutput)
}

// ReleaseMapInput is an input type that accepts ReleaseMap and ReleaseMapOutput values.
// You can construct a concrete instance of `ReleaseMapInput` via:
//
//          ReleaseMap{ "key": ReleaseArgs{...} }
type ReleaseMapInput interface {
	pulumi.Input

	ToReleaseMapOutput() ReleaseMapOutput
	ToReleaseMapOutputWithContext(context.Context) ReleaseMapOutput
}

type ReleaseMap map[string]ReleaseInput

func (ReleaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Release)(nil)).Elem()
}

func (i ReleaseMap) ToReleaseMapOutput() ReleaseMapOutput {
	return i.ToReleaseMapOutputWithContext(context.Background())
}

func (i ReleaseMap) ToReleaseMapOutputWithContext(ctx context.Context) ReleaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseMapOutput)
}

type ReleaseOutput struct{ *pulumi.OutputState }

func (ReleaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Release)(nil)).Elem()
}

func (o ReleaseOutput) ToReleaseOutput() ReleaseOutput {
	return o
}

func (o ReleaseOutput) ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput {
	return o
}

type ReleaseArrayOutput struct{ *pulumi.OutputState }

func (ReleaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Release)(nil)).Elem()
}

func (o ReleaseArrayOutput) ToReleaseArrayOutput() ReleaseArrayOutput {
	return o
}

func (o ReleaseArrayOutput) ToReleaseArrayOutputWithContext(ctx context.Context) ReleaseArrayOutput {
	return o
}

func (o ReleaseArrayOutput) Index(i pulumi.IntInput) ReleaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Release {
		return vs[0].([]*Release)[vs[1].(int)]
	}).(ReleaseOutput)
}

type ReleaseMapOutput struct{ *pulumi.OutputState }

func (ReleaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Release)(nil)).Elem()
}

func (o ReleaseMapOutput) ToReleaseMapOutput() ReleaseMapOutput {
	return o
}

func (o ReleaseMapOutput) ToReleaseMapOutputWithContext(ctx context.Context) ReleaseMapOutput {
	return o
}

func (o ReleaseMapOutput) MapIndex(k pulumi.StringInput) ReleaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Release {
		return vs[0].(map[string]*Release)[vs[1].(string)]
	}).(ReleaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseInput)(nil)).Elem(), &Release{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseArrayInput)(nil)).Elem(), ReleaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseMapInput)(nil)).Elem(), ReleaseMap{})
	pulumi.RegisterOutputType(ReleaseOutput{})
	pulumi.RegisterOutputType(ReleaseArrayOutput{})
	pulumi.RegisterOutputType(ReleaseMapOutput{})
}