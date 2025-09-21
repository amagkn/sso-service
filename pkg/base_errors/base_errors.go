package base_errors

import (
	"fmt"
)

func WithPath(path string, err error) error {
	return fmt.Errorf("%s: %w", path, err)
}
