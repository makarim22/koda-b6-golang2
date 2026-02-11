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

type num struct{
	First first
	Second second
}

type first []int

type second []int


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

	num := num{
		First: first{16,16,20,20,},
		Second: second{16,15,16,10,},
	}

  fmt.Println(Hello.World)

  fmt.Println(we.are.the.best)

  fmt.Println(num.First[1] + num.Second[2])
}


 


