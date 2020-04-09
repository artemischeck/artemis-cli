UP=$(pgrep redis | wc -l);
if [ "$UP" -ne 1 ];
then
    echo "Redis Server is down";
fi