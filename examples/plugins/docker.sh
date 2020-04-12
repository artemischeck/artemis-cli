if [ $(docker inspect -f '{{.State.Running}}' $CONTAINER_NAME) = "true" ]; then
    # echo yup
else
    echo "$CONTAINER_NAME is down"
fi
