# go-nagrapi

**go-nagrapi** provides a simple REST API for **Nagios** based on its `status.dat` file.

## Usage

### As a software

Software can be used as follow : `go-nagrapi -s status.dat`

### As a Docker container

It can (or must) be deployed in a Docker container as follow :

`docker run -d --name api.1 -p 8080:8080 -v /var/log/centreon-engine/status.dat:/usr/bin/status.dat fallais/go-nagrapi`

## API

Once started, you can access this page : `http://localhost:8080/state`
