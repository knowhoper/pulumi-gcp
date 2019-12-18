// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type IamAuditConfig struct {
	s *pulumi.ResourceState
}

// NewIamAuditConfig registers a new resource with the given unique name, arguments, and options.
func NewIamAuditConfig(ctx *pulumi.Context,
	name string, args *IamAuditConfigArgs, opts ...pulumi.ResourceOpt) (*IamAuditConfig, error) {
	if args == nil || args.AuditLogConfigs == nil {
		return nil, errors.New("missing required argument 'AuditLogConfigs'")
	}
	if args == nil || args.OrgId == nil {
		return nil, errors.New("missing required argument 'OrgId'")
	}
	if args == nil || args.Service == nil {
		return nil, errors.New("missing required argument 'Service'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["auditLogConfigs"] = nil
		inputs["orgId"] = nil
		inputs["service"] = nil
	} else {
		inputs["auditLogConfigs"] = args.AuditLogConfigs
		inputs["orgId"] = args.OrgId
		inputs["service"] = args.Service
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:organizations/iamAuditConfig:IamAuditConfig", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IamAuditConfig{s: s}, nil
}

// GetIamAuditConfig gets an existing IamAuditConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamAuditConfig(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IamAuditConfigState, opts ...pulumi.ResourceOpt) (*IamAuditConfig, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["auditLogConfigs"] = state.AuditLogConfigs
		inputs["etag"] = state.Etag
		inputs["orgId"] = state.OrgId
		inputs["service"] = state.Service
	}
	s, err := ctx.ReadResource("gcp:organizations/iamAuditConfig:IamAuditConfig", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IamAuditConfig{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IamAuditConfig) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IamAuditConfig) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
func (r *IamAuditConfig) AuditLogConfigs() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["auditLogConfigs"])
}

func (r *IamAuditConfig) Etag() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["etag"])
}

// The numeric ID of the organization in which you want to manage the audit logging config.
func (r *IamAuditConfig) OrgId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["orgId"])
}

// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
func (r *IamAuditConfig) Service() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["service"])
}

// Input properties used for looking up and filtering IamAuditConfig resources.
type IamAuditConfigState struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs interface{}
	Etag interface{}
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId interface{}
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service interface{}
}

// The set of arguments for constructing a IamAuditConfig resource.
type IamAuditConfigArgs struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs interface{}
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId interface{}
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service interface{}
}