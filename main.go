// mySender project main.go
package main

import (
	"fmt"
	"mySender/src/util"
	"net/http"
	"io/ioutil"
	"strings"
	"strconv"
	"os"
	"errors"
	"time"
	"sync"
)

var request *http.Request
var client *util.Client

type SenderMap struct {
	lock sync.RWMutex
	senderMap map[string]int
}

var SM *SenderMap

func init() {
	SM = &SenderMap{
		senderMap: map[string]int{"alllines":0,"successLine":0},
	}
}

func main() {
	fmt.Println("Hello World!")
	usl := util.Input.Usl
	th, err1 := getIntInput(util.Input.Thread, 10)
	conns, err2 := getIntInput(util.Input.Connections, 10)
	dTime, err3 := getIntInput(util.Input.Dtime, 5)
	if err1 != nil || err2 != nil || err3 != nil{
		fmt.Println(util.HelpInfo)
		os.Exit(0)
	}
	var header map[string]string
	if h := util.Input.Header; h != "" {
		if hIndex := strings.Split(h, ":"); len(hIndex) == 2{
			header = make(map[string]string)
			header[hIndex[0]] = hIndex[1]
		}
	}
	var err error
	request, err = util.NewRequest(http.MethodGet, usl, http.NoBody, header)
	if err != nil {
		fmt.Println("Create Request Client err, please check Inputs")
	}
	fmt.Println(header)
	client = util.NewClient()

	t := th
	j := 0
	// chaThIndex := make([]chan int, dTime)
	for i := 0; i < dTime; i++ {
		// chaThIndex[i] = make(chan int, th)
		t = th
		for t > 0 {
			// go toRequest(chaThIndex[i], conns)
			go toRequest(conns)
			t--
		}
		j = 0
		fmt.Println("j:",j)
		fmt.Println("th",th)
		//for j < th {
		//	for len(chaThIndex[i]) > 0{
		//		a, _ := <- chaThIndex[i]
		//		j =+ a
		//	}
		//}
		// close(chaThIndex[i])
		time.Sleep(time.Second)
	}
	//alllines := SM.senderMap["alllines"]
	//successLine := SM.senderMap["successLine"]
	//fmt.Printf("Running %ds test with %s\n", dTime, usl)
	//fmt.Printf("%d threads %d connections\n", th, conns)
	//fmt.Printf(`all line size : %d  success size: %d
	//access rate : %d/s`, alllines, successLine, successLine/dTime)
	//fmt.Println("\n")
}

// func toRequest(ch chan int, conns int){
func toRequest(conns int){
	//defer func() {
	//	ch <- 1
	//	SM.lock.Unlock()
	//}()
	var chs, cha int
	for conns > 0 {
		response, err := client.SendRequest(request)
		if err == nil {
			rb , _ := ioutil.ReadAll(response.Body)
			if response.Status == "200 OK" {
				chs++
			}
			fmt.Printf("get reponse success, res: %s, status: %s, response header: %s\n,", string(rb), response.Status, string(response.Header))
		}else {
			fmt.Errorf("send request err, %s \n", err.Error())
		}
		cha++
		conns--
	}
	//SM.lock.Lock()
	//SM.senderMap["successLine"] =+ chs
	//SM.senderMap["alllines"] =+ cha
}

func getIntInput(s string, df int) (int, error) {
	if s == "" {
		return df, nil
	} else {
		if t, err := strconv.Atoi(s); err != nil {
			return 0, errors.New("input err")
		}else {
			return t, nil
		}
	}
}
