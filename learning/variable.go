package main
import (
  "fmt"
  "os"
)

func main() {

    TTLHours = os.Getenv("TTL_HOURS", 48)
    fmt.Println("TTL Hours: ", TTLHours)


}
