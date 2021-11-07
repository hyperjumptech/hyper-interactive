# Hyper-Interactive

Hyper-Interaction is a very small tiny library to help
making a command line app become interactive. It allowes user to
fill in qustionaire, making configuration or just a simple
information input.

```shell
go get -u github.com/hyperjumptech/hyper-interactive
```

## Simple text input

```go
import (
	"fmt"
	interactive "github.com/hyperjumptech/hyper-interactive"
)

...
name := interactive.Ask("Whats your name ?", "Bruce Wayne", false)
fmt.Println("Hi " + name)
...
```

The command line will look like

```shell
$ interact
Whats your name ? [default "Bruce Wayne"] : Ferdinand
Hi Ferdinand
```

## Default Answer

If you see `[default __]`, simply hit `enter` will use the default value as your answer.

## Confirming your Answer

All question type have a confirmation flag in their argument.
For example

```go
// the last argument is a confirmation flag
name := interactive.Ask("Whats your name ?", "Bruce Wayne", true) 
fmt.Println("Hi " + name)
```

This will result

```shell
$ interact
Whats your name ? [default "Bruce Wayne"] : Ferdinand
"Ferdinand", are you sure? (y/n/Y/N) [default : Y] ? Y
Hi Ferdinand
```

## You will get result

The logic in all of those ask function will make sure the
user fill in the correct answer.

- User are ensured to choose valid option
- User are ensured to specify valid number
- User are ensured to specify valid time format
- etc

## Asking for Options

```go
options := []string {
    "One","Two","Three",
    "Four","Five","Six",
    "Seven","Eight","Nine",
    "Ten",
}
choosen := interactive.Select("Please choose", options, 1,1, true)
fmt.Printf("You choose number %d\n", choosen)
```

This will be shown as

```shell
$ interact
Please choose :
(1) One    (3) Three  (5) Five   (7) Seven  (9) Nine                         
(2) Two    (4) Four   (6) Six    (8) Eight  (10) Ten                         
Choose from number above [default : (1) One] ? 7
(5) Five - Are you sure ?  (y/n/Y/N) [default : Y] ? Y
You choose number 7
```

## Supported Questions

- `func Select(question string, options []string, startFrom, defaultOption int, confirm bool) int`
- `func AskNumber(question string, from, to, def int, confirm bool) int`
- `func AskTime(question string, def time.Time, confirm bool) time.Time`. The answer must follow format `2006-01-02 15:04:05 -0700`
- `func Ask(question, defaultAnswer string, confirm bool) string`
- `func Confirm(question string, def bool) bool`
- more to come...