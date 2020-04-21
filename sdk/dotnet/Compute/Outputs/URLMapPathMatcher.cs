// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class URLMapPathMatcher
    {
        public readonly string? DefaultService;
        public readonly string? Description;
        public readonly Outputs.URLMapPathMatcherHeaderAction? HeaderAction;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.URLMapPathMatcherPathRule> PathRules;
        public readonly ImmutableArray<Outputs.URLMapPathMatcherRouteRule> RouteRules;

        [OutputConstructor]
        private URLMapPathMatcher(
            string? defaultService,

            string? description,

            Outputs.URLMapPathMatcherHeaderAction? headerAction,

            string name,

            ImmutableArray<Outputs.URLMapPathMatcherPathRule> pathRules,

            ImmutableArray<Outputs.URLMapPathMatcherRouteRule> routeRules)
        {
            DefaultService = defaultService;
            Description = description;
            HeaderAction = headerAction;
            Name = name;
            PathRules = pathRules;
            RouteRules = routeRules;
        }
    }
}