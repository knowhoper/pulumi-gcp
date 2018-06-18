// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Three different resources help you manage your IAM policy for pubsub topic. Each of these resources serves a different use case:
// 
// * `google_pubsub_topic_iam_policy`: Authoritative. Sets the IAM policy for the topic and replaces any existing policy already attached.
// * `google_pubsub_topic_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the topic are preserved.
// * `google_pubsub_topic_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the topic are preserved.
// 
// ~> **Note:** `google_pubsub_topic_iam_policy` **cannot** be used in conjunction with `google_pubsub_topic_iam_binding` and `google_pubsub_topic_iam_member` or they will fight over what your policy should be.
// 
// ~> **Note:** `google_pubsub_topic_iam_binding` resources **can be** used in conjunction with `google_pubsub_topic_iam_member` resources **only if** they do not grant privilege to the same role.
type TopicIAMMember struct {
	s *pulumi.ResourceState
}

// NewTopicIAMMember registers a new resource with the given unique name, arguments, and options.
func NewTopicIAMMember(ctx *pulumi.Context,
	name string, args *TopicIAMMemberArgs, opts ...pulumi.ResourceOpt) (*TopicIAMMember, error) {
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.Topic == nil {
		return nil, errors.New("missing required argument 'Topic'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["member"] = nil
		inputs["project"] = nil
		inputs["role"] = nil
		inputs["topic"] = nil
	} else {
		inputs["member"] = args.Member
		inputs["project"] = args.Project
		inputs["role"] = args.Role
		inputs["topic"] = args.Topic
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:pubsub/topicIAMMember:TopicIAMMember", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TopicIAMMember{s: s}, nil
}

// GetTopicIAMMember gets an existing TopicIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopicIAMMember(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TopicIAMMemberState, opts ...pulumi.ResourceOpt) (*TopicIAMMember, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["member"] = state.Member
		inputs["project"] = state.Project
		inputs["role"] = state.Role
		inputs["topic"] = state.Topic
	}
	s, err := ctx.ReadResource("gcp:pubsub/topicIAMMember:TopicIAMMember", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TopicIAMMember{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *TopicIAMMember) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *TopicIAMMember) ID() *pulumi.IDOutput {
	return r.s.ID
}

// (Computed) The etag of the topic's IAM policy.
func (r *TopicIAMMember) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

func (r *TopicIAMMember) Member() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["member"])
}

// The project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *TopicIAMMember) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The role that should be applied. Only one
// `google_pubsub_topic_iam_binding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *TopicIAMMember) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// The topic name or id to bind to attach IAM policy to.
func (r *TopicIAMMember) Topic() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["topic"])
}

// Input properties used for looking up and filtering TopicIAMMember resources.
type TopicIAMMemberState struct {
	// (Computed) The etag of the topic's IAM policy.
	Etag interface{}
	Member interface{}
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `google_pubsub_topic_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// The topic name or id to bind to attach IAM policy to.
	Topic interface{}
}

// The set of arguments for constructing a TopicIAMMember resource.
type TopicIAMMemberArgs struct {
	Member interface{}
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `google_pubsub_topic_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// The topic name or id to bind to attach IAM policy to.
	Topic interface{}
}