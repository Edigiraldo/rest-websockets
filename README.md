To run this proyect in your local environment:  
- Go to the root of the proyect and then\
`cd database`
- Run\
`docker compose up`
- Go to root file and run\
`go mod tidy`
- Finally, run the server with\
`go run main.go`

locally:
- To sing up:\
	Method: POST\
	URL: localhost:1234/signup\
	Body: { "email": "your@email.com", "password": "yourpassword"}
- To login:\
	Method: POST\
	URL: localhost:1234/login\
	Body: { "email": "your@email.com", "password": "yourpassword"}
- To Get user info:\
	Method: GET\
	URL: localhost:1234/me\
	Header: { "Authorization": "the.received.token"}
- To create a new post:\
	Method: POST\
	URL: localhost:1234/posts\
	Header: { "Authorization": "the.received.token"}\
	Body: { "content": "The post content"}
- To get a post:\
	Method: GET\
	URL: localhost:1234/posts/{id}\
	Header: { "Authorization": "the.received.token"}\
- To update a the post content:\
	Method: PATCH\
	URL: localhost:1234/posts/{id}\
	Header: { "Authorization": "the.received.token"}\
	Body: { "content": "The new post content"}
- To delete a post:\
	Method: DELETE\
	URL: localhost:1234/posts/{id}\
	Header: { "Authorization": "the.received.token"}
- To create a websocket connection:\
	URL: localhost:1234/ws\
	Header: { "Authorization": "the.received.token"}
