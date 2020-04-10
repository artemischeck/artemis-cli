UP=$(pgrep $UTIL_SERVICE_NAME | wc -l)
if [ "$UP" -ne 1 ]; then
    echo "$UTIL_SERVICE_NAME is down"
fi
