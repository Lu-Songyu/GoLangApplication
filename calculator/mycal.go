package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your first num: ")
	num1, _ := reader.ReadString('\n')
	float1, err1 := strconv.ParseFloat(strings.TrimSpace(num1), 64)
	if err1 != nil {
		fmt.Println("Error in Num1: ", err1)
	}
	fmt.Println("Enter your second num: ")
	num2, _ := reader.ReadString('\n')
	float2, err2 := strconv.ParseFloat(strings.TrimSpace(num2), 64)
	if err2 != nil {
		fmt.Println("Error in Num2: ", err2)
	}

	fmt.Println("Choose your operation")
	input, _ := reader.ReadString('\n')
	op := strings.TrimSpace(input)
	var float3 float64
	switch op {
	case "+":
		float3 = float1 + float2
	case "-":
		float3 = float1 - float2
	case "*":
		float3 = float1 * float2
	case "/":
		float3 = float1 / float2
	default:
		fmt.Println(float3)

	}
	float3 = math.Round(float3*100) / 100

	fmt.Printf("num %v %v num %v is %v\n", float1, op, float2, float3)

}
