package unison

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func testCommandAction(ctx *Context, msg *discordgo.Message, request string) error {
	return nil
}

func TestCommandParser(t *testing.T) {
	var cmd = &Command{
		Name:   "test",
		Usage:  "unit testing kek",
		Action: testCommandAction,
		Flags: &struct {
			Input  string
			Output []string
		}{},
	}

	cmd.buildCommand()

	userInput := "test lol kek \"testing test lol2\" --input=\"test\""
	cmd.parseInput(userInput)
}

func TestCommandParserWithContentCheck(t *testing.T) {
	var args struct {
		Input  string
		Output []string
	}
	var cmd = &Command{
		Name:   "test",
		Usage:  "unit testing kek",
		Action: testCommandAction,
		Flags:  &args,
	}

	cmd.buildCommand()

	userInput := "test lol kek \"testing test lol2\" --input=\"test\""
	cmd.parseInput(userInput)

	if args.Input != "test" {
		t.Errorf("Incorrectly parsed user input. Expected `test`, got %s\n", args.Input)
	}
}

func TestCommandParserWithoutFlags(t *testing.T) {
	var cmd = &Command{
		Name:   "test",
		Usage:  "unit testing kek",
		Action: testCommandAction,
	}

	cmd.buildCommand()

	userInput := "test lol kek \"testing test lol2\" --input=\"test\""
	cmd.parseInput(userInput)
}
