package main
import (
	"fmt"
	"os"
)
func main() {
	if len(os.Args)!=3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	
	data, err := os.ReadFle(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	result := ProcessText(string(data))

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}