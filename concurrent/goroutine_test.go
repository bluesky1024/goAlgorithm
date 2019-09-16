package concurrent

import (
	"fmt"
	"io/ioutil"
	"math"
	"net"
	"os"
	"reflect"
	"sync"
	"testing"
	"text/tabwriter"
	"time"
)

//如果没有join点，可能在main goroutine 结束之前都还没执行
func TestGoroutineWithoutJoin(t *testing.T) {
	i := 1
	printNum := func() {
		fmt.Println(i)
	}
	go printNum()

	for {
		if i >= 100000 {
			break
		}
		i++
	}
}

func TestBasicGoroutine(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	i := 1
	printNum := func() {
		defer wg.Done()
		//打印出来的i不固定
		fmt.Println(i)
	}
	go printNum()

	for {
		if i >= 100000 {
			break
		}
		i++
	}
	wg.Wait()
}

func TestUseValInHeapInGoroutine(t *testing.T) {
	wg := sync.WaitGroup{}
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}

func TestPassValToGoroutine(t *testing.T) {
	wg := sync.WaitGroup{}
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(str string) {
			defer wg.Done()
			fmt.Println(str)
		}(salutation)
	}
	wg.Wait()
}

func TestChangeValPassedToGoroutine(t *testing.T) {
	wg := sync.WaitGroup{}
	ori := []int{1, 2, 3}
	pOri := []*int{&ori[0], &ori[1], &ori[2]}
	for _, pNode := range pOri {
		wg.Add(1)
		go func(node *int) {
			defer wg.Done()
			fmt.Println(*node)
			*node = *node * (*node)
		}(pNode)
	}
	wg.Wait()
	fmt.Println("after change")
	fmt.Println(ori)
}

//go test -bench BenchmarkBasicGoroutine -run =^$ -cpu 1,2,4,8
func BenchmarkBasicGoroutine(b *testing.B) {
	wg := sync.WaitGroup{}
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}

func TestMutex(t *testing.T) {
	var count = 0
	var lock sync.Mutex
	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Println("incrementing: ", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Println("decrementing: ", count)
	}

	var arithmetic sync.WaitGroup
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("complete.")
}

func TestRWMutex(t *testing.T) {
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1 * time.Millisecond)
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)

		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}

		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutext\n")
	for i := 0; i < 5; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}
}

//sync.NewCond 能让for轮训更有效的等待
//重点在于 c.Signal() 能通知 goroutine 阻塞的调用 c.Wait() ，提示条件已经被触发
//怎么感觉这种方式就是变相的channel通知
func TestSyncCond(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)
	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("removed from queue")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("adding to queue", len(queue))
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}

//increment 和 decrement 两个函数只调用了 increment
//说明针对同一个once，once.Do 只调用了一次，不管Do的是什么函数
func TestOnce(t *testing.T) {
	var count int
	increment := func() {
		count++
	}
	decrement := func() {
		count--
	}
	var once sync.Once
	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
			once.Do(decrement)
		}()
	}

	increments.Wait()
	fmt.Println("counts:", count)
}

//"create new instance"只执行了一次
//说明pool中的实例得到了复用
func TestBasicPool(t *testing.T) {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create new instance")
			return struct{}{}
		},
	}
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()
}

func connectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			fmt.Println("cannot listen:", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				fmt.Println("cannot appect connection:", err)
				continue
			}
			connectToService()
			fmt.Fprintln(conn, "")
			conn.Close()
		}
	}()
	return &wg
}

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}
	for i := 0; i < 100; i++ {
		p.Put(p.New())
	}
	return p
}

func startNewNetworkDaemonWithPool() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connPool := warmServiceConnCache()
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			fmt.Println("cannot listen:", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				fmt.Println("cannot appect connection:", err)
				continue
			}
			svcConn := connPool.Get()
			fmt.Fprintln(conn, "")
			connPool.Put(svcConn)
			conn.Close()
		}
	}()
	return &wg
}

func init() {
	//按照书本，理论上应该withPool会比普通的快1000倍，但实际上并没有。。。
	//daemonStarted := startNetworkDaemon()
	//daemonStarted := startNewNetworkDaemonWithPool()
	//daemonStarted.Wait()
}

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}

func TestSingleWayChan(t *testing.T) {
	//单向管道是无法进行逆向操作的，以下操作会报错
	//writeStream := make(chan<- interface{})
	//readStream := make(<-chan interface{})
	//
	//<-writeStream
	//readStream <- struct{}{}
}

func TestGetDataFromChanColosed(t *testing.T) {
	intStream := make(chan int)
	close(intStream)
	integer, ok := <-intStream
	fmt.Println(integer, ok)
}

func TestRangeChanClosed(t *testing.T) {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 0; i < 5; i++ {
			intStream <- i * 10
		}
	}()
	<-intStream
	//close(chan) 可以让以下range退出循环
	for integer := range intStream {
		fmt.Println(integer)
	}
}

//close chan 可以用来进行多个协程的同时通知
func TestCloseChanForInform(t *testing.T) {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Println(i, "has begun")
		}(i)
	}
	fmt.Println("unblocking goroutines...")
	fmt.Println("before close begin")
	close(begin)
	wg.Wait()
	fmt.Println("after wg wait")
}

/*
这种方式很重要
chanOwner函数仅对外提供了一个只读管道，
函数内部维护了管道的输入和关闭，
因此外部调用者只需要知道它应该如何处理阻塞读取和chan的关闭，
使得系统条理更加清晰

总之---尽量保持channel所有权的范围足够小
*/
func TestFuncPackChanInputAndClose(t *testing.T) {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i < 5; i++ {
				resultStream <- i * 10
			}
		}()
		return resultStream
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Println("received: ", result)
	}
	fmt.Println("done")
}

func TestGetChanFromTwoGoroutine(t *testing.T) {
	ch1 := make(chan int)
	go func(ch chan int) {
		select {
		case a := <-ch:
			fmt.Println("goroutine 1 get ch", a, reflect.TypeOf(a))
		}
	}(ch1)

	go func(ch chan int) {
		select {
		case a := <-ch:
			fmt.Println("goroutine 2 get ch", a, reflect.TypeOf(a))
		}
	}(ch1)

	go func(ch chan int) {
		select {
		case a := <-ch:
			fmt.Println("goroutine 3 get ch", a, reflect.TypeOf(a))
		}
	}(ch1)

	go func(ch chan int) {
		select {
		case a := <-ch:
			fmt.Println("goroutine 4 get ch", a, reflect.TypeOf(a))
		}
	}(ch1)

	ch1 <- 10
	close(ch1)
	time.Sleep(1 * time.Second)
}

func TestBasicSelect(t *testing.T) {
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	go func() {
		c1 <- 1
	}()
	go func() {
		c2 <- struct{}{}
	}()
	time.Sleep(1 * time.Microsecond)
	select {
	case <-c1:
		fmt.Println("c1 turn do something")
	case <-c2:
		fmt.Println("c2 turn do something")
	default:
		fmt.Println("get no chan")
	}
}

func TestBasicSelectWithoutDefault(t *testing.T) {
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- 1
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- struct{}{}
	}()
	select {
	case <-c1:
		fmt.Println("c1 turn do something")
	case <-c2:
		fmt.Println("c2 turn do something")
	}
	fmt.Println("end")
}

func TestBasicSelectTimes(t *testing.T) {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)
	cnt1 := 0
	cnt2 := 0
	for i := 0; i < 1000; i++ {
		select {
		case <-c1:
			cnt1++
		case <-c2:
			cnt2++
		}
	}
	fmt.Println("cnt1:", cnt1, "cnt2:", cnt2)
}

func TestBasicSelectWithTimeout(t *testing.T) {
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- struct{}{}
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- struct{}{}
	}()
	select {
	case <-c1:
		fmt.Println("c1 turn do something")
	case <-c2:
		fmt.Println("c2 turn do something")
	case <-time.After(1 * time.Second):
		fmt.Println("time out")
	}
	fmt.Println("end")
}
