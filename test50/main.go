package main

import "fmt"

func main(){

	m := map[string][]string{

		"ahmed_f": []string{"football", "tv"},
		"H_f": []string{"go"},
		"HH_f": []string{"a", "b", "c"},
	}


	fmt.Printf("%#v\n", m)

	m["m_f"] = []string{"army", "flag", "jeep"}

	for k, v := range m {
		
		fmt.Println("---------------------")
		fmt.Println(k, v)

		for i, s := range v {

			fmt.Println(i, s)
		}


	}

}
