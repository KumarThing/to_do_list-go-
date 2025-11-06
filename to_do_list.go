package main
import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"encoding/json"

)

func main() {
	newReader := bufio.NewReader(os.Stdin)
	var toDoList [] string

	fmt.Println("To do list:")
	fmt.Println("add list")
	fmt.Println("show list")
	fmt.Println("exit\n")

	for {
		input, _ := newReader.ReadString('\n')
		input = strings.TrimSpace(input)

		 if input == "add list" {
			fmt.Print("- input list: ")

			list, _ := newReader.ReadString('\n')
			list = strings.TrimSpace(list)

			exists := false
			for _ , item := range toDoList {
				if strings.EqualFold(list, item) {
					exists = true 
					break
				} 
			}
			if exists {
				fmt.Println("Item alredy existed.")
			} else {

			toDoList = append(toDoList, list)
				fmt.Println("list added",list)}
		
			
			


		} else if input == "show list" {
			fmt.Println("\nYour to do list:")
			if len(toDoList) == 0 {
				fmt.Println("No items yet!")
				
			} else {
				for i, list := range toDoList{
					fmt.Printf("%d. %s\n", i+ 1, list)

					
				}
			}
		}  else if input == "exit" {
			fmt.Println("Byee")
			break
		} else if input == "save file" {
			savefile(toDoList)
		} else if input == "load file" {
			loadfile(toDoList)
			
		} else {
			fmt.Println("Enter correct operation. Error!!!")
			continue
		}
		fmt.Println("\nTo do list:")
		fmt.Println("add list")
		fmt.Println("show list")
		fmt.Println("exit\n")
		}



	}
	func savefile(toDoList[] string) {
		file,err := os.Create("to_do_list.json")
		if err != nil {
			fmt.Println("Error saving file",err)
			return
		}
		defer file.Close()

		data, err := json.MarshalIndent(toDoList, "", " ")
		if err != nil {
			fmt.Println("Error conveting txt to json.",err)
			return
		}
		file.Write(data)
		fmt.Println("File saved.")
	}

	func loadfile(toDoList[] string) {
		data, err := os.ReadFile("to_do_list.json")
		if err != nil {
			fmt.Println("Error loading file",err)
			return
		}
		err = json.Unmarshal(data, toDoList)
		if err != nil {
			fmt.Println("Error parsing JSON", err)
		}
		fmt.Println("To_do_list data loaded successfully.")

	}
