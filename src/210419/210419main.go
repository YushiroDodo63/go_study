package main

import (
  "fmt"
)

func main() {
//   var ns1 [5]int
//   var ns2 = [5]int{10, 20, 30, 40, 50}
//   ns3 := [...]int{10, 20, 30, 40, 50}
//   ns4 := [...]int{5: 50, 10:100}
//   fmt.Printf("%q\n",ns1)
//   fmt.Println(ns2)
//   fmt.Println(ns3)
//   fmt.Println(ns4)
  
// // map 3-34
//   var m1 map[string]int
//   fmt.Println(m1)
//   var m2 = make(map[string]int)
//   fmt.Println(m2)
//   var m3 = make(map[string]int, 10)
//   fmt.Println(m3)

//   // map3-35
//   m4 := map[string]int{"x": 10, "y": 20}
//   fmt.Println(m4["x"])
//   m4["z"] = 30
//   fmt.Println(m4)

// try 3-40
//   type Game struct {
// 	  Name string
// 	  GameID int
// 	  Point int
//   }
//   user1 := Game{Name: "dodo", GameID: 1, Point: 10}
//   fmt.Println(user1)


// 3-47
//   fmt.Println(swap(1))

// 3-49
//   fmt.Println(swap(1,2))
//   var sum int
//   sum = 5 + 6 + 3

// try3-58
	// for i := 1; i <= 100; i++{
	// 	fmt.Print(i)
	// 	if i%2 == 0{
	// 		fmt.Println("-偶数")
	// 	} else {
	// 		fmt.Println("-奇数")
	// 	}
	// }

// try3-59
	n, m := 10, 20
	swap2(&n, &m)
	fmt.Println(n, m)
}

// 3-47
// func swap(x int) string{
// 	var s string = string(x)
// 	return s
// }

// 3-49
// func swap(x, y int) (x2, y2 int){
// 	y2, x2 = x, y
// 	return
// }

// try 3-58
// func swap(x, y int) (x2, y2 int){
// 	y2, x2 = x, y
// 	return
// }

// try 3-59
func swap2(x, y *int) (x2, y2 *int){
	y2, x2 = x, y
	return
}
