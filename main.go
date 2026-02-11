package main
import "fmt"



type World struct{
	Message string
}

type Hello struct{
	World World
}


type best struct{
	Message string
}

type the struct{
	best best
}

type are struct{
	the the
}

type we struct{
	are are
}


func main (){
  myObj := Hello{
	World: World{
		Message: "Hello World",
	},
  }

  myobj2 := we{
	are: are{
		the: the{
			best: best{
				Message: "koda",
			},
		},
	},
  }

  fmt.Println(myObj.World.Message)

  fmt.Println(myobj2.are.the.best.Message)
}
