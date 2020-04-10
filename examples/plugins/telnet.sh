if nc -w $TIMEOUT -z $HOST $PORT; then
    #
else
    echo "Failed to connect to ${HOST}:${PORT}.Output: ($?)."
fi
