package Buffer_IO_File

import (
	"bufio"
	"fmt"
	"os"
	
)

func File() {
    var f *os.File = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}



