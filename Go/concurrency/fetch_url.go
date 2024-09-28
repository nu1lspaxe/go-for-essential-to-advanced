/*
	ch chan data : 從 channel 寫入/讀取 數據
	ch chan <- data : 從 channel 寫入 數據
	ch <- chan data : 從 channel 讀取 數據

	用 channel 必須使用 sync.WaitGroup -> 否則造成 fatal error: all goroutines are asleep - deadlock!
*/

package concurrency

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func FetchUrl() {
	ch := make(chan string)
	var wg sync.WaitGroup

	for _, url := range os.Args[1:] {
		wg.Add(1)
		go fetch(url, ch, &wg)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	wg.Wait()
}

func fetch(url string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()
	res, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprintf("Error happened : %s", err)
		return
	}

	// 將 res.Body Copy 至 io.Discard stream 中 (可將 io.Discard 當成隨意丟入資訊的trash bin)
	nbytes, err := io.Copy(io.Discard, res.Body)
	res.Body.Close() // 記得關閉 res

	if err != nil {
		ch <- fmt.Sprintf("Error happened : %s", err)
		return
	}

	spend := time.Since(start).Seconds()
	ch <- fmt.Sprintf("URL: %s \nSpend time: %.2f \nData size: %d", url, spend, nbytes)
}
