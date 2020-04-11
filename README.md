# API Monitoring Client for Artemis Health Check

## Installation

### Build from source

    git clone https://github.com/felixcheruiyot/healthcheck.git
    cd healthcheck
    export GOPATH=$PWD
    go build main.go
    mv main healthcheck

### How to run

    healthcheck -dir=PATH_TO_CONFIG_FILES #Check configs file for example

### Monitoring a RESTFul service

    LABEL Service A
    #Service type
    # Options: REST, SOAP, UTIL, TELNET
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


### Monitoring UTIL Services

    LABEL Redis

    SERVICE_TYPE SCRIPT

    #Interval between each monitor. Time value in seconds.
    INTERVAL 60

    #Timeout delay. Time value in seconds.
    TIMEOUT 10

    #Local command or script to run for data
    CMD PATH_TO_PLUGINS_FOLDER/redis.sh

Redis.sh script

    UP=$(pgrep redis | wc -l);
    if [ "$UP" -ne 1 ];
    then
        echo "Redis Server is down";
    fi

Check [Plugins Folder](configs/plugins) for exaple scripts
