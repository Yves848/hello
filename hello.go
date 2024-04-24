package hello

import (
	"bufio"
	"flag"
	"fmt"

	"os"
	"unicode/utf8"

	tsize "github.com/kopoli/go-terminal-size"
)

func main2() {
	type test struct {
		name string
		age  int
	}

	var t []test
	name := flag.String("name", "git", "name of the package")
	flag.Parse()
	t = append(t, test{"zhang", 12})
	t = append(t, test{"li", 13})
	fmt.Println(t)
	reader := bufio.NewReader(os.Stdin)
	peekByte, err := reader.Peek(1)
	fmt.Println(peekByte)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(peekByte) > 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
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
	s, err2 := tsize.GetSize()
	if err2 == nil {
		fmt.Println("Current Size is ", s.Width, "x", s.Height)
	}
}
