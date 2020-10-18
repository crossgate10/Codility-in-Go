package main

import "log"

func main() {
	// TODO: 是否存在O(n)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// upper bound 上三角
			if j >= i {
				log.Printf("i: %d, j: %d", i, j)
			}

			// lower bound 下三角
			//if j <= i {
			//	log.Printf("i: %d, j: %d", i, j)
			//}
		}
	}
}
