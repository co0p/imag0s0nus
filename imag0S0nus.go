package main

import "fmt"

import (
  "log"
  "os"
  "strconv"
	"image"

  _ "image/png"
)


type Pixel struct {
  x int64
  y int64
  c float64
}

func main() {

  filename := ""
  startPosition := new(Pixel)

  // read arguments
  if len(os.Args) < 4 {
    fmt.Printf("usage: %v <path/to/image> <startpos x> <startpos y>\n", os.Args[0])
    return
  }

  filename = os.Args[1]

  x, err := strconv.ParseInt(os.Args[2], 10, 64)
  if err != nil {
    log.Fatalf("Failed reading <startpos x>")
  }
  startPosition.x = x

  y, err := strconv.ParseInt(os.Args[3], 10, 64)
  if err != nil {
    log.Fatalf("Failed reading <startpos y>")
  }
  startPosition.y = y



  // read image
  fmt.Printf("loading file '%v', using %v,%v as a starting point\n", filename, startPosition.x, startPosition.y)
  // Decode the JPEG data. If reading from file, create a reader with
	//
  reader, err := os.Open(filename)
  if err != nil {
     log.Fatal(err)
  }
  defer reader.Close()
  m, _, err := image.Decode(reader)
  if err != nil {
  	log.Fatal(err)
  }
  bounds := m.Bounds()
  fmt.Printf("bounds:%v", bounds)

  // extract sound

  // play sound

}
