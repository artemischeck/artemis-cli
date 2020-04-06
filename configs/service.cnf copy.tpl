LABEL xxx

#Service type
# Options: REST, SOAP, UTIL, TELNET
SERVICE_TYPE REST

#Authentication type
AUTH_TYPE BASIC

#Auth data
AUTH_DATA xxxx:yyyy

#Request Content Type
CONTENT_TYPE application/json

#Request URL
URL https://talklift.com

#Service port if not included in URL
PORT 80

#Request type
REQUEST POST

#Request data
DATA xxx

#Interval between each monitor. Time value in seconds.
INTERVAL 60

#Timeout delay. Time value in seconds.
TIMEOUT 10

#Local command or script to run for data
#CMD ./localscript.sh
