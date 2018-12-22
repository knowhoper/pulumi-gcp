# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetLBIPRangesResult(object):
    """
    A collection of values returned by getLBIPRanges.
    """
    def __init__(__self__, http_ssl_tcp_internals=None, networks=None, id=None):
        if http_ssl_tcp_internals and not isinstance(http_ssl_tcp_internals, list):
            raise TypeError('Expected argument http_ssl_tcp_internals to be a list')
        __self__.http_ssl_tcp_internals = http_ssl_tcp_internals
        """
        The IP ranges used for health checks when **HTTP(S), SSL proxy, TCP proxy, and Internal load balancing** is used
        """
        if networks and not isinstance(networks, list):
            raise TypeError('Expected argument networks to be a list')
        __self__.networks = networks
        """
        The IP ranges used for health checks when **Network load balancing** is used
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_lbip_ranges():
    """
    Use this data source to access IP ranges in your firewall rules.
    
    https://cloud.google.com/compute/docs/load-balancing/health-checks#health_check_source_ips_and_firewall_rules
    """
    __args__ = dict()

    __ret__ = await pulumi.runtime.invoke('gcp:compute/getLBIPRanges:getLBIPRanges', __args__)

    return GetLBIPRangesResult(
        http_ssl_tcp_internals=__ret__.get('httpSslTcpInternals'),
        networks=__ret__.get('networks'),
        id=__ret__.get('id'))