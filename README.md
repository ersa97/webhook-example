# webhook-example

this is a super simple example of how webhook works 

## How it works?
- Let's try and pretend that there is a user subscribing to an app (client)
- The client redirects the payment to the server
- But how does the client know that the user has completed his/her payment and will have to be granted additional access?
- Simple!
- We use webhook (an HTTP-based callback function) from the server to the client (now the server is the one that hits the client's API) <span style=" color: #B10DC9;">http://localhost:8080/webhook</span>
- Now basically the server is becoming the client and vice versa

## Get started
- Please run the client and server separately
- Try and hit the endpoint from the server side
- The callback message will be shown on the client-side
