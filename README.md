# reverseproxy

This is a very simple reverse proxy which logs requests to a syslog host.

# Usage 

```
Usage of ./reverseproxy:
  -port=":8080": Port to listen for connections
  -syslog-host="localhost": Host to send UDP syslog messages
  -syslog-port="514": Port to send UDP syslog messages
  -url="https://api.tempo-db.com": URL to proxy
```

# License

Copyright (c) 2013 Mark Wolfe Licensed under the MIT license.