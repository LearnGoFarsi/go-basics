package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// a map is a reference to a hashtable that maps a key to a value map[key] = value
	// The key type must be comparable using ==
	rankingMap := make(map[string]int)
	rankingMap["Go"] = 1
	rankingMap["Python"] = 2
	rankingMap["JavaScript"] = 3
	rankingMap["C#"] = 4
	rankingMap["Java"] = 1 << 10 // :D

	fmt.Println(rankingMap)

	delete(rankingMap, "Java")
	delete(rankingMap, "dummy")
	fmt.Println(rankingMap)

	// order is random. To enumerate the key/value in order, we must sort the key explicitly
	keys := sortedKeys(&rankingMap)
	for _, k := range keys {
		println(k, rankingMap[k])
	}

	// poor way to check if a key exists
	//rankingMap["C++"] = 0
	if v := rankingMap["C++"]; v == 0 { // v==nil does not work because type int cannot be nil, rather 0
		fmt.Println("C++ is not found")
	}

	// safe way to check if a key exists
	//rankingMap["C++"] = 0
	if _, ok := rankingMap["C++"]; !ok {
		fmt.Println("C++ is not found")
	} else {
		fmt.Println("C++ found")
	}

	// What if a key needs to be slice? Slice is not comparable
	smap := make(map[string]string)
	smap[strings.Join([]string{"a"}, ",")] = "Value of a"
	smap[strings.Join([]string{"a", "b"}, ",")] = "Value of a,b"
	smap[strings.Join([]string{"x", "y"}, ",")] = "Value of x,y"
	fmt.Println(smap[strings.Join([]string{"x", "y"}, ",")])
}

// m is pass by reference and not pass by value
func sortedKeys(m *map[string]int) []string {
	keys := []string{}
	for k, _ := range *m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}
