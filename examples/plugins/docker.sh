INSPECT=$(docker inspect -f '{{.State.Running}}' $CONTAINER_NAME)
if [ "$INSPECT" -ne true ]; then
    echo "$CONTAINER_NAME is down"
fi