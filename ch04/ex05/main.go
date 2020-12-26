// Write an in-place function to eliminate adjacent duplicates in a `[]string` slice
// In-place means that you should update the original string rather than creating a new one.

package main

import "fmt"


func main() {
	strSlice := []string{"ab", "ab", "fhsdh", "asy", "tr", "q34yu", "ZXFch", "bb", "bb", "bb", "asyher", "ngfd"}
	eliAdjDup(&strSlice)
	fmt.Println(strSlice)
}

func eliAdjDup(ss *[]string) {
	if len(*ss) == 0 || len(*ss) == 1 {
		return
	}

	for i := 0; i < len(*ss) - 1; {
		if (*ss)[i] == (*ss)[i + 1] {
			if i + 2 < len(*ss) {
				*ss = append((*ss)[:i + 1], (*ss)[i + 2:]...)
			} else {
				*ss = (*ss)[:i + 1]
			}
			
		} else {
			i++
		}
	}
}