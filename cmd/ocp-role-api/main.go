package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, ocp-role-api!")
}

func OpenAndPrintFileInCycle(fname string, n int) error {
	data := make([]byte, 128)

	for i := 0; i < n; i++ {
		err := func() error {
			f, err := os.Open(fname)
			if err != nil {
				return fmt.Errorf("can't open file: %w", err)
			}
			defer f.Close()

			for {
				cnt, err := f.Read(data)
				if err != nil && err != io.EOF {
					return fmt.Errorf("can't read file: %w", err)
				}
				if cnt == 0 {
					break
				}
				fmt.Printf("%s\n", data[:cnt])
			}

			return nil
		}()

		if err != nil {
			return err
		}
	}
	return nil
}
