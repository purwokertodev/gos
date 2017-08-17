**Gos**

[![Build Status](https://travis-ci.org/wuriyanto48/gos.svg?branch=master)](https://travis-ci.org/wuriyanto48/gos)

# Usage

- **Download to your project:**

```shell
go get github.com/wuriyanto48/gos
```

```go
package main

import(
	"fmt"
	"github.com/wuriyanto48/gos"
)

func main() {
	g := gos.New(`\%$()}&"':`, true, true, true)
	input := `&:wuriyanto0 alex`
	err := g.Validate(input)
	if err != nil {
		fmt.Println("DO SOMETHING WITH ERROR")	
	}
	
}
```

# What is that ?

```go
g := gos.New(`\%$()}&"':`, true, true, true)
```

- First Parameter `\%$()}&"':` : Allowed input to your string
- Second Parameter `boolean` : Allow empty to your string, whic mean your input can be empty ("")
- Third Parameter `boolean` : Allow Number to your string
- Fourth Parameter `boolean` : Allow Alpabetical to your string



