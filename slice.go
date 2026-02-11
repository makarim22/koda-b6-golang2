package main

import "fmt"

var scores = []int{50, 75, 66, 20, 32, 90}




func slice(){
newValue := 88
index := 3

newScores := make([]int, len(scores))

copy(newScores, scores)

scores[3]=88
sliceOne := scores[:index]
sliceTwo := newScores[index:]

sliceAdd := append(sliceOne, newValue)
sliceNew := append(sliceAdd, sliceTwo...)
fmt.Println(sliceNew)

 for i, val := range sliceNew {
    fmt.Printf("Indeks: %d, Nilai: %d\n", i, val)
}

}