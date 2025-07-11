// cmd/blockchainnftregistrytechx/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftregistrytechx/internal/blockchainnftregistrytechx"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftregistrytechx.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
