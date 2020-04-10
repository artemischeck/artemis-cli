UP=$(pgrep $SERVICE_NAME | wc -l)
if [ "$UP" -ne 1 ]; then
    echo "$SERVICE_NAME is down"
fi
