package is_integer

import "fmt"

// IsInt - является ли число целым или с плавающей запятой
func IsInt(bits uint32, bias int) bool {
	exponent := int(bits>>23) - bias - 23
	coefficient := (bits & ((1 << 23) - 1)) | (1 << 23)
	intTest := coefficient & (1<<uint32(-exponent) - 1)

	fmt.Printf("\nExponent: %d Coefficient: %d IntTest: %d\n",
		exponent,
		coefficient,
		intTest)

	if exponent < -23 {
		fmt.Printf("NOT INTEGER\n")
		return false
	}

	if exponent < 0 && intTest != 0 {
		fmt.Printf("NOT INTEGER\n")
		return false
	}

	fmt.Printf("INTEGER\n")
	return true
}
