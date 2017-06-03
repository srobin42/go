//Package secret takes a number, compares it in binary, returns secret handshake sequence, string array output.
package secret

const testVersion = 1

type binFunc func(uint) string

//Handshake takes number, returns secret handshake phrase, string array output.
func Handshake(i uint) []string {
	var shake []string
	tests := []binFunc{testOne, testTwo, testFour, testEight}

	if i > 16 {
		//reverse order of tests
		for j := 1; j >= 0; j-- {
			opp := 3 - j
			tests[j], tests[opp] = tests[opp], tests[j]
		}
	}
	//run the binary compare tests
	for _, val := range tests {
		if tmp := val(i); tmp != "" {
			shake = append(shake, tmp)
		}
	}
	return shake
}

func testOne(u uint) string {
	s := ""
	if u&1 == 1 {
		s = "wink"
	}
	return s
}

func testTwo(u uint) string {
	s := ""
	if u&2 == 2 {
		s = "double blink"
	}
	return s
}

func testFour(u uint) string {
	s := ""
	if u&4 == 4 {
		s = "close your eyes"
	}
	return s

}

func testEight(u uint) string {
	s := ""
	if u&8 == 8 {
		s = "jump"
	}
	return s
}
