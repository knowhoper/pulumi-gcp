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
    public sealed class EdgeCacheOriginTimeout
    {
        /// <summary>
        /// The maximum duration to wait for the origin connection to be established, including DNS lookup, TLS handshake and TCP/QUIC connection establishment.
        /// Defaults to 5 seconds. The timeout must be a value between 1s and 15s.
        /// </summary>
        public readonly string? ConnectTimeout;
        /// <summary>
        /// The maximum time across all connection attempts to the origin, including failover origins, before returning an error to the client. A HTTP 503 will be returned if the timeout is reached before a response is returned.
        /// Defaults to 5 seconds. The timeout must be a value between 1s and 15s.
        /// </summary>
        public readonly string? MaxAttemptsTimeout;
        /// <summary>
        /// The maximum duration to wait for data to arrive when reading from the HTTP connection/stream.
        /// Defaults to 5 seconds. The timeout must be a value between 1s and 30s.
        /// </summary>
        public readonly string? ResponseTimeout;

        [OutputConstructor]
        private EdgeCacheOriginTimeout(
            string? connectTimeout,

            string? maxAttemptsTimeout,

            string? responseTimeout)
        {
            ConnectTimeout = connectTimeout;
            MaxAttemptsTimeout = maxAttemptsTimeout;
            ResponseTimeout = responseTimeout;
        }
    }
}