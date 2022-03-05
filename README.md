### ErrorCallBack

```go
package main

import (
	"github.com/AndreySHSH/ecback"
	"os"
)

func main() {
	ecb := ecback.InitErrCallBack(ecback.ECBack{
		CallBackUrl: "https://example.com/event",
		ShowLog:     true,
		ApplocationTitle: "test-v1.1.1",
	})

	_, err := os.Open("asd")
	ecb.E(err, func(event *ecback.ECBack) *ecback.ECBack {
		event.StopExecution = true
		return event
	})
}

```