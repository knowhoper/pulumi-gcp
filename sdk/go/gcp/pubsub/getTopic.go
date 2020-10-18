// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get information about a Google Cloud Pub/Sub Topic. For more information see
// the [official documentation](https://cloud.google.com/pubsub/docs/)
// and [API](https://cloud.google.com/pubsub/docs/apis).
func LookupTopic(ctx *pulumi.Context, args *LookupTopicArgs, opts ...pulumi.InvokeOption) (*LookupTopicResult, error) {
	var rv LookupTopicResult
	err := ctx.Invoke("gcp:pubsub/getTopic:getTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTopic.
type LookupTopicArgs struct {
	// The name of the Cloud Pub/Sub Topic.
	Name string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getTopic.
type LookupTopicResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                     string                         `pulumi:"id"`
	KmsKeyName             string                         `pulumi:"kmsKeyName"`
	Labels                 map[string]string              `pulumi:"labels"`
	MessageStoragePolicies []GetTopicMessageStoragePolicy `pulumi:"messageStoragePolicies"`
	Name                   string                         `pulumi:"name"`
	Project                *string                        `pulumi:"project"`
}