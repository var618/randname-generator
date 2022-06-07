package generator

import (
	"fmt"
	"testing"
)

func TestGenerateRoomWithoutSeparator(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Println(GenerateRoomWithoutSeparator())
	}
}
