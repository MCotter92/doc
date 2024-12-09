# Structs: Document and documentStore
  - I believe that this is pretty much done.
# JSON (de)serialization or "(un)marshalling"
  - This functionality needs to be refactored so it can be called by the API.
# File monitoring - the event listener server
  - Not started
# a REST API - create, add, delete, update and search for "documents" in the database (JSON)
  - Not started
# CLI - should tie it all together 
  - Not started 
  - the CLI should: 
    - start and keep running the events listener server
    - start and keep running the api server
    - send requests to the api server to do things


# To-Do List

## Rest API - Echo
[x] Read about Echo 
[] Try to implement a post request to the global.JSON
  - I think that this should actually be pretty easy (famous last words). According to the echo cookbook, if I want to PUT into my global.json, I create a func that takes in my struct, does whatever formatting I need, and then pass that func into echo.GET("./global.Json", func). 

  - makeDocument() should take in a title string and then parse out the location, title and extension, store those into my struct and then marshal that into my json. Does get.POST() do the marshalling? I don't think so. I think the marshalling happens in makeDocument and get.POST() returns a pointer to an echo.Route(). 
    - echo.Route:  
        type Route struct {
          Method string `json:"method"`
          Path   string `json:"path"`
          Name   string `json:"name"`
      }

      so echo.Route is a struct that has attributes Method, Path and name. :w
