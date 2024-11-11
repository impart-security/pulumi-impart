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
 * // Create a new tag metadata
 * const example = new impart.TagMetadata("example", {
 *     name: "tag",
 *     description: "tag description",
 *     externalUrl: "http://example.com",
 *     labels: [resource.impart_label.example.slug],
 *     riskStatement: "risk statement",
 * });
 * ```
 */
export class TagMetadata extends pulumi.CustomResource {
    /**
     * Get an existing TagMetadata resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TagMetadataState, opts?: pulumi.CustomResourceOptions): TagMetadata {
        return new TagMetadata(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'impart:index/tagMetadata:TagMetadata';

    /**
     * Returns true if the given object is an instance of TagMetadata.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TagMetadata {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TagMetadata.__pulumiType;
    }

    /**
     * The description for the tag.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The external URL for the tag.
     */
    public readonly externalUrl!: pulumi.Output<string | undefined>;
    /**
     * The applied labels.
     */
    public readonly labels!: pulumi.Output<string[] | undefined>;
    /**
     * The tag name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The risk statement for the tag.
     */
    public readonly riskStatement!: pulumi.Output<string | undefined>;

    /**
     * Create a TagMetadata resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TagMetadataArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TagMetadataArgs | TagMetadataState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TagMetadataState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["externalUrl"] = state ? state.externalUrl : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["riskStatement"] = state ? state.riskStatement : undefined;
        } else {
            const args = argsOrState as TagMetadataArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["externalUrl"] = args ? args.externalUrl : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["riskStatement"] = args ? args.riskStatement : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TagMetadata.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TagMetadata resources.
 */
export interface TagMetadataState {
    /**
     * The description for the tag.
     */
    description?: pulumi.Input<string>;
    /**
     * The external URL for the tag.
     */
    externalUrl?: pulumi.Input<string>;
    /**
     * The applied labels.
     */
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The tag name.
     */
    name?: pulumi.Input<string>;
    /**
     * The risk statement for the tag.
     */
    riskStatement?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TagMetadata resource.
 */
export interface TagMetadataArgs {
    /**
     * The description for the tag.
     */
    description?: pulumi.Input<string>;
    /**
     * The external URL for the tag.
     */
    externalUrl?: pulumi.Input<string>;
    /**
     * The applied labels.
     */
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The tag name.
     */
    name: pulumi.Input<string>;
    /**
     * The risk statement for the tag.
     */
    riskStatement?: pulumi.Input<string>;
}
