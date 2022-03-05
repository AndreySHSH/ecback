### ErrorCallBack

```go
package main

import (
	"github.com/AndreySHSH/ecback/ecback"
	"os"
)

func main() {
	ecb := ecback.InitErrCallBack(ecback.ECBack{
		CallBackUrl: "https://example.com/event",
		Debug:       true,
	})

	_, err := os.Open("asd")
	ecb.E(err, func(event *ecback.ECBack) *ecback.ECBack {
		event.StopExecution = true
		return event
	})
}

```