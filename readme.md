# GoQuery

```go
import (
    . "github.com/StompRocket/GoQuery/goquery"
)

func main() {
    someDiv, err := Q(".some-div")
    if err != nil { panic(err) }

    someDiv.ClassList.Add("another-class")
    println(someDiv.ClassList.Value)
}
```