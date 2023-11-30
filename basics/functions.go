package main

import "fmt"
//It receives 2 strings and returns a string.
//Since here both parameters are strings, it is not necessary to declare both as strings
func concat(s1, s2 string) string {
  return s1 + s2
}

func functions(){
  fmt.Println(concat("Hello", "World from functions!"))   
}

func yearsUntilEvents(age int)(
  yearsUntilAdult,
   yearsUntilDrinking, 
   yearsUntilCarRental int){

  yearsUntilAdult = 18 - age
  if yearsUntilAdult < 0 {
    yearsUntilAdult = 0
    fmt.Println(yearsUntilAdult)
  }

   yearsUntilDrinking = 21 - age
  if yearsUntilDrinking < 0 {
    yearsUntilDrinking = 0
    fmt.Println(yearsUntilDrinking)
  }

  yearsUntilCarRental = 25 - age
  if yearsUntilCarRental < 0 {
    yearsUntilCarRental = 0
    fmt.Println(yearsUntilCarRental)
  }
  return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
}

var result int
func divide(dividend, divisor int)int{
  if divisor == 0{
    result = 0
    return result
  }
  result = dividend/divisor
  return  result
}