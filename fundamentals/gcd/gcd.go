package gcd

// GCD 求两个非负整数的最大公约数
// Page number - Chinese Edition: 1
// Page number : 4
func GCD(i, j int) int {
	if j == 0 {
		return i
	}
	s := i % j
	return GCD(j, s)
}
