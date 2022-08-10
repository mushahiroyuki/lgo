# math_server

A  simple server for testing an upcoming project

## How to build

`make build`

or if you want to make a docker container:

`make docker-build`

## How to run

`./math_server -port=9090`

This launches the server, listening on port 9090. If you don't specify a port, it defaults to 8080.

Using docker:

`make docker-run`

Launches the server on port 8080. To customize the port:

`docker run -p 9090:9090 math_server:latest -port 9090`

## API

There's a single query parameter `expression` that expects an arithmetic expression. `(`, `)`, `+`, `-`, `*`, and `/` are supported.
Order of operations is respected.
There must be a space between each operator and number. If the expression is valid, the result is returned as a string with a status code 200.
If the expression is invalid, an error message is returned with a 400 status code.

From the command line:

`curl -G --data-urlencode "expression= ( 2 + 2 ) * -10" localhost:8080`

`curl -G --data-urlencode "expression= 2 + 2 * -10" localhost:8080`


