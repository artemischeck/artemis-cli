UP=$(pgrep mysql | wc -l);
if [ "$UP" -ne 1 ];
then
    echo "MySQL is down";
fi