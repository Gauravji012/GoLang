package main

import (
	// "bufio"
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }
	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// fmt.Println(fileInfo.IsDir(), fileInfo.ModTime(), fileInfo.Mode(), fileInfo.Name(), fileInfo.Size(), fileInfo.Sys())

	//read the file
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// buf := make([]byte, 25)
	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }
	// for i := 0; i < len(buf); i++ {
	// 	fmt.Println(d, string(buf[i]))
	// }

	// reading file with another approach
	// data, err := os.ReadFile("example.txt") // but this fun loads complete data at a time so if we have large size data tthen it loads & memory may leak

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	//read folder
	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// // fileInfo, err := dir.ReadDir(2) // ReadDir(-1) gives all folders
	// fileInfo, err := dir.ReadDir(-1)
	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	//create a file
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// f.WriteString("hi go")
	// f.WriteString("nice language")
	// fmt.Println("")
	// bytes := []byte("hey golang")
	// f.Write(bytes)
	// f.WriteAt(bytes, 0) // write from the offset pointer // attach the previous present data after offset content

	// read data from 1 file & transfer it to another file using streamming fashion / approach.

	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer destFile.Close()

	// // read the data with the help of buffer.io
	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }
	// writer.Flush()
	// fmt.Println("written to new file successfully")

	// delete a file
	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(" file deleted successfully ")
}
