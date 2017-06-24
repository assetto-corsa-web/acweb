# go-resp

This module implements a simple generic response type, which can be used to respond to client requests for REST endpoints. It supports three types of responses:

1. success
2. failure
3. error

Additionally they all have a response code, a short message (error description) and custom data, which is optional. The HTTP status code can be configured globally.

## Installation

```
go get github.com/DeKugelschieber/go-resp
```

## General usage

To send a response, just call one of the three functions:

```
import "github.com/DeKugelschieber/go-resp"
// ...

// code 0, message "success" and no custom data
resp.Success(w, 0, "success", nil)

// code 1, message "wrong input" and fields that were missing/empty
resp.Failure(w, 1, "wrong input", MissingFields{/*...*/})

// code 2, the technical error message, no custom data
resp.Error(w, 2, err.Error(), nil)
```

I recommend to use 0 as "no error" code, but you can also have multiple success responses and use the type to identify a successful response. Within Go the type can be identified by *resp.SUCCESS*, *resp.FAILURE* and *resp.ERROR*. Don't forget to return within the http handler after you've called the response function.

The final JSON object will look like:

```
{
    "type": 0|1|2,
    "code": number,
    "msg": string,
    "data": null|object
}
```

To change the HTTP status code, assign the following three variables to the desired status codes:

```
SuccessHttpCode
FailureHttpCode
ErrorHttpCode
```

By default, they are all set to *http.StatusOK*. I recomend to set the HTTP status codes for technical errors.

For the full documentation please visit https://godoc.org/github.com/DeKugelschieber/go-resp.

## License

MIT
