package collector

import (
	"context"
	"time"
)

func main() {
	for {
		price := GenerateMockPrice()
		PublishPrice(context.Background(), price)
		time.Sleep(2 * time.Second)
	}
}
