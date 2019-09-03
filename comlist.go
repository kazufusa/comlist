// +build !windows

package comlist

import "fmt"

func Coms() (coms []string, err error) {
	return coms, fmt.Errorf("Operationg system not supported")
}
