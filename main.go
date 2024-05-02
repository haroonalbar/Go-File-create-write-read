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
  // n,err:= writer.Write([]byte("Hello Boi\nSooo\nWhat's up"))
  // fmt.Println(n,err)
  n,err := io.WriteString(writer,"Boii")
  fmt.Println(n,err)
  
  // Seek is used to change the file pointer which determines where the file is read or write from 
  // 1st parameter is for setting the offset in bytes here the offset is none 
  // 2nd parameter defines where the pointer is 0 for start 1 for current and 2 for end here it's on the start
  // so the pointer is now at the start with no offset
  // file.Seek(0,0)

  // open  file
  file,_ = os.Open("hello.txt")
  // create a file reader 
  reader := io.Reader(file)
  // use buffer to store the read value
  // 40 bytes should be enough for this use case
  // buffer:= make([]byte, 1)

  // read the value and sotre it in buffer
  // returns no of bytes and error
  // n,err = reader.Read(buffer)
  // fmt.Printf("no of bytes: %v \nError: %v \n Buffer value: %v ",n,err,string(buffer))

  //using minibuffer of size 1 and read all content in the file
  // mini_buffer:= make([]byte, 1)
  // for {
  // n,err = reader.Read(mini_buffer)
  // fmt.Printf("%v %v %v \n ",string(mini_buffer),n,err)
  //   if err != nil{
  //     break
  //   }
  // }

  // other way use ReadAll
  // this is not a good option for big files to read on 
  // since the buffer size would be equal to the content in the file 
  // which will require a huge memory

  buffer, err := io.ReadAll(reader)
  fmt.Printf("%v %v %v \n ",string(buffer),err,len(buffer))

  seeker := reader.(io.Seeker)

  seeker.Seek(-2,2)

  buffer, err = io.ReadAll(reader)
  fmt.Printf("%v %v %v \n ",string(buffer),err,len(buffer))

  // finally close the file
	file.Close()
}
