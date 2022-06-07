package generator

import (
	"fmt"
	"testing"
)

func TestGenerateRoomWithoutSeparator(t *testing.T) {
	str := GenerateRoomWithoutSeparator()
	fmt.Println(str)
}
