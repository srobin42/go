//Package cryptosquare uses square code to create a secret message.
package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

const testVersion = 2

//Encode creates the secret message string.
func Encode(original string) (secret string) {
	sq := []string{}
	sub := ""

	original = strings.ToLower(original)
	reg := regexp.MustCompile("[^a-z0-9_]")
	//remove non-letter, non-number characters
	secret = reg.ReplaceAllString(original, "$1")
	//encode message
	if secret != "" {
		l := len(secret)
		rows := int(math.Ceil(math.Sqrt(float64(l))))
		//break into even rows of text
		for i, val := range secret {
			sub += string(val)
			if (i+1)%rows == 0 {
				sq = append(sq, sub)
				sub = ""
			} else if (i + 1) == l {
				//pad last chunk if necessary
				sq = append(sq, padRow(sub, rows))
				sub = ""
			}
		}
		//take letters from rows column by column, left to right
		for c := 0; c < rows; c++ {
			for _, val := range sq {
				if val[c] != ' ' {
					sub += string(val[c])
				}
			}
			sub += " "
		}
		secret = strings.TrimRight(sub, " ")
	}
	return secret
}

func padRow(sub string, rows int) string {
	t := make([]byte, rows)
	for i := range t {
		t[i] = ' '
	}
	copy(t[:len(sub)], sub)
	return string(t)
}
