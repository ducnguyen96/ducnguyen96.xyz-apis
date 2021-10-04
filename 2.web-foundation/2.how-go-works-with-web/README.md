![1](./images/3.3.illustrator.png)

![2](./images/3.3.http.png)

1. Create a listening socket, listen to a port and wait for clients.
2. Accept requests from clients.
3. Handle requests, read HTTP header. If the request uses POST method, read data in the message body and pass them to handlers. Finally, socket returns response data to clients.
