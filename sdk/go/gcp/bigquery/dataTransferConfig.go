// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/bigquery_data_transfer_config.html.markdown.
type DataTransferConfig struct {
	s *pulumi.ResourceState
}

// NewDataTransferConfig registers a new resource with the given unique name, arguments, and options.
func NewDataTransferConfig(ctx *pulumi.Context,
	name string, args *DataTransferConfigArgs, opts ...pulumi.ResourceOpt) (*DataTransferConfig, error) {
	if args == nil || args.DataSourceId == nil {
		return nil, errors.New("missing required argument 'DataSourceId'")
	}
	if args == nil || args.DestinationDatasetId == nil {
		return nil, errors.New("missing required argument 'DestinationDatasetId'")
	}
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.Params == nil {
		return nil, errors.New("missing required argument 'Params'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["dataRefreshWindowDays"] = nil
		inputs["dataSourceId"] = nil
		inputs["destinationDatasetId"] = nil
		inputs["disabled"] = nil
		inputs["displayName"] = nil
		inputs["location"] = nil
		inputs["params"] = nil
		inputs["project"] = nil
		inputs["schedule"] = nil
	} else {
		inputs["dataRefreshWindowDays"] = args.DataRefreshWindowDays
		inputs["dataSourceId"] = args.DataSourceId
		inputs["destinationDatasetId"] = args.DestinationDatasetId
		inputs["disabled"] = args.Disabled
		inputs["displayName"] = args.DisplayName
		inputs["location"] = args.Location
		inputs["params"] = args.Params
		inputs["project"] = args.Project
		inputs["schedule"] = args.Schedule
	}
	inputs["name"] = nil
	s, err := ctx.RegisterResource("gcp:bigquery/dataTransferConfig:DataTransferConfig", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DataTransferConfig{s: s}, nil
}

// GetDataTransferConfig gets an existing DataTransferConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataTransferConfig(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DataTransferConfigState, opts ...pulumi.ResourceOpt) (*DataTransferConfig, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["dataRefreshWindowDays"] = state.DataRefreshWindowDays
		inputs["dataSourceId"] = state.DataSourceId
		inputs["destinationDatasetId"] = state.DestinationDatasetId
		inputs["disabled"] = state.Disabled
		inputs["displayName"] = state.DisplayName
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["params"] = state.Params
		inputs["project"] = state.Project
		inputs["schedule"] = state.Schedule
	}
	s, err := ctx.ReadResource("gcp:bigquery/dataTransferConfig:DataTransferConfig", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DataTransferConfig{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DataTransferConfig) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DataTransferConfig) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *DataTransferConfig) DataRefreshWindowDays() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["dataRefreshWindowDays"])
}

func (r *DataTransferConfig) DataSourceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dataSourceId"])
}

func (r *DataTransferConfig) DestinationDatasetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["destinationDatasetId"])
}

func (r *DataTransferConfig) Disabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["disabled"])
}

func (r *DataTransferConfig) DisplayName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["displayName"])
}

func (r *DataTransferConfig) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

func (r *DataTransferConfig) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *DataTransferConfig) Params() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["params"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *DataTransferConfig) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *DataTransferConfig) Schedule() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["schedule"])
}

// Input properties used for looking up and filtering DataTransferConfig resources.
type DataTransferConfigState struct {
	DataRefreshWindowDays interface{}
	DataSourceId interface{}
	DestinationDatasetId interface{}
	Disabled interface{}
	DisplayName interface{}
	Location interface{}
	Name interface{}
	Params interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Schedule interface{}
}

// The set of arguments for constructing a DataTransferConfig resource.
type DataTransferConfigArgs struct {
	DataRefreshWindowDays interface{}
	DataSourceId interface{}
	DestinationDatasetId interface{}
	Disabled interface{}
	DisplayName interface{}
	Location interface{}
	Params interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Schedule interface{}
}