# Running the `main.go` File

To run the `main.go` file with a specific port number, use the following command:

go run main.go -p 1278

# Testing the Functionality of GET

To test the functionality of GET, open a command prompt and use the following examples:

1. **HTML File Types**
`curl http://127.0.0.1:1278/index.html`

2. **CSS File Types**
`curl http://127.0.0.1:1278/css/styles.css`

# Testing the Functionality of POST

To test the functionality of POST, use the following example:

`curl -X POST -H "Content-Type: text/plain" DEF "%i" http://127.0.0.1:1278/data/vipNames.txt`

# GET command to run n clients in parallel on a single command prompt without blocking it

To run multiple GET requests in parallel on a single command prompt without blocking it, use the following command (for Windows command prompt):

`for /L %i in (1, 1, 20) do start /b curl http://127.0.0.1:1278/data/vipNames.txt`



 