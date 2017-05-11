## Telefonnummer

_Telefonnummer_ is phone number in Swedish. This package formats all Swedish phone numbers, both mobile and landline, to a standard format.

### Usage
```
package main

import (
	"fmt"
	"github.com/believer/telefonnummer-go"
)

func main() {
	fmt.Printf(telefonnummer.Parse("222")) // "Röstbrevlåda"
	fmt.Printf(telefonnummer.Parse("0701234567")) // 070-123 45 67
	fmt.Printf(telefonnummer.Parse("468123456")) // 08-12 34 56
	fmt.Printf(telefonnummer.Parse("031626262")) // 031-62 62 62
	fmt.Printf(telefonnummer.Parse("050012345")) // 0500-123 45
}
```

### Tests
```
go test
```
