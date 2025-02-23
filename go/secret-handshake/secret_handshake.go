package secret

import (
	"slices"
	"strconv"
)

// 00001 = wink
// 00010 = double blink
// 00100 = close your eyes
// 01000 = jump
// 10000 = Reverse the order of the operations in the secret handshake.

var allActions = []string{"wink", "double blink", "close your eyes", "jump", "reverse"}

func Handshake(code uint) (actions []string) {
	bin := strconv.FormatUint(uint64(code), 2)
	for i := len(bin) - 1; i >= 0; i-- {
		if bin[i] == '0' {
			continue
		}

		j := len(bin) - 1 - i
		switch j {
		case 0, 1, 2, 3:
			actions = append(actions, allActions[j])
		case 4:
			slices.Reverse(actions)
		}
	}
	return actions
}
