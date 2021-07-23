// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.NetworkServices.Outputs
{

    [OutputType]
    public sealed class EdgeCacheServiceRoutingHostRule
    {
        /// <summary>
        /// A human-readable description of the resource.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The list of host patterns to match.
        /// Host patterns must be valid hostnames with optional port numbers in the format host:port. * matches any string of ([a-z0-9-.]*).
        /// The only accepted ports are :80 and :443.
        /// Hosts are matched against the HTTP Host header, or for HTTP/2 and HTTP/3, the ":authority" header, from the incoming request.
        /// </summary>
        public readonly ImmutableArray<string> Hosts;
        /// <summary>
        /// The name of the pathMatcher associated with this hostRule.
        /// </summary>
        public readonly string PathMatcher;

        [OutputConstructor]
        private EdgeCacheServiceRoutingHostRule(
            string? description,

            ImmutableArray<string> hosts,

            string pathMatcher)
        {
            Description = description;
            Hosts = hosts;
            PathMatcher = pathMatcher;
        }
    }
}