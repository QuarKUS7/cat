package main

import (
	//"flag"
	"fmt"
	"io"
	"os"
)

func readAndPrint(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	buf := make([]byte, 8*1024)

	for {
		n, err := file.Read(buf)

		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			//Zaloguj chybu
			break
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Missin filename")
		os.Exit(1)
	}
	for _, filename := range os.Args[1:] {
		readAndPrint(filename)
	}

}
