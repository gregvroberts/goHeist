package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {

  //provide a unique seed value
  rand.Seed(time.Now().UnixNano())

  
  isHeistOn := true //how successful our heist is
  eludedGuards := rand.Intn(100) //used to simulate if we made it past the guards or not

  if eludedGuards >= 50 {
    fmt.Print("Looks like you've managed to make it past the guards. Good job")
  } else {
    isHeistOn = false
    fmt.Print("Plan a better disguise next time?")
  }

  openedVault := rand.Intn(100) //used to determine if we can open the vault
if isHeistOn && openedVault >= 70 {
  fmt.Print("Grab and GO!")
} else if isHeistOn {
  isHeistOn = false
  fmt.Print("Oh no! The vault can't be opened")
}

leftSafely:= rand.Intn(5)
if isHeistOn {
  switch leftSafely{
  case 0:
    isHeistOn = false
    fmt.Print("You got caught on your way out!")
  case 1:
    isHeistOn = false
    fmt.Print("Turns out the vault doors don't open from the inside...")
  case 2:
    isHeistOn = false
    fmt.Print("You didn't account for security lasers....")
  case 3:
    isHeistOn = false
    fmt.Print("The guards released tear gas...")
  default:
    fmt.Print("Start the getaway car!")
  }
}
if isHeistOn {
  //how much was stolen in the heist
  amtStolen := 10000 + rand.Intn(1000000)
  fmt.Print("You successfully stole $", amtStolen)
}

  fmt.Print(isHeistOn)

}
