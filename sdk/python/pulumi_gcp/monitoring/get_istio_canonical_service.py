# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetIstioCanonicalServiceResult',
    'AwaitableGetIstioCanonicalServiceResult',
    'get_istio_canonical_service',
]

@pulumi.output_type
class GetIstioCanonicalServiceResult:
    """
    A collection of values returned by getIstioCanonicalService.
    """
    def __init__(__self__, canonical_service=None, canonical_service_namespace=None, display_name=None, id=None, mesh_uid=None, name=None, project=None, service_id=None, telemetries=None):
        if canonical_service and not isinstance(canonical_service, str):
            raise TypeError("Expected argument 'canonical_service' to be a str")
        pulumi.set(__self__, "canonical_service", canonical_service)
        if canonical_service_namespace and not isinstance(canonical_service_namespace, str):
            raise TypeError("Expected argument 'canonical_service_namespace' to be a str")
        pulumi.set(__self__, "canonical_service_namespace", canonical_service_namespace)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if mesh_uid and not isinstance(mesh_uid, str):
            raise TypeError("Expected argument 'mesh_uid' to be a str")
        pulumi.set(__self__, "mesh_uid", mesh_uid)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if service_id and not isinstance(service_id, str):
            raise TypeError("Expected argument 'service_id' to be a str")
        pulumi.set(__self__, "service_id", service_id)
        if telemetries and not isinstance(telemetries, list):
            raise TypeError("Expected argument 'telemetries' to be a list")
        pulumi.set(__self__, "telemetries", telemetries)

    @property
    @pulumi.getter(name="canonicalService")
    def canonical_service(self) -> str:
        return pulumi.get(self, "canonical_service")

    @property
    @pulumi.getter(name="canonicalServiceNamespace")
    def canonical_service_namespace(self) -> str:
        return pulumi.get(self, "canonical_service_namespace")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="meshUid")
    def mesh_uid(self) -> str:
        return pulumi.get(self, "mesh_uid")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> Optional[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> str:
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter
    def telemetries(self) -> Sequence['outputs.GetIstioCanonicalServiceTelemetryResult']:
        return pulumi.get(self, "telemetries")


class AwaitableGetIstioCanonicalServiceResult(GetIstioCanonicalServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIstioCanonicalServiceResult(
            canonical_service=self.canonical_service,
            canonical_service_namespace=self.canonical_service_namespace,
            display_name=self.display_name,
            id=self.id,
            mesh_uid=self.mesh_uid,
            name=self.name,
            project=self.project,
            service_id=self.service_id,
            telemetries=self.telemetries)


def get_istio_canonical_service(canonical_service: Optional[str] = None,
                                canonical_service_namespace: Optional[str] = None,
                                mesh_uid: Optional[str] = None,
                                project: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIstioCanonicalServiceResult:
    """
    A Monitoring Service is the root resource under which operational aspects of a
    generic service are accessible. A service is some discrete, autonomous, and
    network-accessible unit, designed to solve an individual concern

    A monitoring Istio Canonical Service is automatically created by GCP to monitor
    Istio Canonical Services.

    To get more information about Service, see:

    * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services)
    * How-to Guides
        * [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
        * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)

    ## Example Usage
    ### Monitoring Istio Canonical Service

    ```python
    import pulumi
    import pulumi_gcp as gcp

    default = gcp.monitoring.get_istio_canonical_service(canonical_service="prometheus",
        canonical_service_namespace="istio-system",
        mesh_uid="proj-573164786102")
    ```


    :param str canonical_service: The name of the canonical service underlying this service.
           Corresponds to the destination_canonical_service_name metric label in label in Istio metrics.
    :param str canonical_service_namespace: The namespace of the canonical service underlying this service.
           Corresponds to the destination_canonical_service_namespace metric label in Istio metrics.
    :param str mesh_uid: Identifier for the mesh in which this Istio service is defined.
           Corresponds to the meshUid metric label in Istio metrics.
    :param str project: The ID of the project in which the resource belongs.
           If it is not provided, the provider project is used.
    """
    __args__ = dict()
    __args__['canonicalService'] = canonical_service
    __args__['canonicalServiceNamespace'] = canonical_service_namespace
    __args__['meshUid'] = mesh_uid
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:monitoring/getIstioCanonicalService:getIstioCanonicalService', __args__, opts=opts, typ=GetIstioCanonicalServiceResult).value

    return AwaitableGetIstioCanonicalServiceResult(
        canonical_service=__ret__.canonical_service,
        canonical_service_namespace=__ret__.canonical_service_namespace,
        display_name=__ret__.display_name,
        id=__ret__.id,
        mesh_uid=__ret__.mesh_uid,
        name=__ret__.name,
        project=__ret__.project,
        service_id=__ret__.service_id,
        telemetries=__ret__.telemetries)