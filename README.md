# go-nagrapi

**go-nagrapi** provides a simple REST API for **Nagios** based on its `status.dat` file.

## Usage

### As a software

Software can be used as follow : `go-nagrapi -s status.dat`

### As a Docker container

It can (or must) be deployed in a Docker container as follow :

`docker run -d --name api -p 8080:8080 -v /var/log/centreon-engine:/data fallais/go-nagrapi`

**Note** : The folder that is mapped must contain the `status.dat` Nagios file.

## API

Once started, you can access this page : `http://localhost:8080/state`
