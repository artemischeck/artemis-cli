NAME xxx

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
URL https://xxx

#Service port if not included in URL
#PORT=[Optional]

#Request type
REQUEST POST

#Request data
DATA xxx

#Interval between each monitor. Time on seconds.
INTERVAL=300

#Local command or script to run for data
#CMD ./localscript.sh
