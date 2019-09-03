# comlist

Gets the list of Windows COM ports.

```
package main

import (
	"fmt"
	"log"

	"github.com/kazufusa/comlist"
)

func main() {
	coms, err := comlist.Coms()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coms)
}
```

```
> go run ./cmd
[COM4]
```
