package main
import
(
    "fmt"
    "time"
)

func producer(threadID int , wg sync.WaitGroup, ch chan string) {
  count := 0
  for {
    time.Sleep(time.Second * 1)
    count ++
    data := strconv.Itoa(threadID) + "---" + strconv.Itoa(count)
    ch <- data
  }
  wg.Done()
}
func consumer(wg sync.WaitGroup, ch chan string) {
  for data := range ch {
    time.Sleep(time.Second * 1)
    fmt.Printf(format: "Consumer", data)

  }
  wg.Done()

}
func main(){
  //多个生产者和消费者
  P_NUM := 3
  C_NUM := 2

  chanStream = make(chan string, 10)
  wgPd := new(sync.WaitGroup)
  wgCs := new(sync.WaitGroup)

  for i := 0; i < P_NUM; i++ {
    wgPd.Add(delta:1)
     go producer(i,wgPd,chanSrteam)
  }

  for i := 0; i < C_NUM; i++ {
    wgCs.Add(delta:1)
     go consumer(wgCs,chanSrteam)
  }

}
