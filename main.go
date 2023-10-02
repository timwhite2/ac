package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func main() {
	fmt.Println("begin kmpVsNative -------------------------------")
	kmpVsNative()
	fmt.Println("end kmpVsNative -------------------------------")
	fmt.Println()
	fmt.Println()
	fmt.Println("begin acVsKmp -------------------------------")
	acVsKmp()
	fmt.Println("end acVsKmp -------------------------------")
	//GenBtcAddr(1)
}

func kmpVsNative() {
	content, err := os.ReadFile("files/kmp_text.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	filename := "files/kmp_pattern.txt"
	pattern, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	pastTime(false, text, []string{string(pattern)}, NaiveMatch)
	pastTime(false, text, []string{string(pattern)}, KMP)
}

func acVsKmp() {
	content, err := os.ReadFile("files/verbos3.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	filename := "files/addresses_1000.txt"
	addresses, err := ReadAddressesFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	pastTime(true, text, addresses, KMP)
	pastTime(true, text, addresses, AC)
}

func pastTime(isPrintPattern bool, text string, addresses []string, fun func(text string, addresses []string) map[int][]string) {
	before := time.Now()
	matches := fun(text, addresses)
	after := time.Now()
	duration := after.Sub(before)

	fmt.Printf("func %s use %.3f s\n", getFunctionName(fun), duration.Seconds())
	for position, ptrs := range matches {
		if isPrintPattern {
			fmt.Printf("in %d match pattern: %s\n", position, ptrs)
		} else {
			fmt.Printf("in %d match \n", position)
		}
	}
}

func getFunctionName(f interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	parts := strings.Split(fullName, ".")
	return parts[len(parts)-1]
}

func ReadAddressesFromFile(filename string) ([]string, error) {
	var addresses []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		address := scanner.Text()
		addresses = append(addresses, address)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return addresses, nil
}
