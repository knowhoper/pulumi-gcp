// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Beta only
 * ## Example Usage
 * ### Basic_monitored_project
 * A basic example of a monitoring monitored project
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const primary = new gcp.monitoring.MonitoredProject("primary", {metricsScope: "my-project-name"});
 * const basic = new gcp.organizations.Project("basic", {
 *     projectId: "id",
 *     orgId: "123456789",
 * });
 * ```
 *
 * ## Import
 *
 * MonitoredProject can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:monitoring/monitoredProject:MonitoredProject default locations/global/metricsScopes/{{metrics_scope}}/projects/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:monitoring/monitoredProject:MonitoredProject default {{metrics_scope}}/{{name}}
 * ```
 */
export class MonitoredProject extends pulumi.CustomResource {
    /**
     * Get an existing MonitoredProject resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MonitoredProjectState, opts?: pulumi.CustomResourceOptions): MonitoredProject {
        return new MonitoredProject(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:monitoring/monitoredProject:MonitoredProject';

    /**
     * Returns true if the given object is an instance of MonitoredProject.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MonitoredProject {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MonitoredProject.__pulumiType;
    }

    /**
     * Output only. The time when this `MonitoredProject` was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Required. The resource name of the existing Metrics Scope that will monitor this project. Example: locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}
     */
    public readonly metricsScope!: pulumi.Output<string>;
    /**
     * Immutable. The resource name of the `MonitoredProject`. On input, the resource name includes the scoping project ID and monitored project ID. On output, it contains the equivalent project numbers. Example: `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}`
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a MonitoredProject resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MonitoredProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MonitoredProjectArgs | MonitoredProjectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MonitoredProjectState | undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["metricsScope"] = state ? state.metricsScope : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as MonitoredProjectArgs | undefined;
            if ((!args || args.metricsScope === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metricsScope'");
            }
            inputs["metricsScope"] = args ? args.metricsScope : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["createTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MonitoredProject.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MonitoredProject resources.
 */
export interface MonitoredProjectState {
    /**
     * Output only. The time when this `MonitoredProject` was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Required. The resource name of the existing Metrics Scope that will monitor this project. Example: locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}
     */
    metricsScope?: pulumi.Input<string>;
    /**
     * Immutable. The resource name of the `MonitoredProject`. On input, the resource name includes the scoping project ID and monitored project ID. On output, it contains the equivalent project numbers. Example: `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}`
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MonitoredProject resource.
 */
export interface MonitoredProjectArgs {
    /**
     * Required. The resource name of the existing Metrics Scope that will monitor this project. Example: locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}
     */
    metricsScope: pulumi.Input<string>;
    /**
     * Immutable. The resource name of the `MonitoredProject`. On input, the resource name includes the scoping project ID and monitored project ID. On output, it contains the equivalent project numbers. Example: `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}`
     */
    name?: pulumi.Input<string>;
}