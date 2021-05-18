package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
	"github.com/Enviy/pipelineBuilder/builder"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	fmt.Println("Enter age: ")
	age, _ := reader.ReadString('\n')
	age = strings.TrimSuffix(age, "\n")

	builder.BuildScript(name, age)
}
