UP=$(pgrep $SERVICE_NAME | wc -l)
if [ "$UP" -lt 1 ]; then
    echo "$SERVICE_NAME is down"
fi
