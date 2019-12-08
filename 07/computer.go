package main

//Computer computes
func Computer(n []int, ins []int, i int) (int, []int, int) {

	result := 0
	inputs := 0
	for i < len(n) {
		inst := getOpCode(n[i])

		switch inst.code {
		case 1:
			a := getModedValue(inst.m1, i+1, n)
			b := getModedValue(inst.m2, i+2, n)
			c := getModedIndex(inst.m3, i+3, n)
			n[c] = a + b
			i += 4

		case 2:
			a := getModedValue(inst.m1, i+1, n)
			b := getModedValue(inst.m2, i+2, n)
			c := getModedIndex(inst.m3, i+3, n)
			n[c] = a * b
			i += 4
		case 3:
			c := getModedIndex(inst.m1, i+1, n)
			n[c] = ins[inputs]
			inputs++
			i += 2
		case 4:
			c := getModedValue(inst.m1, i+1, n)
			i += 2
			return c, n, i
		case 5:
			a := getModedValue(inst.m1, i+1, n)
			b := getModedValue(inst.m2, i+2, n)

			if a != 0 {
				i = b
			} else {
				i += 3
			}
		case 6:
			a := getModedValue(inst.m1, i+1, n)
			b := getModedValue(inst.m2, i+2, n)

			if a == 0 {
				i = b
			} else {
				i += 3
			}
		case 7:
			a := getModedValue(inst.m1, i+1, n)
			b := getModedValue(inst.m2, i+2, n)
			c := getModedIndex(inst.m3, i+3, n)

			if a < b {
				n[c] = 1
			} else {
				n[c] = 0
			}

			i += 4
		case 8:
			a := getModedValue(inst.m1, i+1, n)
			b := getModedValue(inst.m2, i+2, n)
			c := getModedIndex(inst.m3, i+3, n)
			if a == b {
				n[c] = 1
			} else {
				n[c] = 0
			}

			i += 4
		case 99:
			return result, n, 99
		}
	}
	return result, n, 99
}
