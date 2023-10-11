package concur

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type itemId string

type uerReq struct {
	name     string
	item     itemId
	quantity uint
	price    uint
}
type itemsList []string

var Items = itemsList{"phone", "laptop", "keyboard"}

func RequestGenerator(ctx context.Context, n int, name string, items itemsList, ch1 chan uerReq) {
	for i := 0; i < n; i++ {
		newRequest := uerReq{
			name:     name,
			item:     itemId(items[rand.Intn(len(items))]),
			quantity: uint(rand.Intn(5)),
			price:    uint(rand.Intn(100)),
		}
		fmt.Printf("sum:%d\n", (newRequest.price * newRequest.quantity))
		ch1 <- newRequest
	}
	close(ch1)
}

func ProcessesRequests(ctx context.Context, ch1 chan uerReq) {
	var total int
	for newRequest := range ch1 {
		total += int(newRequest.price * newRequest.quantity)
		fmt.Println("done")
	}

	fmt.Println("total:", total)

}

func ShopProg() {
	// Створення контексту з дедлайном 2 секунди
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Пам'ятайте про закриття контексту

	ch1 := make(chan uerReq)
	done := make(chan bool) // Канал для сигналу про завершення

	go RequestGenerator(ctx, 30, "name", Items, ch1)
	go func() {
		ProcessesRequests(ctx, ch1)
		done <- false // Сигнал про завершення горутини randIntToChan
	}()

	select {
	case <-ctx.Done():
		// Операція завершилася через дедлайн або була скасована
		err := ctx.Err()
		if err == context.DeadlineExceeded {
			fmt.Println("Операція не встигла завершитися протягом дедлайну")
		} else if err == context.Canceled {
			fmt.Println("Операція була скасована")
		}
	case <-done:
	}
}
