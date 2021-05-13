package socket

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	server := fmt.Sprintf("%s:%d", "localhost", 80)
	fmt.Println(server)
}
