func romanToInt(s string) int {
	k := strings.Split(s, "")
	o := 0
	for i, z := range k {
		if z == "I" {
			if i != len(k)-1 && (k[i+1] == "V" || k[i+1] == "X") {
				o -= 1
			} else {
				o += 1
			}
		} else if z == "V" {
			o += 5
		} else if z == "X" {
			if i != len(k)-1 && (k[i+1] == "L" || k[i+1] == "C") {
				o -= 10
			} else {
				o += 10
			}
		} else if z == "L" {
			o += 50
		} else if z == "C" {
			if i != len(k)-1 && (k[i+1] == "D" || k[i+1] == "M") {
				o -= 100
			} else {
				o += 100
			}
		} else if z == "D" {
			o += 500
		} else if z == "M" {
			o += 1000
		}
	}

	return o
}
