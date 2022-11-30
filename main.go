package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		url := "http://192.168.1.37:8099/api/tranform?url=http://192.168.2.202/index.php/Bashu/Report/310997&targetPath=/巴蜀/巴蜀"
		url += ".pdf&callBack=https://www.baidu.com"
		fmt.Println(url)
		//http.Get(url)
	}
}
