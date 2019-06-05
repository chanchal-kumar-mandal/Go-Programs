/**
 * functins
 */

// package main

// import(
// 	"fmt"
// )

// func main() {
// 	var a int = 100
// 	var b int = 200
// 	var ret int

// 	ret = max(a, b)

// 	fmt.Printf("Max value is %d", ret)

// 	m, n := swap("Jim", "Umair")
// 	fmt.Println(m, n)
// }

// func max(num1, num2 int) int {
// 	var result int

// 	if (num1 > num2) {
// 		result = num1
// 	} else {
// 		result = num2
// 	}

// 	return result
// }

// func swap(x, y string) (string, string) {
// 	return y, x
// }


/**
 * Print Hello World
 */ 

// package main

// import (
// 	"fmt"
// )

// func main(){
// 	fmt.Println("Hello WOrld!")
// }


/**
 * Error Handling
 */

// package main

// import (
// 	"fmt"
// 	"math"
// 	"errors"
// )

// // Use errors.New to construct a basic error message
// func Sqrt(value float64) (float64, error){
// 	if value < 0 {
// 		return 0, errors.New("Math: negative number passed to Sqrt")
// 	}
// 	return math.Sqrt(value), nil
// }

// func main(){
// 	result, error := Sqrt(-1)

// 	if error != nil {
// 		fmt.Println(error)
// 	} else {
// 		fmt.Println(result)
// 	}

// 	result1, error := Sqrt(9)

// 	if error != nil {
// 		fmt.Println(error)
// 	} else {
// 		fmt.Println(result1)
// 	}
// }


/**
 * Hello World
 */
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("Hello Word")
// }



/**
 * Interface
 */
// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Shape interface {
// 	area() float64
// }

// type Circle struct {
// 	x, y, radius float64
// }

// type Rectangle struct {
// 	width, height float64
// }

// func(circle Circle) area() float64 {
// 	return math.Pi * circle.radius * circle.radius
// } 

// func(rectangle Rectangle) area() float64 {
// 	return rectangle.width * rectangle.height
// }

// func getArea(shape Shape) float64 {
// 	return shape.area()
// }

// func main() {
// 	circle := Circle{x:0, y:0, radius:5}
// 	rectangle := Rectangle{width:5, height:10}

// 	fmt.Printf("Area of the circle %f\n", getArea(circle))
// 	fmt.Printf("Area of the rectangle %f\n", getArea(rectangle))
// }

