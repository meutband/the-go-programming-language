# Problem Set

## 8.1
Modify `clock2` to accept a port number, and write a program, `clockwall`, that acts as a client of several clock servers at once, reading the times from each one and displaying the results in a table, akin to the wall of clocks seen on some business offices. If you have access to geographically distributed computers, run instances remotely; otherwise run local instances on different ports with fake time zones.

```
    $ TZ=US/Eastern ./clock2 -port 8010 &
    $ TZ=Asia/Tokyo ./clock2 -port 8020 &
    $ TZ=Europe/London ./clock2 -port 8030 &
    $ clockwall NewYork=localhost:8010 Tokyo=localhost:8020 Londin=localhost:8030
```

## 8.2
Implement a concurrent File Transfer Protocol (FTP) server. The server should interpret commands from each client such as `cd` to change directory, `ls` to list a directory, `get` to send contents of a file, and `close` to close the connection. You can use the standard `ftp` command as the client, or write your own.

## 8.3
In `netcat3`, the interface value `conn` has the concrete type `*net.TCPConn`, which represents a TCP connection. A TCP connection consists of two halves that may be closed independently using its `CloseRead` and `CloseWrite` methods. Modify the main goroutine of `netcat3` to close only the write half of the connection so that the program will continue to print the final echoes from the `reverb1` server even after the standard input has been closed. (Doing this for the `reverb2` server is harder; see Exercise 8.4.)

## 8.4
Modify the `reverb2` server to use a `sync.WaitGroup` per connection to count the number of actice `echo` goroutines. When it falls to zero, close the write half of the TCP connection as described in Exercise 8.3. Verify that your modified `netcat3` client from that exercise waits for the final echoes of multiple concurrent shouts, even after the standard input has been closed. 

## 8.5
Take an existing CPU-bound sequential program, such as the Mandelbrot program of Section 3.3 or the 3-D surface computation of Section 3.2, and execute its main loop in parallel using channels for communication. How much faster does it run on a multiprocessor machine? What is the optimal number of goroutines to use?

## 8.6
Add depth-limiting to the concurrent crawler. That is, if the user sets `-depth=3`, then only URLs reachable by at most three links will be fetched.

## 8.7
Write a concurrent program that creates a local mirror of a web site, fetching each reachable page and writing it to a directory on the local disk. Only pages within the original domain (for instance, `golang.org`) should be fetched. URLs within mirrored pages should be altered as needed so that they refer to the mirrored page, not the original.

## 8.8
Using a select statement, add a timeout to the echo server from Section 8.3 so that it disconnects any client that shouts nothing within 10 seconds.

## 8.9
Write a version of `du` that computes and periodically displays seperate totals for each of the root directories.

## 8.10
HTTP requests may be cancelled by closing the optional `Cancel` channel in the `http.Request` struct. Modify the web crawler of Section 8.6 to support cancellation.

Hint: the `http.Get` convenience function does not give you an opportunity to customize a `Request`. Instead, create the request using `http.NewRequest`, set its `Cancel` field, then perform the request by calling `http.DefaultClient.Do(req)`.

## 8.11
Following the approach of `mirroredQuery` in Section 8.4.4, implement a variant of `fetch` that requests several URLs concurrently. As soon as the first response arrives, cancel the other requests. 

## 8.12
Make the broadcaster announce the current set of clients to each new arrival. This requires that the `clients` set and the `entering` and `leaving` channels record the client name too.

## 8.13
Make the chat server disconnect idle clients, such as those that have sent no messages in the last five minutes. Hint: calling `conn.Close()` in another goroutine unblocks active `Read` calls such as the one done by `input.Scan()`.

## 8.14
Change the chat server's network protocol so that each client provides its name on entering. Use that name instead of the network address when prefixing each message with its sender's identity.

## 8.15
Failure of any client program to read data in a timely manner ultimately causes all clients to get stuck. Modify the broadcaster to skip a message rather than wait if a client writer is not ready to accept it. Alternatively, add buffering to each client's outgoing message channel so that most messages are not dropped; the broadcaster should use a non-blocking send to this channel.