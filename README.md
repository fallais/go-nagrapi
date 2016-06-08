# go-nagrapi

**go-nagrapi** provides a simple REST API for **Nagios** based on its `status.dat` file.

## Docker

It can (or must) be deployed in a Docker container as follow : `docker run -d --name api.1 -p 5555:8080 -v /var/log/centreon-engine/status.dat:/usr/bin/status.dat fallais/go-nagrapi`
