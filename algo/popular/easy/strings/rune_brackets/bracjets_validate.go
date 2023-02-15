package rune_brackets

/*
Классическая задача - надо проверить, что у каждой открывающийся скобки, есть своя закрывающаяся.
При этом, они  должны быть как бы а одном уровне.
Как если бы это ide проверяла код на открытые/закрытые скобки.

Given string s.
s contains only '[', ']', '(', ')', '{', '}'

Examples:
"[](){}" -> true
"[({})]" -> true
"[(])" -> false
"}{" -> false
"[](){[({})]}[][({})](){}[[(])](){}[]([({})]){}[](){}"
*/

// Validate - used for validate string of brackets.
func Validate(s string) bool {
	checkMap := map[rune]rune{
		'[': ']',
		//']': '[',
		'(': ')',
		//')': '(',
		'{': '}',
		//'}': '{',
	}

	stack := NewStack()

	var lastSymbol rune
	for _, symbol := range s {
		lastSymbol = stack.Front()
		if checkMap[lastSymbol] == symbol {
			stack.Pop()

			continue
		}

		stack.Push(symbol)
	}

	if stack.Size() == 0 {
		return true
	}

	return false
}
