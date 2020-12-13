# goaround

Workarounds for Go.

## Usage

Right now goaround only includes one function `NotNil`.
`NotNil` can be used to check if the function argument passed in is nil.
It will panic when encountering a nil value.

```go
import (
	"github.com/weakish/goaround"
)

func foo(bar []string) int {
	goaround.NotNil(bar)
	return len(bar)
}
```

## License

0BSD