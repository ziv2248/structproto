package converter

import (
	"fmt"
)

func newConvErr(from interface{}, to string) error {
	return fmt.Errorf("cannot convert %#v (type %[1]T) to %v", from, to)
}
