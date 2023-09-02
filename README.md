# AmariGo
## AmariGo is a Go library for the AmariBot API.

### Installation
```bash
go get github.com/sleepyyui/amari-go
```

### Usage Example
```golang
package main

import (
    "fmt"
    "github.com/sleepyyui/amari-go"
)

func main() {
    bot := amarigo.New("token")
    user, err := bot.GetGuildMember("guildId", "userId")
    if err != nil {
        panic(err)
    }
    fmt.Println(user.Username)
}
```

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributors
- [SleepyYui](https://sleepyyui.com/)

