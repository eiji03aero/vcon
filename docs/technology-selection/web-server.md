# web server framework

# Conditions
- handle http requests and send back desired response data

# Features wanted
## Must to have
- basic http server functionality
  - routing
  - route groping
  - middleware
  - path parameter
- not so much to claim here

## Better to have
- websocket support
  - probably not a property of web-framework

# Candidates
- gin
  - https://github.com/gin-gonic/gin
  - claims to be the fastest atm
  - panic receovery
  - json validator
- echo
  - https://github.com/labstack/echo
  - variety of opt-in features
    - jwt support
- gorilla mux
  - https://github.com/gorilla/mux
  - lack of grouping
  - as official page says, this is more of a tool kit for http server, not a framework
- net/http
  - https://medium.com/@ribice/working-with-go-web-frameworks-gin-and-echo-a04ea4b65f72
  - better portability, should make dev's onboarding easier since interface is that of official, well-known
  - more codes to write
    - middleware
    - no support for grouping
  - should be more matured, less error-prone

# Conclusion
- go with gin
  - for brevity and usability, mas rapido development
