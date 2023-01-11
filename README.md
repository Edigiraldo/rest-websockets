# Web forum with REST and WebSockets

To run this proyect in your local environment:

- Go to the root of the proyect and then

```console
    docker compose up
```

Locally:

- To sing up:

  [POST] localhost:5050/api/v1/signup

  ```json
  {
    "email": "your@email.com",
    "password": "yourpassword"
  }
  ```

- To login:

  [POST] localhost:5050/api/v1/login

  ```json
  {
    "email": "your@email.com",
    "password": "yourpassword"
  }
  ```

- To Get user info:

  [GET] localhost:5050/api/v1/me

  Header

  ```json
  {
    "Authorization": "the.received.token"
  }
  ```

- To create a new post:

  [POST] localhost:5050/api/v1/posts

  Header

  ```json
  {
    "Authorization": "the.received.token"
  }
  ```

  Body

  ```json
  {
    "content": "The post content"
  }
  ```

- To get a post:\

  [GET] localhost:5050/api/v1/posts/{id}

  Header

  ```json
  {
    "Authorization": "the.received.token"
  }
  ```

- To update a the post content:\

  [PATCH] localhost:5050/api/v1/posts/{id}

  Header

  ```json
  {
    "Authorization": "the.received.token"
  }
  ```

  Body

  ```json
  {
    "content": "The new post content"
  }
  ```

- To delete a post:\

  [DELETE] localhost:5050/api/v1/posts/{id}

  Header

  ```json
  {
    "Authorization": "the.received.token"
  }
  ```

- To create a websocket connection:\

  localhost:5050/api/v1/ws

  Header

  ```json
  {
    "Authorization": "the.received.token"
  }
  ```
