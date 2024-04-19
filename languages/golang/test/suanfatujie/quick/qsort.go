package main

import "fmt"

func main() {
	req := []uint64{283, 233, 3232, 4343, 212, 233, 2, 23, 343, 3234223, 21, 112, 3}
	res := qSort(req)
	fmt.Println(res)
}

func qSort(req []uint64) []uint64 {
	if len(req) <= 1 {
		return req
	} else {
		var min, max []uint64

		for i := 1; i < len(req); i++ {
			if req[i] > req[0] {
				max = append(max, req[i])
			} else {
				min = append(min, req[i])
			}
		}
		return append(append(qSort(min), req[0]), qSort(max)...)
	}
}
