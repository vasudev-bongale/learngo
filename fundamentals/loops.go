package main

import "fmt"

/*
There is only one Loop construct i.e, For Loop
"for" can be used in different ways to achieve a while loop, an infinite loop,
a traditional for loop
*/

func main()  {

  // for loop with For clause: (int; cond; post)
  for i := 0; i < 4; i++ {
    fmt.Println(i)
  }

  // while loops
  j := 0
  for j < 5 {
    fmt.Println(j)
    j++
  }

  var runtimes int = 5
  //Infinite loop with Break & Continue
  k := 0
  for {
    if k > runtimes {
      break;
    }
    if k % 2 == 0 {
      k++
      continue;
    }
    fmt.Println("I am running forever!")
    k++
  }
}
