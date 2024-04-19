package main

import "fmt"

func main() {
	req := []uint64{233, 3232, 4343, 212, 233, 2, 23, 343, 3234223, 21, 112, 3}
	res := sSort(req)
	fmt.Println(res)
}

func sSort(req []uint64) []uint64 {
	if len(req) <= 1 {
		return req
	}
	for i := 0; i < len(req); i++ {
		min := i
		for j := i + 1; j < len(req); j++ {
			if req[min] > req[j] {
				min = j
			}
		}
		if min != i {
			req[min], req[i] = req[i], req[min]
		}
	}
	return req
}
