package audioSample

import (
  "app/mqtt"
  "log"
  "os"
)

/*
Receiver ...
*/
func Receiver() {
  Topic := "audioSample/#"
  mqtt.Sub(Topic, func(payload string) {
    dat := []byte(payload)
    GetBinaryToFile("test.mp3", dat)
  })
}

/*
GetBinaryToFile ...
*/
func GetBinaryToFile(fName string, data []byte) {

  currentDir, _ := os.Getwd()
  filePath := currentDir + "/" + fName

  file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
  if err != nil {
    log.Fatalf("failed opening file: %s", err)
  }
  defer file.Close()

  // func (f *File) Write(b []byte) (n int, err error)
  n, err := file.Write(data)
  if err != nil {
    log.Panicln(err)
  }

  log.Println("Save file size : ", n, " bytes")
}
