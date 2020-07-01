package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	defer writer.Flush()

}

var writer = bufio.NewWriter(os.Stdout)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Buffer(make([]byte, 256), 1e9)
	sc.Split(bufio.ScanWords)
}

func Println(data interface{}) {
	writer.WriteString(fmt.Sprintf("%v\n", data))
}

func Print(data interface{}) {
	writer.WriteString(fmt.Sprintf("%v", data))
}

func Next() string {
	sc.Scan()
	return sc.Text()
}

func NextInt() int {
	i, err := strconv.Atoi(Next())
	if err != nil {
		panic(err)
	}
	return i
}

func NextInt64() int64 {
	i64, err := strconv.ParseInt(Next(), 10, 64)
	if err != nil {
		panic(err)
	}
	return i64
}

func NextFloat64() float64 {
	f, err := strconv.ParseFloat(Next(), 64)
	if err != nil {
		panic(err)
	}
	return f
}
