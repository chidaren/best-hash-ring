### BestRing
--- 

#### Target
provide a consistent hash-ring with high performance and high stability :)

#### Example

```go

// initialization
b := bestring.NewBestRing(vNodeNumber)

// add a real server 
b.AddNode(serverAddress)

// find a suitable server
serverHits := b.GetNode(clientRequestKey)

```

#### Todo
1. provide complete features 
2. provide stability report