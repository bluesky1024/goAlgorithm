package concurrent

import "fmt"

/*问题*/
/*
有四个线程1、2、3、4。线程1的功能就是输出1，线程2的功能就是输出2，以此类推.........现在有四个文件ABCD。初始都为空。现要让四个文件呈如下格式：

A：1 2 3 4 1 2....

B：2 3 4 1 2 3....

C：3 4 1 2 3 4....

D：4 1 2 3 4 1....
*/
/*思路*/
/*
重点在于四个线程的顺序控制，对于同一个文件，2的打印必须在1之后，3的打印必须在2之后，以此类推
通过管道控制顺序，1的线程开始来自于4的管道值，。。。
*/
//实现1
func printA() string {
	return "A"
}
func printB() string {
	return "B"
}
func printC() string {
	return "C"
}
func printD() string {
	return "D"
}

func PrintMultiInOrder() {
	chA := make(chan int, 1)
	chB := make(chan int, 1)
	chC := make(chan int, 1)
	chD := make(chan int, 1)

	chD <- 1

	go func(chI chan int, chO chan int) {
		for {
			<-chI
			fmt.Println(printA())
			chO <- 1
		}
	}(chD, chA)

	go func(chI chan int, chO chan int) {
		for {
			<-chI
			fmt.Println(printB())
			chO <- 1
		}
	}(chA, chB)

	go func(chI chan int, chO chan int) {
		for {
			<-chI
			fmt.Println(printC())
			chO <- 1
		}
	}(chB, chC)

	go func(chI chan int, chO chan int) {
		for {
			<-chI
			fmt.Println(printD())
			chO <- 1
		}
	}(chC, chD)
}

//实现2
type printSingleChar struct {
}
