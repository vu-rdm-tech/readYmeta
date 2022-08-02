// readYmeta.go

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func errcntrl(e error) {
	if e != nil {
		panic(e)
	}
}

func read_json_branch(json_map map[string]interface{}) {
	fmt.Println(" ")
	fmt.Println("json_map")
	fmt.Println(json_map)
	fmt.Println(reflect.TypeOf(json_map))

	for k, v := range json_map {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			fmt.Println(reflect.TypeOf(k))
			fmt.Println(reflect.TypeOf(v))
			fmt.Println(reflect.TypeOf(vv))
			read_array(vv)
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}

func read_array(jarr []interface{}) {
	fmt.Println("->")
	fmt.Println(reflect.TypeOf(jarr))
	for i, u := range jarr {
		fmt.Println("-->")
		fmt.Println(i, u)
		fmt.Println(reflect.TypeOf(i))
		fmt.Println(reflect.TypeOf(u))
		if reflect.TypeOf(u) == reflect.TypeOf(jarr) {
			fmt.Println("XXX")
		}

	}
}

func main() {
	msg := "Hello World! Again!"
	fmt.Println(msg)

	json_file, err1 := os.ReadFile("yoda-metadata.json")
	errcntrl(err1)

	// print the file cast as string
	fmt.Print(string(json_file))

	// create interface to json file
	var json_dat interface{}
	// Unmarshal json file to interface
	err2 := json.Unmarshal(json_file, &json_dat)
	errcntrl(err2)
	//create a dictionary of the json interface data
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(json_dat)
	read_json_branch(json_dat.(map[string]interface{}))

	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(reflect.TypeOf(json_dat))
	fmt.Println(json_dat)
	fmt.Println(" ")
	fmt.Println(reflect.TypeOf(json_dat.(map[string]interface{})))
	fmt.Println(json_dat.(map[string]interface{}))

}

/*
Funky GO template builder: https://mholt.github.io/json-to-go/
*/

/*
func old_main_for_history() {
	msg := "Hello World! Again!"
	fmt.Println(msg)

	var buff [32]byte
	fmt.Println(buff[10])

	var buffmulti [5][5]int
	cntr := 1
	for i := range buffmulti[0] {
		for j := range buffmulti[0] {
			fmt.Println(i, j)
			buffmulti[i][j] = cntr
			cntr++
		}
	}

	fmt.Println(buffmulti[:])
	fmt.Println(buffmulti[1])
	fmt.Println(buffmulti[0:1])
	fmt.Println(buffmulti[0:2])

	var numbers [10]int
	for i := 0; i < cap(numbers); i++ {
		numbers[i] = i + 1
	}

	fmt.Println(numbers)

	type point struct {
		x, y int
	}

	var p = point{10, 20}
	fmt.Println(p.x)
	fmt.Println(p.y)

	var m = make(map[string]int)
	m["mike"] = 30
	m["lucy"] = 40
	fmt.Println(m["lucy"])

	var m2 = make(map[int]string)
	m2[50] = "Mike"
	m2[60] = "Lucy"
	fmt.Println(m2[40])
	fmt.Println(m2[50])

}
*/
