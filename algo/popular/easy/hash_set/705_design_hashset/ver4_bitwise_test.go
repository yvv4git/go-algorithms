package _05_design_hashset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVer4(t *testing.T) {
	obj := ConstructorV4()

	obj.Add(1)
	obj.Add(2)

	assert.True(t, obj.Contains(1))
	assert.False(t, obj.Contains(3))

	obj.Add(2)
	assert.True(t, obj.Contains(2))

	obj.Remove(2)
	assert.False(t, obj.Contains(2))

	printBinary(obj.set)
	obj.Remove(2)
	printBinary(obj.set)

	obj.Add(0)
	printBinary(obj.set)
	obj.Add(0)
	printBinary(obj.set)
}

func TestVer4_2(t *testing.T) {
	obj := ConstructorV4()
	// ["MyHashSet","contains","remove","add","add","contains","remove","contains","contains","add","add","add","add","remove","add","add","add","add","add","add","add","add","add","add","contains","add","contains","add","add","contains","add","add","remove","add","add","add","add","add","contains","add","add","add","remove","contains","add","contains","add","add","add","add","add","contains","remove","remove","add","remove","contains","add","remove","add","add","add","add","contains","contains","add","remove","remove","remove","remove","add","add","contains","add","add","remove","add","add","add","add","add","add","add","add","remove","add","remove","remove","add","remove","add","remove","add","add","add","remove","remove","remove","add","contains","add"]
	// [[],[72],[91],[48],[41],[96],[87],[48],[49],[84],[82],[24],[7],[56],[87],[81],[55],[19],[40],[68],[23],[80],[53],[76],[93],[95],[95],[67],[31],[80],[62],[73],[97],[33],[28],[62],[81],[57],[40],[11],[89],[28],[97],[86],[20],[5],[77],[52],[57],[88],[20],[48],[42],[86],[49],[62],[53],[43],[98],[32],[15],[42],[50],[19],[32],[67],[84],[60],[8],[85],[43],[59],[65],[40],[81],[55],[56],[54],[59],[78],[53],[0],[24],[7],[53],[33],[69],[86],[7],[1],[16],[58],[61],[34],[53],[84],[21],[58],[25],[45],[3]]
	// Actual: [null,false,null,null,null,false,null,true,false,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,false,null,false,null,null,false,null,null,null,null,null,null,null,null,true,null,null,null,null,false,null,false,null,null,null,null,null,true,null,null,null,null,true,null,null,null,null,null,null,true,true,null,null,null,null,null,null,null,false,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,false,null]
	// Expect: [null,false,null,null,null,false,null,true,false,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,false,null,true,null,null,true,null,null,null,null,null,null,null,null,true,null,null,null,null,false,null,false,null,null,null,null,null,true,null,null,null,null,true,null,null,null,null,null,null,true,true,null,null,null,null,null,null,null,false,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,false,null]
	//obj := ConstructorV4()
	//t.Logf("->%v", obj.Contains(72)) // false
	//obj.Remove(91)
	//obj.Add(48)
	//obj.Add(41)
	//t.Log(obj.Contains(96))
	//obj.
	actionsMap := map[string]func(int) string{
		"contains": func(v int) string {
			result := obj.Contains(v)
			switch result {
			case true:
				return "true"
			case false:
				return "false"
			default:
				return "null"
			}
		},
		"add": func(v int) string {
			obj.Add(v)
			return "null"
		},
		"remove": func(v int) string {
			obj.Remove(v)
			return "null"
		},
	}

	actionsList := []string{"contains", "remove", "add", "add", "contains", "remove", "contains", "contains", "add", "add", "add", "add", "remove", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "contains", "add", "contains", "add", "add", "contains", "add", "add", "remove", "add", "add", "add", "add", "add", "contains", "add", "add", "add", "remove", "contains", "add", "contains", "add", "add", "add", "add", "add", "contains", "remove", "remove", "add", "remove", "contains", "add", "remove", "add", "add", "add", "add", "contains", "contains", "add", "remove", "remove", "remove", "remove", "add", "add", "contains", "add", "add", "remove", "add", "add", "add", "add", "add", "add", "add", "add", "remove", "add", "remove", "remove", "add", "remove", "add", "remove", "add", "add", "add", "remove", "remove", "remove", "add", "contains", "add"}
	values := []int{72, 91, 48, 41, 96, 87, 48, 49, 84, 82, 24, 7, 56, 87, 81, 55, 19, 40, 68, 23, 80, 53, 76, 93, 95, 95, 67, 31, 80, 62, 73, 97, 33, 28, 62, 81, 57, 40, 11, 89, 28, 97, 86, 20, 5, 77, 52, 57, 88, 20, 48, 42, 86, 49, 62, 53, 43, 98, 32, 15, 42, 50, 19, 32, 67, 84, 60, 8, 85, 43, 59, 65, 40, 81, 55, 56, 54, 59, 78, 53, 0, 24, 7, 53, 33, 69, 86, 7, 1, 16, 58, 61, 34, 53, 84, 21, 58, 25, 45, 3}
	expect := []string{"false", "null", "null", "null", "false", "null", "true", "false", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "false", "null", "true", "null", "null", "true", "null", "null", "null", "null", "null", "null", "null", "null", "true", "null", "null", "null", "null", "false", "null", "false", "null", "null", "null", "null", "null", "true", "null", "null", "null", "null", "true", "null", "null", "null", "null", "null", "null", "true", "true", "null", "null", "null", "null", "null", "null", "null", "false", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "false", "null"}
	for i := 0; i < len(values); i++ {
		act := actionsList[i]
		fn := actionsMap[act]
		result := fn(values[i])
		fmt.Printf("[%d] obj.%s(%d) = %s [%s] \n", i, actionsList[i], values[i], result, expect[i])
		require.Equal(t, expect[i], result)
	}

}

func TestVer4_3(t *testing.T) {
	obj := ConstructorV4()
	printBinary(obj.set)

	obj.Add(95)
	printBinary(obj.set)
	require.True(t, obj.Contains(95))
}

func printBinary(n uint) {
	for i := 31; i >= 0; i-- {
		if n&(1<<i) != 0 {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
	fmt.Println()
}
