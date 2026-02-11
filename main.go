package main
import "fmt"



// type World struct{
// 	Message string
// }

type Hello struct{
	World string
}


// type best struct{
// 	Message string
// }

type the struct{
	best string
}

type are struct{
	the the
}

type we struct{
	are are
}


func main () {
  Hello := Hello{
	World:  "Hello World",
  }

  we := we{
	are: are{
		the: the{
			best: "koda",
			},
		},
	}

  fmt.Println(Hello.World)

  fmt.Println(we.are.the.best)
}


 


