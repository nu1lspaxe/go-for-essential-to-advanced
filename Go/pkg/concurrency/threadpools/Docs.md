# Thread Pools (Work Pool)
若是 goroutine 產生數量規模過大，會導致內存暴漲、處理效率下降甚至引發程序崩潰。使用緩衝佇列可以在一定程度上提高併發，但並不是最佳解，因為當 goroutine 請求速度大於佇列的處理速度時，緩衝區一樣會被塞滿，而後面的請求被阻塞。

在線程池(Thread Pools)的概念中，可以定義最大 goroutine 數量，當 Job Queue 中還有 job 待完成時，從 work pool 中取一個available的 worker (從 Job Queue 中取出 job) 來執行，如此既可保障 goroutine 的可控性，也提高併發處理能力。

<img src='https://miro.medium.com/v2/resize:fit:3756/1*xe4DmSW7U1PNY8vzryKZ6Q.png'>

<br>

## Concept : WaitGroup, Channel, Multi Goroutine
- WaitGroup: Multiple goroutine conduct a same task
- Channel + Select: 主動通知停止

- 多個 Goroutine ? Or Goroutine in Goroutine ?
    > Use Context