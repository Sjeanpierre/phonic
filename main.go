package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var Alpha = map[string]string{
	"A": "Alfa",
	"B": "Bravo",
	"C": "Charlie",
	"D": "Delta",
	"E": "Echo",
	"F": "Foxtrot",
	"G": "Golf",
	"H": "Hotel",
	"I": "India",
	"J": "Juliett",
	"K": "Kilo",
	"L": "Lima",
	"M": "Mike",
	"N": "November",
	"O": "Oscar",
	"P": "Papa",
	"Q": "Quebec",
	"R": "Romeo",
	"S": "Sierra",
	"T": "Tango",
	"U": "Uniform",
	"V": "Victor",
	"W": "Whiskey",
	"X": "X-ray",
	"Y": "Yankee",
	"Z": "Zulu",
	"1": "One",
	"2": "Two",
	"3": "Three",
	"5": "Five",
	"6": "Six",
	"7": "Seven",
	"8": "Eight",
	"9": "Nine",
	"0": "Zero",
}

func main() {
        if len(os.Args) < 2 {
          log.Fatal("No input provided")
        }
	string := os.Args[1]
	parts := strings.Split(string, "")
	if len(parts) > 20 {
		log.Fatal("Input too long, replace reduce to 20 characters")
	}

	for _, part := range parts {
		fmt.Printf("%s\n", Alpha[strings.ToUpper(part)])
	}
}
//todo, add flag to use say command on mac
