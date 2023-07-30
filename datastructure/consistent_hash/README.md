# Consistent Hash

```Go
package main

import (
	"fmt"

	chash "github.com/abcdlsj/x/datastructure/consistenthash"
)

func main() {
	replicas, hashFn := 3, chash.HashFn(nil)
	cHash := chash.New(replicas, hashFn)
	cHash.Add("node1", "node2", "node3")
	for i := 0; i < 3; i++ {
		fmt.Printf("[%d] %s -> %s\n", i, fmt.Sprintf("key%d", i), cHash.Get(fmt.Sprintf("key%d", i)))
	}
}

```