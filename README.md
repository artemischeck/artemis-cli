# API Monitoring Client for Artemis Health Check

## Installation

### Build from source

    git clone https://github.com/artemischeck/artemischeck-cli.git
    cd artemischeck
    export GOPATH=$PWD
    go build main.go
    mv main artemischeck

### How to run

    #Check examples files for ideas
    artemischeck -dir=PATH_TO_CONFIG_FILES

### Monitoring a RESTFul service

    LABEL Service A
    
    #Service type: REST, SOAP, UTIL, TELNET
    SERVICE_TYPE REST

    #Authentication type
    AUTH_TYPE BASIC

    #Auth data
    AUTH_DATA xxxx:yyyy

    #Request Content Type
    CONTENT_TYPE application/json

    #Request URL
    URL https://myapi.com/resource/

    #Service port if not included in URL
    PORT 80

    #Request type
    REQUEST POST

    #Request data
    DATA {"hello", "world"}

    #Interval between each monitor. Time value in seconds.
    INTERVAL 300

    #Timeout delay. Time value in seconds.
    TIMEOUT 10


### Monitoring Plugin Services e.g Redis, MySQL

    LABEL Redis Server

    SERVICE_TYPE PLUGIN

    #Interval between each monitor. Time value in seconds.
    INTERVAL 60

    #Timeout delay. Time value in seconds.
    TIMEOUT 10

    # Utility service name e.g mysql, apache2, redis, nginx
    UTIL_SERVICE_NAME redis

### Telnet Service

    LABEL Telnet Test Server

    SERVICE_TYPE TELNET

    #Interval between each monitor. Time value in seconds.
    INTERVAL 60

    #Timeout delay. Time value in seconds.
    TIMEOUT 10

    HOST 8.8.8.8

    PORT 80

### Docker Container Monitoring

    LABEL Docker Container Check

    SERVICE_TYPE DOCKER

    #Interval between each monitor. Time value in seconds.
    INTERVAL 60

    #Timeout delay. Time value in seconds.
    TIMEOUT 10

    # Docker container name
    # Consider naming your containers for easy reference e.g docker run --name postgres postgres
    CONTAINER_NAME postgres


Check [Examples Folder](examples) for sample setup files.
