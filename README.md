# Middleware based in echo framework
---

### Build a middleware using echo framework

First of all, you should create a handler which sends how many days left until 1 Jan 2027 and response with HTTP 200 OK status code.

Secondly, build a middleware, which checks HTTP header `User-Role` presents and contains `admin` and prints `red button user detected` to the console (using default log package or any 3rd party) if so.

## Run

```shell
go run github.com/Artemiadze/Middleware-Echo/cmd/main
```

## Test

```shell
curl --location --request GET http://127.0.0.1:8080/status --header "User-Role: admin"
```

## License

You can share the source code or use it in learning purposes.