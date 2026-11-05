package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strings"
)

func buildHackedStrings(n int) []string {
	aBlock, bBlock := thueMorseBlocks()
	const blockCount = 7
	const suffix = "aaaaaaa"

	result := make([]string, 0, n)
	for k := 0; k < n; k++ {
		var b strings.Builder
		for bit := range blockCount {
			if (k>>bit)&1 == 0 {
				b.WriteString(aBlock)
			} else {
				b.WriteString(bBlock)
			}
		}
		b.WriteString(suffix)
		result = append(result, b.String())
	}
	return result
}

func thueMorseBlocks() (string, string) {
	const blockLen = 64
	aBlock := make([]byte, blockLen)
	bBlock := make([]byte, blockLen)
	for i := range blockLen {
		if bits.OnesCount(uint(i))%2 == 0 {
			aBlock[i] = 'a'
			bBlock[i] = 'b'
		} else {
			aBlock[i] = 'b'
			bBlock[i] = 'a'
		}
	}
	return string(aBlock), string(bBlock)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	if _, err := fmt.Fscan(reader, &n); err != nil {
		return
	}

	for _, s := range buildHackedStrings(n) {
		fmt.Fprintln(writer, s)
	}
}
