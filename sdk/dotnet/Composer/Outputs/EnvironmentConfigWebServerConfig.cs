// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Composer.Outputs
{

    [OutputType]
    public sealed class EnvironmentConfigWebServerConfig
    {
        /// <summary>
        /// Machine type on which Airflow web server is running. It has to be one of: composer-n1-webserver-2,
        /// composer-n1-webserver-4 or composer-n1-webserver-8.
        /// Value custom is returned only in response, if Airflow web server parameters were
        /// manually changed to a non-standard values.
        /// </summary>
        public readonly string MachineType;

        [OutputConstructor]
        private EnvironmentConfigWebServerConfig(string machineType)
        {
            MachineType = machineType;
        }
    }
}