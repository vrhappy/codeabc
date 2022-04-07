// testf
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type student struct {
	name string
}

func randn3(n int) []int {
	if n == 1 {
		return []int{0}
	} else if n == 2 {
		return []int{0, 1}
	}
	var a, b, c int
	a, b, c = rand.Intn(n), rand.Intn(n), rand.Intn(n)
	for b == a {
		b = rand.Intn(n)
	}
	for c == a || c == b {
		c = rand.Intn(n)
	}
	return []int{a, b, c}
}

func randn_m(n int, m int) []int {
	res := make([]int, n)
	for i, _ := range res {
		res[i] = i
	}
	if n < m {
		return res
	}
	res1 := make([]int, m)
	for i, _ := range res1 {
		cur_len := len(res)
		num := rand.Intn(cur_len)
		res1[i] = res[num]
		res = append(res[:num], res[num+1:]...)
	}
	return res1
	// res := make([]int, n)
	// for i, _ := range res {
	// 	res[i] = i
	// }
	// if n < m {
	// 	return res
	// }
	// res1 := make([]int, m)
	// res2 := res
	// for i, _ := range res1 {
	// 	cur_len := len(res2)
	// 	num := rand.Intn(cur_len)
	// 	res1[i] = res2[num]
	// 	res2 = append(res2[:num], res2[num+1:]...)
	// }
	// return res1

}

func doTask(num int, basicNum int, outChan *chan int, taskWg *sync.WaitGroup) {
	if num > basicNum {
		time.Sleep(time.Millisecond * 600)
		*outChan <- num
	}
	defer func() {
		taskWg.Done()
	}()
}

func main() {
	list1 := []int{1, 2, 3, 4, 5, 6, 8, 10, 11, 13, 19, 22, 100, 99, 220, 339, 22, 23, 40, 999, 111, 121, 130}
	list2 := []int{1, 12, 3, 41, 5, 6, 8, 101, 11, 13, 19, 22, 100, 99, 220, 339, 22, 23, 40, 999, 1011, 1121, 1301}
	list3 := []int{1, 21, 31, 4, 51, 61, 8, 10, 11, 13, 19, 22, 100, 99, 220, 339, 22, 23, 40, 999, 111, 121, 130, 1011}
	delist := make([][]int, 0)
	delist = append(delist, list1, list2, list3)
	base := 1300
	result_list := make([]int, 0)

	wg := &sync.WaitGroup{}
	startTime := time.Now().UnixMilli()
	fmt.Println("delist length is ", len(delist))
	for i := 0; i < len(delist); i++ {
		// go func() {

		outchan := make(chan int, 100)
		taskwg := &sync.WaitGroup{}
		go func(i int) {
			wg.Add(1)
			if i >= len(delist) {
				return
			}
			for _, v := range delist[i] {

				taskwg.Add(1)
				go doTask(v, base, &outchan, taskwg)
			}
			taskwg.Wait()
			close(outchan)
			// for v := range outchan {
			// 	fmt.Println("--round %v find v: %v", i, v)
			// }
			if v, ok := <-outchan; ok == true {
				fmt.Println(">>outChan=%v for %v", v, i)
				result_list = append(result_list, v)
				for v := range outchan {
					result_list = append(result_list, v)
				}
			} else {
				fmt.Println("<<no result in outchan for %v", i)
			}
			defer wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Printf("costTime:%v ms\n", time.Now().UnixMilli()-startTime)
	fmt.Printf("result list is %v\n", result_list)
	fmt.Println("**all are done**")
	return
	fmt.Println("Hello World!")
	res := randn3(10)
	fmt.Println("res is %v", res)
	res1 := randn_m(100, 5)
	fmt.Println("res1 is %v", res1)
	// m := map[int][]student{
	// 	1: {{name: "1"}, {name: "11"}},
	// }
	// m[1][0].name = "2" // 编译错误： cannot assign to struct field m[1].name in map
	// // m := map[int][]student{1: {"1", "2", "3"}}
	// // m[1][0] = "11"
	// fmt.Println(m)
	// // fmt.Println(
	// // 	unsafe.Sizeof(interface{}(0)),
	// // 	unsafe.Sizeof(struct{}{}),
	// // )
}
