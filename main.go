package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
  //to Create a file
  //this reaturns the file and error
	file, _ := os.Create("hello.txt")

  // create file writer
  writer := io.Writer(file)

  // write to the file 
  //reaturns no of bytes and error
  n,err:= writer.Write([]byte("Hello Boi\nSooo\nWhat's up"))
  fmt.Println(n,err)
  n,err = io.WriteString(writer,"Boii")
  fmt.Println(n,err)
  
  // Seek is used to change the file pointer which determines where the file is read or write from 
  // 1st parameter is for setting the offset in bytes here the offset is none 
  // 2nd parameter defines where the pointer is 0 for start 1 for current and 2 for end here it's on the start
  // so the pointer is now at the start with no offset
  file.Seek(0,0)

  // create a file reader 
  reader := io.Reader(file)
  // use buffer to store the read value
  // 40 bytes should be enough for this use case
  buffer:= make([]byte, 40)

  // read the value and sotre it in buffer
  // returns no of bytes and error
  n,err = reader.Read(buffer)
  fmt.Printf("no of bytes: %v \nError: %v \n Buffer value: %v ",n,err,string(buffer))

  // finally close the file
	file.Close()
}
