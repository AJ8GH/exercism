package diamond

import "errors"

const (
	first = 'A'
	last  = 'Z'
)

/*

 A -> 0

  A
 B B -> 1
  A

   A
  B B
 C   C -> 3
  B B
	 A

    A
   B B
  C   C
 D     D -> 5
  C   C
   B B
	  A

     A
    B B
   C   C
  D     D
 E       E -> 7
  D     D
   C   C
    B B
	   A

	0 1 2 3 4 5 6
  0 1 3 5 7 9 11

	1 2 3 4 5 6
	0 1 3 5 7 9
*/

func Gen(char byte) (string, error) {
	if char < first || char > last {
		return "", errors.New("bad char")
	}
	if char == first {
		return string(first), nil
	}

	diff := int(char) - int(first)
	padding := diff*2 - 1

	out := ""
	for i := first; i <= rune(char); i++ {
		line := ""
		switch i {
		case first:
			for j := 0; j < diff; j++ {
				line += " "
			}
			line += string(rune(i))
			for j := 0; j < diff; j++ {
				line += " "
			}
		case rune(char):
			line += string(char)
			for j := 0; j < padding; j++ {
				line += " "
			}
			line += string(char)
		default:
			currentDiff := i - first
			currentPadding := currentDiff*2 - 1
			middle := string(i)
			for j := 0; j < int(currentPadding); j++ {
				middle += " "
			}
			middle += string(i)
			sidePadding := padding - int(currentPadding)
			for j := 0; j < sidePadding/2; j++ {
				line += " "
			}
			line += middle
			for j := 0; j < sidePadding/2; j++ {
				line += " "
			}
		}
		out += line + "\n"
	}

	for i := rune(char - 1); i >= first; i-- {
		line := ""
		switch i {
		case first:
			for j := 0; j < diff; j++ {
				line += " "
			}
			line += string(rune(i))
			for j := 0; j < diff; j++ {
				line += " "
			}
		default:
			currentDiff := i - first
			currentPadding := currentDiff*2 - 1
			middle := string(i)
			for j := 0; j < int(currentPadding); j++ {
				middle += " "
			}
			middle += string(i)
			sidePadding := padding - int(currentPadding)
			for j := 0; j < sidePadding/2; j++ {
				line += " "
			}
			line += middle
			for j := 0; j < sidePadding/2; j++ {
				line += " "
			}
		}
		out += line
		if i != first {
			out += "\n"
		}
	}
	return out, nil
}
