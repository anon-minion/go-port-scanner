package main
import (
	"fmt"	
	"net"
	"sync"
	"time"
	"os"
)
func main(){
 	var wg sync.WaitGroup
	host := os.Args[1]
	fmt.Println("Searching for Open Ports")
	fmt.Print("Please Be Patient.... \n\n")
	time.Sleep(78888)
	fmt.Println("Port    State     Service")
	for i := 1; i <= 7500; i++ {
		 wg.Add(1)
		go func(port int) {
			var service string
	 		defer wg.Done()
			address := fmt.Sprintf("%v:%d", host, port)
			conn, err := net.Dial("tcp", address)
			if err != nil {
			return
			}
			conn.Close()

			if port == 135{
				service = "msrpc"
			} else if port == 445 {
				service = "microsoft-ds"
			} else {
				service = "Unknown"
			}
			
			fmt.Printf("%v    %v        %v\n", port, "open", service)
		}(i)
	}
	wg.Wait()
}