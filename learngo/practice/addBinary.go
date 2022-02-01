package practice

func AddBinary(a string, b string) string {
	res, s := "", ""
	c := 0
	i, j := len(a)-1, len(b)-1
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		sum := (a[i] - 48) + (b[j] - 48) + uint8(c)
		c, s = getCarryAndSum(sum)
		res = s + res
	}
	if i >= 0 {
		for ; i >= 0; i-- {
			sum := (a[i] - 48) + uint8(c)
			c, s = getCarryAndSum(sum)
			res = s + res
		}
	}
	if j >= 0 {
		for ; j >= 0; j-- {
			sum := (b[j] - 48) + uint8(c)
			c, s = getCarryAndSum(sum)
			res = s + res
		}
	}
	if c == 1 {
		res = "1" + res
	}
	return res
}

func getCarryAndSum(sum uint8) (int, string) {
	switch sum {
	case 0:
		return 0, "0"
	case 1:
		return 0, "1"
	case 2:
		return 1, "0"
	case 3:
		return 1, "1"
	default:
		return 0, "0"
	}
}
