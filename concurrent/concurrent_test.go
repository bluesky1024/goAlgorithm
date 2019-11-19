package concurrent

import (
	"fmt"
	"testing"
	"time"
)

func TestPrintMultiInOrder(t *testing.T) {
	fmt.Println("start")
	PrintMultiInOrder()
	time.Sleep(5 * time.Second)
}

func TestGenDealLock(t *testing.T) {
	GenDealLock()
}
