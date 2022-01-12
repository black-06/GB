## Usage

```go
package you

import (
	"github.com/black-06/gb"
)

func ValidateDivisionCode() (bool, error) {
	gbDivisionCode, err := gb.GetGB("GBT2260")
	if err != nil {
		return false, err
	}

	// get a fake division code
	divisionCode := gbDivisionCode.Fake()

	// check the data is a right division code
	return gbDivisionCode.Validate(divisionCode), nil
}
```

About gb name, see [GB Name Map](/README.md#GB-Name-Map)