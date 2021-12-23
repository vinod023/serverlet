package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	str := `Your childhood teacher did not wrong you when they taught you that there should be three, or four, or five sentences in a paragraph. It is important to understand, however, that the aim in teaching this was not to impart a hard-and-fast rule of grammar, drawn from an authoritative-but-dusty book. The true aim of this strategy was to teach you that your ideas must be well supported to be persuasive and effective`

	res := strings.Split(str, " ")

	data := make(map[string]int)

	for _, val := range res {

		mapData, ok := data[val]
		if ok {
			mapData++
			data[val] = mapData
		} else {
			data[val] = 1
		}
	}

	keysData := []string{}
	for key, val := range data {
		keysData = append(keysData, strconv.Itoa(val)+"-"+key)
	}

	sort.Slice(keysData, func(i, j int) bool {
		return keysData[i] > keysData[j]
	})

	for i, val := range keysData {
		if i == 10 {
			break
		}
		fmt.Println(val)
	}
}
