To run this proyect in your local environment:  
- Go to the root of the proyect and then\
`cd database`
- Build the docker image for postgress db\
`docker build . -t rest-web-sockets-postgress-db`
- Run your docker container\
`docker run -p 4321:5432 rest-web-sockets-postgress-db`
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
