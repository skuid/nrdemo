# New Relic Infrastructure Debugging

The following repository contains a demo go application that reports to New
Relic that runs in a container. It also launches the New Relic Infrastructure
agent and a `ping` container that queries the demo application.

## Setup

You must have installed:

* Docker [for mac](https://www.docker.com/docker-mac) or [for windows](https://www.docker.com/docker-windows)
* [Docker-compose](https://docs.docker.com/compose/install/) (comes with docker for mac/windows)
* A New Relic license key

Once you have the docker tools installed, edit the `docker-compose.yml` file and
add the relevant New Relic License key to the Infrastructure agent and demoapp.

```yaml
version: '3'
services:
  pinger:
    image: alpine
    volumes:
    - ./ping.sh:/ping.sh
    command: /ping.sh
  demoapp:
    build:
      context: ./demoapp
    ports:
    - "8080:8080"
    environment:
      NEW_RELIC_LICENSE_KEY: '' # <<<<< Add License Key
  nrinfra:
    command: /usr/bin/newrelic-infra -config=/etc/newrelic/newrelic-infra.yml
    environment:
      NRIA_DISPLAY_NAME: 'dockerformac'
      NRIA_LICENSE_KEY: '' # <<<<< Add License Key
    image: newrelic/infrastructure:latest
    privileged: true
    network_mode: "host"
    volumes:
    - /:/host:ro
    - /etc/hostname:/etc/hostname:ro
    - /etc/os-release:/etc/os-release:ro
    - /var/run/docker.sock:/var/run/docker.sock
    - './newrelic-infra.yml:/etc/newrelic/newrelic-infra.yml'
```

After editing the file and adding the license key, run:

```bash
docker-compose up -d
```

To check the demo application is running, run any of the following commands:

```bash
docker-compose logs
# or
curl http://localhost:8080/
# or
open http://localhost:8080/
```

## Verify

After a few minutes the server `dockerformac` should show up in the New Relic
infrastucture UI, and the `demoapp` should show up in RPM.

What we're exepecting to see but are not:

1. Processes on the server in the Infrastructure UI
1. Association of the `demoapp` to the `dockerformac` server in the maps view
1. Infrastructure Inventory metadata around other docker containers

