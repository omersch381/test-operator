.. _prerequisites:

Prerequisites
=============

First, you will need an OpenStack Platform running on OpenShift. See, the
`ci-framework documentation <https://ci-framework.readthedocs.io/en/latest/>`_
to get you started. It will get you through the installation of such environment.

After the installations is completed, you can source the credentials to the
environment as follows:

.. code-block:: bash

    eval $(${HOME}/ci-framework-data/bin/crc oc-env)
    export KUBECONFIG="${HOME}/.crc/machines/crc/kubeconfig"
    oc login -u kubeadmin -p 12345678 https://api.crc.testing:6443

Once you source the credentials, please check that you are inside the
openstack project.

.. code-block:: bash

   oc project openstack
   Now using project "openstack" on server "https://api.crc.testing:6443".

.. note::
    12345678 is a default password set by ci-framework
