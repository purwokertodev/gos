package gos

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	number = `0-9`
	alpha  = `A-Za-z`
)

type gos struct {
	//Allowed input to your string
	input string
	//Allow empty to your string
	allowEmpty bool
	//Allow Number to your string
	allowNumber bool
	//Allow Alpabetical to your string
	allowAlpha bool
}

func New(allowedInput string, allowEmpty, allowNumber, allowAlpha bool) gos {
	return gos{
		allowedInput,
		allowEmpty,
		allowNumber,
		allowAlpha,
	}
}

func (g gos) Validate(input string) error {
	var inputs []string
	var allowedInput string

	for _, v := range g.input {
		inputs = append(inputs, string(v))
	}
	allowedInput = strings.Join(inputs, `\`)

	var ai string

	if g.allowEmpty && !g.allowNumber && !g.allowAlpha {
		ai = fmt.Sprintf(`^[%s ]*$`, allowedInput)
	} else if g.allowEmpty && g.allowNumber && g.allowAlpha {
		ai = fmt.Sprintf(`^[%s%s%s ]*$`, alpha, number, allowedInput)
	} else if g.allowEmpty && !g.allowNumber && g.allowAlpha {
		ai = fmt.Sprintf(`^[%s%s ]*$`, alpha, allowedInput)
	} else if g.allowEmpty && g.allowNumber && !g.allowAlpha {
		ai = fmt.Sprintf(`^[%s%s ]*$`, number, allowedInput)
	} else if !g.allowEmpty && g.allowNumber && g.allowAlpha {
		ai = fmt.Sprintf(`^[%s%s%s ]+$`, alpha, number, allowedInput)
	} else if !g.allowEmpty && !g.allowNumber && g.allowAlpha {
		ai = fmt.Sprintf(`^[%s%s ]+$`, alpha, allowedInput)
	} else if !g.allowEmpty && g.allowNumber && !g.allowAlpha {
		ai = fmt.Sprintf(`^[%s%s ]+$`, number, allowedInput)
	} else {
		ai = fmt.Sprintf(`^[%s%s ]+$`, number, allowedInput)
	}
	check := regexp.MustCompile(ai).MatchString
	if !check(input) {
		return errors.New("Invalid Input")
	}
	return nil

}
