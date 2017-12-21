roles-ws
=========

API that interfaces with gitlab, control and serve roles (ansible) informations



Endpoints
----------------

/api/v1/roles
/api/v1/roles/:id/meta/:version
/api/v1/roles/:id/params/:version



Endpoint Response Examples
----------------

$ curl http://127.0.0.1:8000/api/v1/roles

```[
    {
        "ID": 161,
        "Name": "openstack-cinder",
        "Url": "https://gitlab.home/ansible-roles/openstack-cinder.git",
        "Versions": [
            "0.0.9",
            "0.0.8",
            "0.0.7",
            "0.0.6",
            "0.0.5",
            "0.0.4",
            "0.0.3",
            "0.0.2",
            "0.0.13",
            "0.0.12",
            "0.0.11",
            "0.0.10",
            "0.0.1"
        ]
    },
    {
        "ID": 160,
        "Name": "openstack-neutron-controller",
        "Url": "https://gitlab.home/ansible-roles/openstack-neutron-controller.git",
        "Versions": [
            "0.0.5",
            "0.0.4",
            "0.0.3",
            "0.0.2",
            "0.0.1"
        ]
    },
    {
        "ID": 152,
        "Name": "openstack-nova-controller",
        "Url": "https://gitlab.home/ansible-roles/openstack-nova-controller.git",
        "Versions": [
            "0.0.9",
            "0.0.8",
            "0.0.7",
            "0.0.6",
            "0.0.5",
            "0.0.4",
            "0.0.3",
            "0.0.2",
            "0.0.12",
            "0.0.11",
            "0.0.10",
            "0.0.1"
        ]
    },
    {
        "ID": 148,
        "Name": "rabbitmq",
        "Url": "https://gitlab.home/ansible-roles/rabbitmq.git",
        "Versions": [
            "0.1.1",
            "0.1.0",
            "0.0.4",
            "0.0.3",
            "0.0.2",
            "0.0.1"
        ]
    },
    {
        "ID": 147,
        "Name": "openstack-glance",
        "Url": "https://gitlab.home/ansible-roles/openstack-glance.git",
        "Versions": [
            "mitaka_1.0.7",
            "mitaka_1.0.6",
            "mitaka_1.0.5",
            "mitaka_1.0.4",
            "mitaka_1.0.3",
            "mitaka_1.0.2",
            "mitaka_1.0.1",
            "mitaka_1.0.0",
            "2.1.5",
            "2.1.4",
            "2.1.3",
            "2.1.2",
            "2.1.1",
            "2.1.0",
            "2.0.1",
            "2.0.0",
            "1.0.0"
        ]
    },
    {
        "ID": 143,
        "Name": "openstack-keystone",
        "Url": "https://gitlab.home/ansible-roles/openstack-keystone.git",
        "Versions": [
            "mitaka_1.0.6",
            "mitaka_1.0.5",
            "mitaka_1.0.4",
            "mitaka_1.0.3",
            "mitaka_1.0.2",
            "mitaka_1",
            "2.1.3",
            "2.1.2",
            "2.1.1",
            "2.1.0",
            "2.0.1",
            "2.0.0",
            "1.0.0"
        ]
    },
    {
        "ID": 134,
        "Name": "mysql-galaxy",
        "Url": "https://gitlab.home/ansible-roles/mysql-galaxy.git",
        "Versions": null
    },
    {
        "ID": 130,
        "Name": "php-fpm",
        "Url": "https://gitlab.home/ansible-roles/php-fpm.git",
        "Versions": [
            "1.0.0"
        ]
    },
    {
        "ID": 129,
        "Name": "nginx",
        "Url": "https://gitlab.home/ansible-roles/nginx.git",
        "Versions": [
            "1.0.0"
        ]
    },
    {
        "ID": 128,
        "Name": "vimbadmin",
        "Url": "https://gitlab.home/ansible-roles/vimbadmin.git",
        "Versions": [
            "1.0.3",
            "1.0.2",
            "1.0.1",
            "1.0.0"
        ]
    },
    {
        "ID": 127,
        "Name": "roundcube",
        "Url": "https://gitlab.home/ansible-roles/roundcube.git",
        "Versions": [
            "1.0.5",
            "1.0.4",
            "1.0.3",
            "1.0.2",
            "1.0.1",
            "1.0.0"
        ]
    },
    {
        "ID": 126,
        "Name": "dovecot",
        "Url": "https://gitlab.home/ansible-roles/dovecot.git",
        "Versions": [
            "1.0.4",
            "1.0.3",
            "1.0.2",
            "1.0.1",
            "1.0.0"
        ]
    },
    {
        "ID": 125,
        "Name": "openvpn-server",
        "Url": "https://gitlab.home/ansible-roles/openvpn-server.git",
        "Versions": [
            "1.0.1",
            "1.0.0"
        ]
    },
    {
        "ID": 124,
        "Name": "mysql",
        "Url": "https://gitlab.home/ansible-roles/mysql.git",
        "Versions": [
            "1.0.7",
            "1.0.6",
            "1.0.5",
            "1.0.4",
            "1.0.3",
            "1.0.2",
            "1.0.1",
            "1.0.0"
        ]
    },
    {
        "ID": 123,
        "Name": "postfix",
        "Url": "https://gitlab.home/ansible-roles/postfix.git",
        "Versions": [
            "1.0.8",
            "1.0.7",
            "1.0.6",
            "1.0.5",
            "1.0.4",
            "1.0.3",
            "1.0.2",
            "1.0.1",
            "1.0.0"
        ]
    },
    {
        "ID": 122,
        "Name": "gdnsd",
        "Url": "https://gitlab.home/ansible-roles/gdnsd.git",
        "Versions": [
            "1.0.0"
        ]
    },
    {
        "ID": 73,
        "Name": "dspam",
        "Url": "https://gitlab.home/ansible-roles/dspam.git",
        "Versions": [
            "1.0.0"
        ]
    }
]
```

$ curl http://127.0.0.1:8000/api/v1/roles/161/meta/0.0.13

```{
    "author": "Davide Guerri",
    "categories": [
        "cloud"
    ],
    "company": "Hewlett-Packard Development Company, L.P.",
    "dependencies": [],
    "description": null,
    "license": "Apache",
    "min_ansible_version": 1.7,
    "platforms": [
        {
            "name": "Ubuntu"
        }
    ],
    "versions": [
        "trusty"
    ]
}
```

$ curl http://127.0.0.1:8000/api/v1/roles/161/params/1.00

```{
    "keystone_admin_port": 35357,
    "keystone_auth_type": "password",
    "keystone_hostname": "localhost",
    "keystone_port": 5000,
    "keystone_protocol": "http",
    "keystone_version": "v2.0",
    "openstack_cinder_auth_strategy": "keystone",
    "openstack_cinder_backends": "nfs",
    "openstack_cinder_create_db": false,
    "openstack_cinder_create_db_user": false,
    "openstack_cinder_database_url": "sqlite:////var/lib/cinder/cinder.sqlite",
    "openstack_cinder_db_host": "127.0.0.1",
    "openstack_cinder_db_name": "cinder",
    "openstack_cinder_db_passwd": "cinder",
    "openstack_cinder_db_user": "cinder",
    "openstack_cinder_ip": "{{ ansible_default_ipv4.address }}",
    "openstack_cinder_lock_path": "{{ openstack_cinder_state_path }}/lock",
    "openstack_cinder_log_dir": "/var/log/cinder",
    "openstack_cinder_memcached_servers": "127.0.0.1:11211",
    "openstack_cinder_nfs_mount_attempts": 5,
    "openstack_cinder_nfs_mount_options": "vers=3",
    "openstack_cinder_nfs_mount_point_base": "{{ openstack_cinder_state_path }}/mnt",
    "openstack_cinder_nfs_shares": [],
    "openstack_cinder_nfs_shares_config_file": "/etc/cinder/nfs_shares",
    "openstack_cinder_nfs_sparsed_volumes": true,
    "openstack_cinder_osapi_volume_workers": 2,
    "openstack_cinder_passwd": 123,
    "openstack_cinder_port": 8776,
    "openstack_cinder_rabbitmq_hostname": "localhost",
    "openstack_cinder_rabbitmq_passwd": 123,
    "openstack_cinder_rabbitmq_username": "openstack",
    "openstack_cinder_state_path": "/var/lib/cinder",
    "openstack_cinder_user": "cinder",
    "openstack_cinder_volume_driver": "cinder.volume.drivers.nfs.NfsDriver"
}
```
