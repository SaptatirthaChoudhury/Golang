package Buffer_IO_File

import (
	"bufio"
	"fmt"
	"os"
	
)

func Buffer() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// fmt.Print(scanner.Scan())
		// fmt.Println(">", scanner.Text())

		text := scanner.Text()
		if text == "exit" {
			break
		}

		fmt.Println(">", text)
	}
}



