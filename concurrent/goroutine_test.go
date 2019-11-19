package concurrent

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net"
	"os"
	"reflect"
	"strconv"
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
		if i >= 10000000 {
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

// 本来想试试是否可以在select case中监控变长的数组管道，似乎不行
func TestCloseMultiGoroutines(t *testing.T) {
	//程序退出控制
	var wait sync.WaitGroup

	//生产者生存控制管道
	stopChan := make(chan bool)

	//生产者
	nodeNum := 10
	for i := 0; i < nodeNum; i++ {
		wait.Add(1)
		go func(ind int, stopChan <-chan bool) {
			defer func() {
				fmt.Println("stop goroutine", ind)
				wait.Done()
			}()
			stopTime := time.After(1 * time.Second)
			for {
				select {
				case <-stopTime:
					fmt.Println(ind, "time out")
					return
				default:
				}
				select {
				case <-stopChan:
					return
				default:
				}
				time.Sleep(100 * time.Millisecond)
				fmt.Println(ind, "is working")
			}
		}(i, stopChan)
	}

	fmt.Println("wait 2s")
	time.Sleep(2 * time.Second)

	fmt.Println("close chan")
	close(stopChan)

	wait.Wait()
	fmt.Println("close finish")
}

func TestSelectVariableLengthChan(t *testing.T) {
	//节点数据准备
	type chanNode struct {
		sign chan bool
	}

	nodeNum := 3
	nodes := make([]chanNode, nodeNum)

	for i := 0; i < nodeNum; i++ {
		nodes[i] = chanNode{
			sign: make(chan bool),
		}
	}

	//程序退出控制
	var wait sync.WaitGroup

	//生产者生存控制管道
	stopChan := make(chan bool)

	//生产者
	for i := 0; i < nodeNum; i++ {
		wait.Add(1)
		go func(ind int, stopChan <-chan bool) {
			defer func() {
				wait.Done()
				fmt.Println(ind, "stop produce")
			}()
			for {
				select {
				case <-stopChan:
					return
				case nodes[ind].sign <- true:
					fmt.Println("node", ind, "produce")
				}
			}
		}(i, stopChan)
	}

	//消费者协程
	var consumerWait sync.WaitGroup
	for i := 0; i < nodeNum*5; i++ {
		consumerWait.Add(1)
		go func(ind int) {
			defer func() {
				consumerWait.Done()
				fmt.Println(ind, "stop consumer")
			}()
			stopTime := time.After(1 * time.Second)
			for {
				select {
				case <-stopTime:
					return
				default:
				}
				select {
				case <-nodes[0].sign:
					fmt.Println("node 0 consume")
				case <-nodes[1].sign:
					fmt.Println("node 1 consume")
				case <-nodes[2].sign:
					fmt.Println("node 2 consume")
				default:
				}
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	//等待所有消费者停止消费
	consumerWait.Wait()
	fmt.Println("consumer stop")

	//通知所有生产者协程关闭
	close(stopChan)

	//等待所有生产者协程关闭
	wait.Wait()
	fmt.Println("produce stop")

	fmt.Println("task finish")
}

//如何优雅的关闭多生产者多消费者参与的管道，来自网上博客
func TestCloseChannelPerfectly(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const MaxRandomNumber = 1000
	const NumReceivers = 100
	const NumSenders = 10

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the moderator goroutine shown below.
	// Its reveivers are all senders and receivers of dataCh.
	toStop := make(chan string, 1)
	// The channel toStop is used to notify the moderator
	// to close the additional signal channel (stopCh).
	// Its senders are any senders and receivers of dataCh.
	// Its reveiver is the moderator goroutine shown below.
	// It must be a buffered channel.

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(MaxRandomNumber)
				if value == 0 {
					// Here, the try-send operation is to notify the
					// moderator to close the additional signal channel.
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				// The try-receive operation here is to try to exit the
				// sender goroutine as early as possible. Try-receive
				// try-send select blocks are specially optimized by the
				// standard Go compiler, so they are very efficient.
				select {
				case <-stopCh:
					return
				default:
				}

				// Even if stopCh is closed, the first branch in this
				// select block may be still not selected for some
				// loops (and for ever in theory) if the send to dataCh
				// is also non-blocking. If this is not acceptable,
				// then the above try-receive operation is essential.
				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()

			for {
				// Same as the sender goroutine, the try-receive
				// operation here is to try to exit the receiver
				// goroutine as early as possible.
				select {
				case <-stopCh:
					return
				default:
				}

				// Even if stopCh is closed, the first branch in this
				// select block may be still not selected for some
				// loops (and for ever in theory) if the receive from
				// dataCh is also non-blocking. If this is not acceptable,
				// then the above try-receive operation is essential.
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == MaxRandomNumber-1 {
						// The same trick is used to notify
						// the moderator to close the
						// additional signal channel.
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}

					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	// ...
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}

func TestWriteChanInMultiGoroutine(t *testing.T) {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case a := <-ch:
				fmt.Println("get chan num", a)
				//a = a + 1
				break
			}
		}()
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(ii int, ch1 chan<- int) {
			defer wg.Done()
			ch1 <- ii
		}(i, ch)
	}

	wg.Wait()
	fmt.Println("finish")
}

func TestCancelCtxInParentCtx(t *testing.T) {
	pCtx, pCancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		cCtx, _ := context.WithTimeout(ctx, 3*time.Second)
		select {
		case <-cCtx.Done():
			fmt.Println("finish")
			break
		}
	}(pCtx)
	time.Sleep(1 * time.Second)
	pCancel()

	time.Sleep(5 * time.Second)
	fmt.Println("end")
}

func genChanWithCancel(ctx context.Context) <-chan int {
	ch := make(chan int)
	n := 0
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context done")
				close(ch)
				return
			default:
				n++
				fmt.Println("in", n)
				ch <- n
				time.Sleep(time.Second)
			}
		}
	}()

	return ch
}

func TestCloseChanWithContent(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	ch := genChanWithCancel(ctx)

	for i := range ch {
		fmt.Println("out", i)
		time.Sleep(time.Second)
	}
}

func TestChannel(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	num := testChanInOut(ctx)

	time.Sleep(time.Second * 5)

	for i := range num {
		fmt.Println("out ", i)
		time.Sleep(time.Second)
	}
	defer cancel()
}

func testChanInOut(ctx context.Context) <-chan int {
	//缓冲通道和阻塞通道
	c := make(chan int)
	//c := make(chan int, 10)
	n := 0
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context done")
				close(c)
				return
			default:
				n++
				fmt.Println("in ", n)
				c <- n
				fmt.Println("aaa")
				time.Sleep(time.Second)
			}
		}
	}()

	return c
}

func TestOrChannel(t *testing.T) {
	var or func(chs ...<-chan interface{}) <-chan interface{}

	or = func(chs ...<-chan interface{}) <-chan interface{} {
		switch len(chs) {
		case 0:
			return nil
		case 1:
			return chs[0]
		}

		orDone := make(chan interface{})
		go func() {
			defer close(orDone)

			switch len(chs) {
			case 2:
				select {
				case <-chs[0]:
				case <-chs[1]:
				}
			default:
				select {
				case <-chs[0]:
				case <-chs[1]:
				case <-chs[2]:
				case <-or(append(chs[3:], orDone)...):
				}
			}
		}()
		return orDone
	}

	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})

	ch := or(ch1, ch2, ch3)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 1
	}()
	go func() {
		time.Sleep(3 * time.Second)
		ch3 <- 1
	}()

	select {
	case <-ch:
		fmt.Println("get message")
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer func() {
				close(c)
				fmt.Println("call defer", after)
			}()
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
	)
	fmt.Println("done after", time.Since(start))
}

func TestChanWithBufferInAGoroutine(t *testing.T) {
	ch := make(chan int)
	ch <- 1
	select {
	case <-ch:
		fmt.Println("get ch")
	}
	close(ch)
	fmt.Println("end")
}

func AppendWithCopy(i int, x int, a []int) []int {
	if i <= len(a) {
		fmt.Println("aaa")
		a = append(a, 0)
		copy(a[i+1:], a[i:])
		a[i] = x
		return a
	} else {
		a = append(a, x)
		return a
	}
}

func TestSlice(t *testing.T) {

	num1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(num1, len(num1), cap(num1))
	num1 = AppendWithCopy(1, 99, num1)
	fmt.Println(num1, len(num1), cap(num1))

	return

	var numbers []int
	fmt.Println(numbers, len(numbers), cap(numbers))

	numbers = append(numbers, 0)
	fmt.Println(numbers, len(numbers), cap(numbers))

	numbers = append(numbers, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(numbers, len(numbers), cap(numbers))

	s1 := []int{11, 12, 13, 14, 15, 16, 17}
	numbers = append(numbers, s1...)
	fmt.Println(numbers, len(numbers), cap(numbers))

	numbers = numbers[1:]
	fmt.Println(numbers, len(numbers), cap(numbers))

	numbers = append(numbers[:3], numbers[4:]...)
	fmt.Println(numbers, len(numbers), cap(numbers))
}

func TestDefer(t *testing.T) {
	f1 := func() (a int) {
		a = 3
		defer func() {
			a = a + 1
		}()
		return a
	}

	f2 := func() int {
		a := 3
		defer func() {
			a = a + 1
		}()
		return a
	}

	fmt.Println(f1())
	fmt.Println(f2())
}
