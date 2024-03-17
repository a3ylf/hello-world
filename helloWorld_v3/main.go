package main

import (
	"os"
  "fmt"
)
//ROUBEI DO V2 MESMO 
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

func fileNameDeliveryService() (TheNameOfTheFile string) {
  return "Hello_world_content.txtxt"
}

func checkIfTheFileExists(TheNameOfTheFile string) bool{
  _, err := os.ReadFile(TheNameOfTheFile)
  if err != nil {
    _, err = os.ReadFile(TheNameOfTheFile)
    if err != nil {
        err = nil
    }
  }
  return true
} 

type Hello_world___struct struct {
  hello_World string  
}
type Hello_World_In_Bytes struct{
  HelloWORLDin_Bytes []byte
}
func HelloWorld_String_To_Byte_Transformer(helloWoRLD Hello_world___struct ) Hello_World_In_Bytes {
    if string([]byte(helloWoRLD.hello_World)) == string([]byte("Hello world!")){
    return Hello_World_In_Bytes{ HelloWORLDin_Bytes: []byte(string([]byte(string([]byte(helloWoRLD.hello_World),
      ),
          ),
        ),
      ),
    }
  }
  return Hello_World_In_Bytes{}
}
func main() {
  var helloWorld Hello_world___struct
  helloWorld = Hello_world___struct{hello_World: ""}
  helloWorld.hello_World = helloWorlder(helloWorld.hello_World,"H")   
  helloWorld.hello_World = helloWorlder(helloWorld.hello_World, "e")   
  helloWorld.hello_World = helloWorlder(helloWorld.hello_World,   "l")   
  helloWorld.hello_World = helloWorlder(helloWorld.hello_World,     "l")   
  helloWorld.hello_World = helloWorlder(helloWorld.hello_World,       "o")   
  helloWorld.hello_World = helloWorlder(helloWorld.hello_World,         " ")   
  helloWorld.hello_World = helloWorlder(helloWorld.hello_World,           "w")   
  helloWorld.hello_World = helloWorlder(helloWorld.hello_World,             "o")   
  helloWorld.hello_World = helloWorlder(helloWorld.hello_World,               "r")   
  helloWorld.hello_World = helloWorlder(helloWorld.hello_World,                 "l")   
  helloWorld.hello_World = helloWorlder(helloWorld.hello_World,                   "d")   
  helloWorld.hello_World = helloWorlder(helloWorld.hello_World,                     "!")   

    var Helloworld Hello_World_In_Bytes 
  Helloworld = HelloWorld_String_To_Byte_Transformer(helloWorld)

  doNothing := func(err error) {
    return
  } 
  writeHELLOWORLD_INTO_Hello_world_content_txtxt_file := func() {
    err := os.WriteFile(fileNameDeliveryService(),[]byte(string(Helloworld.HelloWORLDin_Bytes)),0644)
    if err != nil {
        doNothing(err)
    }
    if err == nil {
      doNothing(err)
    }
  }

  writeHELLOWORLD_INTO_Hello_world_content_txtxt_file()
  
  checkIfTheFileExists(fileNameDeliveryService())
  

  
  var HELLOWORLDREADFROMFILE Hello_World_In_Bytes
  HELLOWORLDREADFROMFILE.HelloWORLDin_Bytes, _ = os.ReadFile(fileNameDeliveryService())
  helloWorldPrinter(string(string(string(string(HELLOWORLDREADFROMFILE.HelloWORLDin_Bytes)))))

  if "Hello world!" == "Hello world!" {
                  return
  }  

  
}
