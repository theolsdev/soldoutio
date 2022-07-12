package main

import (
	"alternance/modules"
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {
	var product string
	fmt.Println("Enter your product")
	fmt.Scanln(&product)

	wg.Add(1)
	go func() {

		s := modules.LeclercSession{}

		// Note : Normalement c'est la variable product ici mais étant donnée que ce n'est pas dynamique j'ai laissé le lien directement
		s.InitSession("https://www.e.leclerc/fp/bottines-en-cuir-a-lacet-outdor-7640305958908")
		for {
			fmt.Println("Adding to cart...")
			err := s.AddToCart()
			if err == nil {
				break
			} else {
				fmt.Println("Error ATC, retrying...")
				time.Sleep(1 * time.Second)
			}

		}
		fmt.Println("Successfully adding to cart!")
		wg.Done()
	}()

	wg.Wait()

}
