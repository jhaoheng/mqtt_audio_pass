package audioSample

import (
  "app/mqtt"
  "fmt"
  "io/ioutil"
  "os"
)

/*
Sender ...
*/
func Sender() {
  fmt.Println("1 : make audio file to binary")
  dat, _ := GetFileToBinary()

  fmt.Println("\n2 : send it by mqtt\n")

  Topic := "audioSample/test"
  Payload := string(dat[:])
  mqtt.Pub(Topic, Payload)
  return
}

/*
GetFileToBinary ...
*/
func GetFileToBinary() (dat []byte, err error) {

  currentDir, _ := os.Getwd()
  fFolder := "audioSample"
  fName := "64kbps"
  fExt := "mp3"

  fPath := currentDir + "/" + fFolder + "/" + fName + "." + fExt
  getFileSize(fPath)
  dat, err = ioutil.ReadFile(fPath)

  return
}

func getFileSize(fPath string) {
  fi, e := os.Stat(fPath)
  if e != nil {
    panic(e)
  }
  // get the size
  size := fi.Size()
  fmt.Println(" File Size is : ")
  fmt.Println("   ", size, " bytes")
  fmt.Println("   ", float64(size)/1000, " KB")
}
