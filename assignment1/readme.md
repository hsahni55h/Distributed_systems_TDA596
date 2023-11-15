# Running the `main.go` File (same applies for `proxy_main.go`)

To run the `main.go` file with a specific port number (say `1278`), use the following command:

`go run main.go -p 1278`

# Building and running the `main.exe` File (same applies for `proxy_main.go`)

To build the `main.exe` file from `main.go`, use the following command:

`go build main.go`

To run the `main.exe` file with a specific port number (say `1278`), use the following command:

`main.exe -p 1278`

# Testing the Functionality of GET

To test the functionality of GET, open a command prompt and use the following examples:

1. **HTML File Types**
`curl http://127.0.0.1:1278/index.html`

2. **CSS File Types**
`curl http://127.0.0.1:1278/css/styles.css`

# Testing the Functionality of POST

To test the functionality of POST, use the following example:

`curl -X POST -H "Content-Type: text/plain" -d DEF http://127.0.0.1:1278/data/vipNames.txt`

# GET command to run n clients in parallel on a single command prompt without blocking it

To run multiple GET requests in parallel on a single command prompt without blocking it, use the following command (for Windows command prompt):

`for /L %i in (1, 1, 20) do start /b curl http://127.0.0.1:1278/data/vipNames.txt`

If you wish to store the response in a file, use the following command (for Windows command prompt):
`for /L %i in (1, 1, 20) do start /b curl http://127.0.0.1:1278/data/vipNames.txt > out_%i.txt 2> info_%i.txt`

POST command to run n clients in parallel on single command prompt without blocking it

`for %i in (GHI, JKL, MNO, PQR, STQ, UVW, XYZ) do start /b curl -X POST -H "Content-Type: text/plain" -d "%i" http://127.0.0.1:1278/data/vipNames.txt`


# Testing the Functionality of GET with proxy server
`curl -X GET 127.0.0.1:1278/index.html -x 127.0.0.1:1279`

# Testing the Functionality of GET with proxy server - shouldn't be allowed and must be handled gracefully.
`curl -X POST -H "Content-Type: text/plain" -d DEF http://127.0.0.1:1278/data/vipNames.txt -x 127.0.0.1:1279`

Expected response:
```shell
> "Only GET requests are supported"
```
