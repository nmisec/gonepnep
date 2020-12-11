// Example for gonepnep SDK © GoldenXelenium

package main

import (
    "fmt"

    "github.com/nmisec/gonepnep"
)

func main() {
    // Creating a new NeppedAPI session
    session := gonepnep.NewSession("some token")

    // Creating a new SHARP API session
    sharp := gonepnep.NewSHARPSession(*session)

    // Get info about vlfz
    resp, err := sharp.Check(178404926869733376)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(resp)

    // Ban vlfz
    resp, err = sharp.Ban(178404926869733376)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(resp)

    // Unban vlfz
    resp, err = sharp.Unban(178404926869733376)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(resp)

    // Get some "sad" images
    img, err := session.GetImages("sad")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(img.URL)
}
