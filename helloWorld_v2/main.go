package main

import (
	"fmt"
	"os"
)


func helloWorlder(SuperUltraMegaHelloWorldString string,TheImportantCharThatIsGointToBeAddedToTheSuperUltraMegaHelloWorldString string) string {
                  var  TheNowSuperUltraMegaHelloWorldStringWithItsNewAdditionThatIsTheImportantCharThatIsGointToBeAddedToTheSuperUltraMegaHelloWorldString string 
  
              _, err := os.ReadFile("ThisFileReadingHasNoPurposeWhatsoever.ThisIsCompletelyPointless")
        if err != nil                         {
  TheNowSuperUltraMegaHelloWorldStringWithItsNewAdditionThatIsTheImportantCharThatIsGointToBeAddedToTheSuperUltraMegaHelloWorldString =  fmt.Sprintf(SuperUltraMegaHelloWorldString + TheImportantCharThatIsGointToBeAddedToTheSuperUltraMegaHelloWorldString)    }
        return TheNowSuperUltraMegaHelloWorldStringWithItsNewAdditionThatIsTheImportantCharThatIsGointToBeAddedToTheSuperUltraMegaHelloWorldString 
}
func helloWorldPrinter(helloWorld string){
      if string([]byte(helloWorld)) == "Hello world!" {
    fmt.Println(helloWorld)
  }
}
func main() {
    var helloWorld string
  helloWorld = helloWorlder         (helloWorld,"H")
        helloWorld = helloWorlder(helloWorld,"e"         )
  helloWorld = helloWorlder(helloWorld,"l"   , 
    )
  helloWorld = helloWorlder(helloWorld,"l")
  helloWorld = helloWorlder(      helloWorld,"o")
            helloWorld = helloWorlder(helloWorld," ")
  helloWorld = helloWorlder(helloWorld      ,"w")
  helloWorld             =              helloWorlder  (helloWorld,"o")
  helloWorld =        helloWorlder(helloWorld,         "r")
  helloWorld          = helloWorlder(helloWorld,"l")
              helloWorld =     helloWorlder(helloWorld,"d")
  helloWorld = helloWorlder(helloWorld,"!")
  
  helloWorldPrinter(helloWorld);;;;
}
