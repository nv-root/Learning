// context provides a signal for cancellation, timeout, or deadlines to goroutines so they can obey to stop working
package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func workWithCancel(ctx context.Context) {
	for {
		select {
		//listening for cancel signal
		case <-ctx.Done():
			fmt.Println("Stopping...")
			return
		default:
			// if not, keep working
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func workWithTimeout(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Timeout:", ctx.Err())
			return
		default:
			fmt.Println("Working with timeout...", time.Now())
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func checkEvenOdd(ctx context.Context, num int) string {
	select {
	case <-ctx.Done():
		return "Operation cancelled"
	default:
		if num%2 == 0 {
			return fmt.Sprintf("%d is even", num)
		} else {
			return fmt.Sprintf("%d is odd", num)
		}

	}
}

func logWithContext(ctx context.Context) {
	log.Printf("Request Id log: %v", ctx.Value("reqId"))
}

func main() {
	// rootContext of all contexts with no cancel, or deadline, or value
	rootContext := context.Background()

	ctx1, cancel := context.WithCancel(rootContext)
	go workWithCancel(ctx1)

	ctx2, cancel2 := context.WithTimeout(rootContext, 3*time.Second)
	go workWithTimeout(ctx2)
	defer cancel2()

	time.Sleep(4 * time.Second)
	cancel()

	ctx3 := context.WithValue(rootContext, "reqId", "req123")
	reqId := ctx3.Value("reqId")
	if reqId != nil {
		fmt.Println("Request ID:", reqId)
	} else {
		fmt.Println("No request ID found")
	}

	logWithContext(ctx3)

	// used when the right context policy is not known yet
	ctxUnkown := context.TODO()
	result := checkEvenOdd(ctxUnkown, 5)
	fmt.Println(result)
}
