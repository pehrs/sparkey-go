package main

import (
  "fmt"
  "./sparkey"
)

func main () {
  s := sparkey.New("sparkey_db", sparkey.COMPRESSION_NONE, 1024)

  s.Put("first", "Hello")
  s.Put("second", "Worlb")
  s.Put("third", "Goodbye")
  s.Put("fourth", "EOM")
  s.Delete("third")
  s.Flush()


  s.Get("first")
  // s.GetAll()
  // TODO fix Delete?

  // s.PrettyPrintLog();
  s.PrettyPrintHash();

  fmt.Printf("Sparkey info:\n\n")
  fmt.Printf("  Basename:\t\t%s\n", s.Basename)
  fmt.Printf("  CompressionType:\t\t%d\n", s.CompressionType)
  fmt.Printf("  Size:\t\t%d\n", s.Size())
  fmt.Printf("  LogWriter.CompressionType:\t\t%d\n", s.CompressionType)
  fmt.Printf("  LogWriter.BlockSize:\t\t%d\n", s.BlockSize)
  fmt.Printf("  MaxKeyLen:\t\t%d\n", s.MaxKeyLen)
  fmt.Printf("  MaxValueLen:\t\t%d\n", s.MaxValueLen)

  s.Close()
}
