LABEL Postgres DB

SERVICE_TYPE DOCKER

#Interval between each monitor. Time value in seconds.
INTERVAL 60

#Timeout delay. Time value in seconds.
TIMEOUT 10

# Docker container name
CONTAINER_NAME postgres