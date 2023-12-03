package main

import (
	"bufio"
	"fmt"
	"os"
)

func puzzle_1()  {
	file, err := os.Open("./input")
	if err != nil{
		panic("error while reading the input file")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var result_val int
	for scanner.Scan(){
		var first_el *int
		var second_el *int
		current_line := scanner.Bytes()
		i := 0
		e := len(current_line) - 1
		for {
			if i >= len(current_line) || first_el != nil && second_el != nil {
				break
			}
				
			if first_el == nil && current_line[i] >= 48 && current_line[i] <= 58{
				first_n_check := int(current_line[i]) % 48
				first_el = &first_n_check	
			}

			if second_el == nil && current_line[e] >= 48 && current_line[e] <= 58{
				second_n_check := int(current_line[e]) % 48
				second_el = &second_n_check
			}
			i++
			e--
		}
		if first_el == nil && second_el != nil{
			result_val += *second_el
			continue
		}else if first_el != nil && second_el == nil{
			result_val += *first_el
			continue
		}else if first_el != nil && second_el != nil{
			result_val += (*first_el*10)+*second_el
			continue
		}else{
			continue
		}

	}
	fmt.Println("puzzle 1 result",result_val)
}

func main()  {
	puzzle_1()
}


