// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as impart from "@impart-security/pulumi-impart";
 *
 * // Create a new external link
 * const externalLink1 = new impart.ExternalLink("externalLink1", {
 *     description: "A link to Datadog dashboard for client IP address",
 *     entity: "request",
 *     jsonPathElement: "$.client_ip.address",
 *     name: "Datadog client IP address",
 *     url: "https://app.datadoghq.com/dashboard/3tm-mpc-863?tpl_var_ClientIp=9.37.130.233",
 *     vendor: "Datadog",
 * });
 * ```
 */
export class ExternalLink extends pulumi.CustomResource {
    /**
     * Get an existing ExternalLink resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExternalLinkState, opts?: pulumi.CustomResourceOptions): ExternalLink {
        return new ExternalLink(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'impart:index/externalLink:ExternalLink';

    /**
     * Returns true if the given object is an instance of ExternalLink.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExternalLink {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExternalLink.__pulumiType;
    }

    /**
     * The description of the external link.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The entity to which the links should be applied.
     */
    public readonly entity!: pulumi.Output<string>;
    /**
     * A JSONPath to the element for which this link should apply (e.g. $.client_ip.address).
     */
    public readonly jsonPathElement!: pulumi.Output<string>;
    /**
     * The name of the external link.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of spec IDs this external link applies to (empty means all).
     */
    public readonly specIds!: pulumi.Output<string[] | undefined>;
    /**
     * The external URL template with JSONPath element variables.
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * The vendor for the external link.
     */
    public readonly vendor!: pulumi.Output<string>;

    /**
     * Create a ExternalLink resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExternalLinkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExternalLinkArgs | ExternalLinkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExternalLinkState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["entity"] = state ? state.entity : undefined;
            resourceInputs["jsonPathElement"] = state ? state.jsonPathElement : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["specIds"] = state ? state.specIds : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["vendor"] = state ? state.vendor : undefined;
        } else {
            const args = argsOrState as ExternalLinkArgs | undefined;
            if ((!args || args.entity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entity'");
            }
            if ((!args || args.jsonPathElement === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jsonPathElement'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            if ((!args || args.vendor === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vendor'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["entity"] = args ? args.entity : undefined;
            resourceInputs["jsonPathElement"] = args ? args.jsonPathElement : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["specIds"] = args ? args.specIds : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["vendor"] = args ? args.vendor : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ExternalLink.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ExternalLink resources.
 */
export interface ExternalLinkState {
    /**
     * The description of the external link.
     */
    description?: pulumi.Input<string>;
    /**
     * The entity to which the links should be applied.
     */
    entity?: pulumi.Input<string>;
    /**
     * A JSONPath to the element for which this link should apply (e.g. $.client_ip.address).
     */
    jsonPathElement?: pulumi.Input<string>;
    /**
     * The name of the external link.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of spec IDs this external link applies to (empty means all).
     */
    specIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The external URL template with JSONPath element variables.
     */
    url?: pulumi.Input<string>;
    /**
     * The vendor for the external link.
     */
    vendor?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ExternalLink resource.
 */
export interface ExternalLinkArgs {
    /**
     * The description of the external link.
     */
    description?: pulumi.Input<string>;
    /**
     * The entity to which the links should be applied.
     */
    entity: pulumi.Input<string>;
    /**
     * A JSONPath to the element for which this link should apply (e.g. $.client_ip.address).
     */
    jsonPathElement: pulumi.Input<string>;
    /**
     * The name of the external link.
     */
    name: pulumi.Input<string>;
    /**
     * A list of spec IDs this external link applies to (empty means all).
     */
    specIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The external URL template with JSONPath element variables.
     */
    url: pulumi.Input<string>;
    /**
     * The vendor for the external link.
     */
    vendor: pulumi.Input<string>;
}
