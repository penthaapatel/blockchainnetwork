# blockchainnetwork

A basic blockchain implementation written in Go.

### Basic Prototype

A simplified version of a block

```go
type Block struct {
	Index     int    `json:"index"`
	Data      string `json:"data"`
	Timestamp string `json:"timestamp"`
	PrevHash  string `json:"prevhash"`
	Hash      string `json:"hash"`
}
```

How to run :

**Terminal 1**
Start the tcp server.

```bash
$ go run main.go
Starting tcp server on localhost:8080
```

**Terminal 2**
Start a tcp connection on another terminal (Netcat-nc used here). Enter data to generate a new block.

```
$ nc 127.0.0.1 8080
Enter input data for new Block generation: hello blockchain
```

See output on Terminal 1

**Terminal 1**

```
127.0.0.1:50306 joined.
NEW BLOCK CREATED
{
   "index": 1,
   "data": "hello blockchain",
   "timestamp": "2020-01-21 12:55:16.303992792 +0530 IST m=+139.975010023",
   "prevhash": "",
   "hash": "e3565e5344cad9e777a8cb03974c27785f48b14ab292db7e95561c2cda095b4f"
}
```

### Todo

- [ ] Add Proof of Work
