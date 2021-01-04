package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	replacer := strings.NewReplacer("\r\n", "", "\r", "", "\n", "", "%0a", "", "%0d", "")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Json Person printer")
	fmt.Println("======================")
	fmt.Print("Enter name: ")
	nm, _ := reader.ReadString('\n')
	fmt.Print("Enter address: ")
	addr, _ := reader.ReadString('\n')

	nm = replacer.Replace(nm)
	addr = replacer.Replace(addr)

	jsonMap := map[string]string{nm: addr}
	val, _ := json.Marshal(jsonMap)
	fmt.Printf("Json value: %+v", val)
	fmt.Println()
}
