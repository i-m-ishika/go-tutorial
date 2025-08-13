package main

import (
  "fmt"
  "errors"
)

func printMe(param string) {
  fmt.Println("Hello, " + param)
}

func float32Sum(a float32, b float32) float32 {
  return a+b
}

func intDivision(nr int, dr int) (int, int, error) {
  if(dr==0){
    return 0, 0, errors.New("Cannot divide by zero.")
  }
  return nr/dr, nr%dr, nil
}

func summariseDivision(rem int) {
  switch rem{
  case 0:
    fmt.Printf("Exact division")
  case 1: 
    fmt.Printf("Close division")
  default:
    fmt.Printf("Not close division")
  }
}

func main(){
  printMe("User")
  var intNum int
  intNum = 2
  var floatNum float32 = 123.4567
  fmt.Println(float32Sum(floatNum, float32(intNum)))
  res, rem, err := intDivision(intNum+5, intNum)
  if err!=nil{
    fmt.Printf(err.Error())
  } else {    
    fmt.Printf("Result: %v, Remainder: %v\n", res, rem)
  }

//  TODO: arrays and slices
//    capacities, length

//  maps
  var myMap map[string]uint8 = make(map[strin:g]uint8) // key: string val: uint8
  fmt.Println(myMap)
  myMap2 := map[string]uint8{"Adam": 23, "Sarah": 24}
  fmt.Println(myMap2["Adam"])
  fmt.Println(myMap2["A"]) // returns 0. maps in GO always return something, even when key doesn't exist
  // to check if key exists in map
  age, ok := myMap2["Jessica"]
  if ok{
    fmt.Printf("Age: %v", age)
  }else{
    fmt.Println("Jessica not in map")
  }

  //  TODO: arrays and maps iteration
}
