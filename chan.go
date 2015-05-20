package main

import "fmt"
import "time"
 
func Message2() {
        fmt.Println("Message2..");
}

func Message1(chs []chan string) {
        chs[0] <- "hello";
        chs[1] <- "World";
        fmt.Println("Message1...");
}

func main() {

    chs := make([]chan string, 2);
    
    for i:= 0; i<2; i++{
        chs[i] = make(chan string);
    }
    
    go Message1(chs);
    go Message2();

    //写入到channle中的数据要读出来，否则不会有输出
    msg1 := <-chs[0];
    msg2 := <-chs[1];

    fmt.Println(msg1, msg2);
    time.Sleep(2);
}
