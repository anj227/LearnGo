package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true 
  var eludedGuards int 
  eludedGuards = rand.Intn(100)
  fmt.Println("Elude guard: ", eludedGuards)
  if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
  }
  fmt.Println("Is heist still on?", isHeistOn)
  if isHeistOn {
    openedVault := rand.Intn(100)
    fmt.Println("Open the vault: ", openedVault)
    if openedVault >= 70 {
      fmt.Println("Grab and GO!")

      leftSafely := rand.Intn(5)
      switch leftSafely{
        case 0:
          isHeistOn = false 
          fmt.Println("Sigh.. failed 1")
        case 1:
          isHeistOn = false 
          fmt.Println("Sigh.. failed 2")
        case 2:
          isHeistOn = false 
          fmt.Println("Sigh.. failed 3")
        case 3:
          fmt.Println("Start the getaway car!")
        case 4:
          fmt.Println("Yes.. we did it!")
      }
    } else {
      fmt.Println("Oops.. can't open the vault :( ")
    }
  } else {
    fmt.Println("Stay Tuned... Heist will be planned for some other day!")
  }
  if isHeistOn {
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Printf("Yeah! We stole $%d \n", amtStolen)
  }
}
