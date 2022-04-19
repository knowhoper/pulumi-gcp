// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Iam.Inputs
{

    public sealed class DenyPolicyRuleDenyRuleGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// User defined CEVAL expression. A CEVAL expression is used to specify match criteria such as origin.ip, source.region_code and contents in the request header.
        /// Structure is documented below.
        /// </summary>
        [Input("denialCondition", required: true)]
        public Input<Inputs.DenyPolicyRuleDenyRuleDenialConditionGetArgs> DenialCondition { get; set; } = null!;

        [Input("deniedPermissions")]
        private InputList<string>? _deniedPermissions;

        /// <summary>
        /// The permissions that are explicitly denied by this rule. Each permission uses the format `{service-fqdn}/{resource}.{verb}`,
        /// where `{service-fqdn}` is the fully qualified domain name for the service. For example, `iam.googleapis.com/roles.list`.
        /// </summary>
        public InputList<string> DeniedPermissions
        {
            get => _deniedPermissions ?? (_deniedPermissions = new InputList<string>());
            set => _deniedPermissions = value;
        }

        [Input("deniedPrincipals")]
        private InputList<string>? _deniedPrincipals;

        /// <summary>
        /// The identities that are prevented from using one or more permissions on Google Cloud resources.
        /// </summary>
        public InputList<string> DeniedPrincipals
        {
            get => _deniedPrincipals ?? (_deniedPrincipals = new InputList<string>());
            set => _deniedPrincipals = value;
        }

        [Input("exceptionPermissions")]
        private InputList<string>? _exceptionPermissions;

        /// <summary>
        /// Specifies the permissions that this rule excludes from the set of denied permissions given by deniedPermissions.
        /// If a permission appears in deniedPermissions and in exceptionPermissions then it will not be denied.
        /// The excluded permissions can be specified using the same syntax as deniedPermissions.
        /// </summary>
        public InputList<string> ExceptionPermissions
        {
            get => _exceptionPermissions ?? (_exceptionPermissions = new InputList<string>());
            set => _exceptionPermissions = value;
        }

        [Input("exceptionPrincipals")]
        private InputList<string>? _exceptionPrincipals;

        /// <summary>
        /// The identities that are excluded from the deny rule, even if they are listed in the deniedPrincipals.
        /// For example, you could add a Google group to the deniedPrincipals, then exclude specific users who belong to that group.
        /// </summary>
        public InputList<string> ExceptionPrincipals
        {
            get => _exceptionPrincipals ?? (_exceptionPrincipals = new InputList<string>());
            set => _exceptionPrincipals = value;
        }

        public DenyPolicyRuleDenyRuleGetArgs()
        {
        }
    }
}