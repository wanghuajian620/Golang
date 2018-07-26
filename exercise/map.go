/*
 *  Revision History:
 *      Initial: 2018/07/25    Wang Huajian
 */

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m) // 结果是 map[k1:7 k2:13]

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2} // 结果是 map[foo:1 bar:2]
	fmt.Println("map:", n)
}