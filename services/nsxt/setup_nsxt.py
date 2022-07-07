"""
Setup file for vapi client bindings
"""
__author__ = 'VMware, Inc.'
__copyright__ = 'Copyright 2018 VMware, Inc.  All rights reserved. -- VMware Confidential'

import os
try:
    from setuptools import setup, find_packages
except ImportError:
    from distribute_setup import use_setuptools
    use_setuptools()
    from setuptools import setup, find_packages


setup(
    name='nsxt_client_bindings',
    version=os.environ["SDK_BINDING_VERSION"],
    author=__author__,
    description='Client bindings for NSX-T',
    packages=find_packages(),
    classifiers=[],
    install_requires=[
        'setuptools',
        'vapi_common_client',
        'vapi_runtime',
    ]
)