package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input_data := get_input_data()
	answer := parse_data(input_data)
	fmt.Println("Answer: {}", answer)
}

func parse_data(input_data string) int32 {
	segments := strings.Split(input_data, "\n\n")
	seeds := parse_seeds(segments[0])
	fmt.Println(seeds)
	maps := parse_maps(segments[1:])
	fmt.Println(maps)
	return 0
}

func parse_maps(segments []string) map[int][]seed_map_entry {
	maps := make(map[int][]seed_map_entry)
	for i, seg := range segments {
		maps[i] = parse_map_seg(seg)
	}
	return maps
}

func parse_map_seg(seg string) []seed_map_entry {
	var map_entries []seed_map_entry
	entry_lines := strings.Split(seg, "\n")

	for _, line := range entry_lines[1:] {
		if line == "" {
			continue
		}
		part_strings := strings.Split(line, " ")
		parts := string_arr_to_int_arr(part_strings)
		entry := seed_map_entry{dest_range_start: parts[0], source_range_start: parts[1], range_length: parts[2]}
		map_entries = append(map_entries, entry)
	}

	return map_entries
}

func string_arr_to_int_arr(arr []string) []int32 {
	var ints []int32

	for _, val := range arr {
		num, _ := strconv.ParseInt(val, 10, 32)
		ints = append(ints, int32(num))
	}

	return ints
}

func parse_seeds(seed_def string) []int32 {
	data := strings.Replace(seed_def, "seeds: ", "", -1)
	seeds := strings.Split(data, " ")
	return string_arr_to_int_arr(seeds)
}

func get_input_data() string {
	args := os.Args
	fPath := args[1]
	data, _ := os.ReadFile(fPath)
	content := string(data)
	return content
}

type seed_map_entry struct {
	dest_range_start   int32
	source_range_start int32
	range_length       int32
}

const (
	seedsoil int = iota
	soilfert
	fertwater
	waterlight
	lighttemp
	temphumidity
	humiditylocation
)
