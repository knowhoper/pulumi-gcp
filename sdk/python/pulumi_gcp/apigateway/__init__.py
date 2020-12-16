# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .api import *
from .api_config import *
from .api_config_iam_binding import *
from .api_config_iam_member import *
from .api_config_iam_policy import *
from .api_iam_binding import *
from .api_iam_member import *
from .api_iam_policy import *
from .gateway import *
from .gateway_iam_binding import *
from .gateway_iam_member import *
from .gateway_iam_policy import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "gcp:apigateway/api:Api":
                return Api(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:apigateway/apiConfig:ApiConfig":
                return ApiConfig(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:apigateway/apiConfigIamBinding:ApiConfigIamBinding":
                return ApiConfigIamBinding(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:apigateway/apiConfigIamMember:ApiConfigIamMember":
                return ApiConfigIamMember(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:apigateway/apiConfigIamPolicy:ApiConfigIamPolicy":
                return ApiConfigIamPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:apigateway/apiIamBinding:ApiIamBinding":
                return ApiIamBinding(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:apigateway/apiIamMember:ApiIamMember":
                return ApiIamMember(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:apigateway/apiIamPolicy:ApiIamPolicy":
                return ApiIamPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:apigateway/gateway:Gateway":
                return Gateway(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:apigateway/gatewayIamBinding:GatewayIamBinding":
                return GatewayIamBinding(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:apigateway/gatewayIamMember:GatewayIamMember":
                return GatewayIamMember(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp:apigateway/gatewayIamPolicy:GatewayIamPolicy":
                return GatewayIamPolicy(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("gcp", "apigateway/api", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "apigateway/apiConfig", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "apigateway/apiConfigIamBinding", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "apigateway/apiConfigIamMember", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "apigateway/apiConfigIamPolicy", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "apigateway/apiIamBinding", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "apigateway/apiIamMember", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "apigateway/apiIamPolicy", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "apigateway/gateway", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "apigateway/gatewayIamBinding", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "apigateway/gatewayIamMember", _module_instance)
    pulumi.runtime.register_resource_module("gcp", "apigateway/gatewayIamPolicy", _module_instance)

_register_module()