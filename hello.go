package main

import (
	// "bufio"
	"flag"
	"fmt"

	// "os"
	"unicode/utf8"

	tsize "github.com/kopoli/go-terminal-size"
)

func main() {
	type test struct {
		name string
		age  int
	}

	var t []test
	name := flag.String("name", "git", "name of the package")
	flag.Parse()
	t = make([]test, 3)
	t[0] = test{"zhang", 12}
	fmt.Println(t)
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
	fmt.Println(*name)
	// cmd := exec.Command("winget", "search", "git")
	// out, err := cmd.Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(out))
	st := "百度语音输入                          "

	fmt.Println(len(st))
	fmt.Println(utf8.RuneCountInString(st))

	var s tsize.Size
	s, err := tsize.GetSize()
	if err == nil {
		fmt.Println("Current Size is ", s.Width, "x", s.Height)
	}
}
