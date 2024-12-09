# API basics

So it sounds like a REST api (or Representational State Transfer) api is a way to send requests to a database and recieve information back. This works with a request --> response format where the end user essentially fills out a form and the api asks the database for it and boom you have your data (usually in a JSON format which I think makes this super easy for me because my project database is just a JSON file).

You can do other things too. There is a format called CRUD which stands for Create, Read, Update, Delete. These are all things I need, except I need search as well. So I would use GET for recieving data, POST for creating data, PUT for updating data, DELETE for deleting data. For search IDK what to do I think I would need a search algorithm and then run GET on the results or something?

Now here is how I think echo works based on the website. We have 3 basic concepts here: handlers, middleware, and routs. 

## Handlers

Handlers just seem like they simply handle the request. So a handler might include the info you want to pass to server via the api. In the example online we have the handler setting up the request to print "Hello World" with an OK status to tell us (the logger? middleware?) that everything went well. 

## Middleware 

In an application with no middleware, if a client sends a request, the request reaches the server and is handled by some function handler, and is sent back immediately from the server to the client. But in an application with middleware, the request made by the client passes throught stages like logging, authentication, session validation, and so on, the process the business logic. It filters wrong requests before it interacts with the business logic. 

## Routes 

Routes look like where the actual GET or PUT requests go. It looks like they make sure that the requests go to the right thing. So I would define a handler function that has my "Hello World" and status code and whatever else I need and the route takes in my function and where it should go. 

## What does this look like for me? 

Well it looks like I need to define some requests and make sure I am using middleware to just log relevant info for debugging. I think that would be enough for now. 

Lets see, so if I am making a PUT request, I would define a handler function that receives a user defined struct and PUT it in the data. I would use middleware to record what I did and the responses. 
