# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class FhirStore(pulumi.CustomResource):
    dataset: pulumi.Output[str]
    disable_referential_integrity: pulumi.Output[bool]
    disable_resource_versioning: pulumi.Output[bool]
    enable_history_import: pulumi.Output[bool]
    enable_update_create: pulumi.Output[bool]
    labels: pulumi.Output[dict]
    name: pulumi.Output[str]
    notification_config: pulumi.Output[dict]
    self_link: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, dataset=None, disable_referential_integrity=None, disable_resource_versioning=None, enable_history_import=None, enable_update_create=None, labels=None, name=None, notification_config=None, __name__=None, __opts__=None):
        """
        A FhirStore is a datastore inside a Healthcare dataset that conforms to the FHIR (https://www.hl7.org/fhir/STU3/)
        standard for Healthcare information exchange
        
        To get more information about FhirStore, see:
        
        * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1beta1/projects.locations.datasets.fhirStores)
        * How-to Guides
            * [Creating a FHIR store](https://cloud.google.com/healthcare/docs/how-tos/fhir)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if dataset is None:
            raise TypeError("Missing required property 'dataset'")
        __props__['dataset'] = dataset

        __props__['disable_referential_integrity'] = disable_referential_integrity

        __props__['disable_resource_versioning'] = disable_resource_versioning

        __props__['enable_history_import'] = enable_history_import

        __props__['enable_update_create'] = enable_update_create

        __props__['labels'] = labels

        __props__['name'] = name

        __props__['notification_config'] = notification_config

        __props__['self_link'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(FhirStore, __self__).__init__(
            'gcp:healthcare/fhirStore:FhirStore',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
