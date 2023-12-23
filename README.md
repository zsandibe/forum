# Forum-moderation

## Description

This is a web forum project created to enable communication between users through posts and comments. The forum allows users to associate categories to posts, like and dislike posts and comments, and filter posts based on categories, created posts, and liked posts.


+ This project consists in creating a web forum that allows :
    - communication between users.
    - associating categories to posts.
    - liking and disliking posts and comments.
    - filtering posts.

+ This project will help you learn about:
    + The basics of web :
        - HTML
        - HTTP
         -Sessions and cookies
    + Using and setting up Docker
        - Containerizing an application
        - Compatibility/Dependency
        - Creating images
    + SQL language
        - Manipulation of databases
    + The basics of encryption


## Usage

0. Admin credentials

``` email 
admin@gmail.com
```

``` password
admin
```

1. Run server

```
$ go run cmd/main.go
```
2. Or run container

- Build image
```
$ make build
```
- Then run image
```
$ make run-img
```
- To stop
```
$ docker stop {DOCKER IMAGE ID}
```

Server will start at **http://localhost:8081** address


## Authors

@zsandibe
@ymoldabe
