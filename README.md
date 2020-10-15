# README update from Julian

## Architecture

I tried to separate the modules in the design that different engineers can focus on different parts of the project 
in different modules. Different files allow for less merge conflicts and more modular focus for the developer.

The coder package - I could have named it codex, but decided to send this to you sooner rather than later.

I hope I didn't do the binary API part of the task incorrectly, I didn't need to use the spew debugger.

## Future Work

I saw the video on concurrency from Rob Pike - https://www.youtube.com/watch?v=f6kdp27TYZs
I considered using the cool looking chan concurrency technique, but in the interest of time and the words of 
Rob Pike: "Don't overdo it: Sometimes all you need is a reference counter" - minute 38:47

## Final notes

I implemented this in a repo, which can show you some timestamps. I was doing this task concurrently with other 
tasks in my life, and also learning Go. Please don't be too critical. 

Thanks for this career opportunity.

# "Simple Redis"

## Abstract

We want to create a service that allows for remotely setting and getting strings -- a very simplified [Redis][3].
The service uses a binary protocol to communicate with clients.

## Service

We want to create a command line service that could be started like `go run main.go`

The service should be a TCP socket server.  All communication is done over the TCP socket.  

The service should listen to port `5566` as a fixed port for listening/handling connections.

The server must handle multiple clients at the sample time in a thread-safe manner (clients may attempt 
simultaneous operations)

Once a client is connected, it will communicate with the server using the protocol.

## Protocol

The protocol is binary with no built in delimiters.

We only have 1 type of data, which is a *string*.

Strings are:
- 2 byte `length`, in big endian order (network byte order), of maximum length [MaxUint16][1]
- `length` bytes of content data.

As an example, the 8 byte ASCII string `hi there` would be written to the socket as 10 bytes (2 bytes header, 8 bytes content):

| byte# | `0`  | `1`  | `2`  | `3`  | `4`  | `5` | `6`  | `7`  | `8`  | `9`  |
|-------|------|------|------|------|------|-----|------|------|------|------|
| Hex   | `00` | `08` | `68` | `69` | `20` | `74`| `68` | `65` | `72` | `65` |
| ASCII | ` `  | ` `  | `h`  | `i`  | ` `  | `t` | `h`  | `e`  | `r`  | `e`  |

String content is always considered to be UTF-8.

## API

The API is a simple RPC model:  the server waits for commands from the client, acts upon them, and returns a 
result to the client.

| Command   | Arguments     | Return result           | Description                                                                                |
|-----------|---------------|-------------------------|--------------------------------------------------------------------------------------------|
| `SET`     | `key`,`value` | `OK` string             | Sets a given `key` to have `value`                                                         |
| `GET`     | `key`         | `value` of that `key`   | Retrieves a value for a given `key`.  If the value was never set, return `""` blank string |

- Commands are written as a sequence of strings
- Concurrent setters/getters are supported

### Examples

To set a key `greeting` to `hi there`, we would write three strings:

1. `SET`
2. `greeting`
3. `hi there`

and then receive back a single string:

1. `OK`

To get the value of a key, like the one we just set, we would write two strings:

1. `GET`
2. `greeting`

and then receive back the single string:

1. `hi there`

## Details and rules

- You cannot use external data sources: Don't use Redis or any other external database or storage mechanism.
- You cannot use external protocols: Don't use Protobuf or GRPC -- build your own protocol encoders.

### Testing and practical considerations

Demonstrate the server works with a client that issues commands to it.

## References 

- [Golang Math Package Constants][1] : Maximum value for `MaxUint16`
- [go-spew][2] : Tool for printing out binary data for easy debugging
- [redis][3] : in memory key value store

[1]: https://golang.org/pkg/math/#pkg-constants
[2]: https://github.com/davecgh/go-spew
[3]: https://redis.io/