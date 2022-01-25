// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * The NetworkConnectivity Hub resource
 *
 * ## Example Usage
 * ### Basic_hub
 * A basic test of a networkconnectivity hub
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const primary = new gcp.networkconnectivity.Hub("primary", {
 *     description: "A sample hub",
 *     labels: {
 *         "label-one": "value-one",
 *     },
 *     project: "my-project-name",
 * });
 * ```
 *
 * ## Import
 *
 * Hub can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:networkconnectivity/hub:Hub default projects/{{project}}/locations/global/hubs/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:networkconnectivity/hub:Hub default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:networkconnectivity/hub:Hub default {{name}}
 * ```
 */
export class Hub extends pulumi.CustomResource {
    /**
     * Get an existing Hub resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HubState, opts?: pulumi.CustomResourceOptions): Hub {
        return new Hub(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:networkconnectivity/hub:Hub';

    /**
     * Returns true if the given object is an instance of Hub.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Hub {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Hub.__pulumiType;
    }

    /**
     * Output only. The time the hub was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * An optional description of the hub.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Immutable. The name of the hub. Hub names must be unique. They use the following form: `projects/{project_number}/locations/global/hubs/{hub_id}`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project for the resource
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The VPC network associated with this hub's spokes. All of the VPN tunnels, VLAN attachments, and router appliance
     * instances referenced by this hub's spokes must belong to this VPC network. This field is read-only. Network Connectivity
     * Center automatically populates it based on the set of spokes attached to the hub.
     */
    public /*out*/ readonly routingVpcs!: pulumi.Output<outputs.networkconnectivity.HubRoutingVpc[]>;
    /**
     * Output only. The current lifecycle state of this hub. Possible values: STATE_UNSPECIFIED, CREATING, ACTIVE, DELETING
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Output only. The Google-generated UUID for the hub. This value is unique across all hub resources. If a hub is deleted
     * and another with the same name is created, the new hub is assigned a different unique_id.
     */
    public /*out*/ readonly uniqueId!: pulumi.Output<string>;
    /**
     * Output only. The time the hub was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Hub resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: HubArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HubArgs | HubState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HubState | undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["routingVpcs"] = state ? state.routingVpcs : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["uniqueId"] = state ? state.uniqueId : undefined;
            inputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as HubArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["routingVpcs"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["uniqueId"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Hub.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Hub resources.
 */
export interface HubState {
    /**
     * Output only. The time the hub was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * An optional description of the hub.
     */
    description?: pulumi.Input<string>;
    /**
     * Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Immutable. The name of the hub. Hub names must be unique. They use the following form: `projects/{project_number}/locations/global/hubs/{hub_id}`
     */
    name?: pulumi.Input<string>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
    /**
     * The VPC network associated with this hub's spokes. All of the VPN tunnels, VLAN attachments, and router appliance
     * instances referenced by this hub's spokes must belong to this VPC network. This field is read-only. Network Connectivity
     * Center automatically populates it based on the set of spokes attached to the hub.
     */
    routingVpcs?: pulumi.Input<pulumi.Input<inputs.networkconnectivity.HubRoutingVpc>[]>;
    /**
     * Output only. The current lifecycle state of this hub. Possible values: STATE_UNSPECIFIED, CREATING, ACTIVE, DELETING
     */
    state?: pulumi.Input<string>;
    /**
     * Output only. The Google-generated UUID for the hub. This value is unique across all hub resources. If a hub is deleted
     * and another with the same name is created, the new hub is assigned a different unique_id.
     */
    uniqueId?: pulumi.Input<string>;
    /**
     * Output only. The time the hub was last updated.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Hub resource.
 */
export interface HubArgs {
    /**
     * An optional description of the hub.
     */
    description?: pulumi.Input<string>;
    /**
     * Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Immutable. The name of the hub. Hub names must be unique. They use the following form: `projects/{project_number}/locations/global/hubs/{hub_id}`
     */
    name?: pulumi.Input<string>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
}