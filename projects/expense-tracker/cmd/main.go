package main

import (
	"ferranrt/roadmap-sh/expense-tracker/internal/commands"
)

func main() {
	commands.GetRootCmd().Execute()
}
