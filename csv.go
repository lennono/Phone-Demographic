package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func differentiate(number string) (string, string) {
	a := []rune(number)
	length := len(number)
	var NPA = ""
	var NXX = ""

	// If statements for different number lengths
	// Number length can vary based on an iternational number area codes
	if(length == 10){
		for i, r := range a {
			if i < 3 {
				NPA = NPA + string(r)
			} else if i > 2 && i < 6 {
				NXX = NXX + string(r)
			}
		}
	} else if(length == 12){
		for i, r := range a {
			if i > 1 && i < 5 {
				NPA = NPA + string(r)
			} else if i > 4 && i < 8 {
				NXX = NXX + string(r)
			}
		}
	} else{
		for i, r := range a {
			if i > 0 && i < 4 {
				NPA = NPA + string(r)
			} else if i > 3 && i < 7 {
				NXX = NXX + string(r)
			}
		}
	}
	return NPA, NXX
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter first number: ")
	Originating_Number, _ := reader.ReadString('\n')
	fmt.Println(Originating_Number)

	reader1 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter second number: ")
	Full_Terminating_number, _ := reader1.ReadString('\n')
	fmt.Println(Full_Terminating_number) 

	var NPA = ""
	var NXX = ""
	var cliNPA = ""
	var cliNXX = ""

	NPA, NXX = differentiate(Full_Terminating_number)
	cliNPA, cliNXX = differentiate(Originating_Number)

	url := fmt.Sprintf("http://localcallingguide.com/lca_rcdist.php?npa1=%s&nxx1=%s&npa2=%s&nxx2=%s", cliNPA, cliNXX, NPA, NXX)
	resp, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
		
	if err != nil {
		panic(err.Error())
	}

	ts := string(body)
	text := string(`<td headers="olocalcall" data-label="Local call+">Y</td>`) // No need to work through the html document, returns either true or false, based on if the call was local or not 

	if strings.Contains(ts, text) == true {
		fmt.Print("Local call")
	} else {
		fmt.Print("Non-Local call")
	}
}
