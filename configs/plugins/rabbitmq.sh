UP=$(pgrep rabbitmq-server | wc -l);
if [ "$UP" -ne 1 ];
then
    echo "RabbitMQ Server is down";
fi