# Lugo Frontend


The Lugo Frontend is a standalone application, but it may (and is) integrated to the Game Server.

This application listens to the events broadcast by the Game Server through the gRPC service called Broadcast.


## Gridview 

The frontend does not require pre-configuration to display the Gridview over the game field.

Use the query parameters `r` (rows) and `c` (columns) to display the grid. **Both** parameters must be defined
to the frontend display the grid.

e.g. [http://localhost:8081/?r=8&c=10](http://localhost:8081/?r=8&c=10)


