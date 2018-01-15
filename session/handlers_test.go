package session

import "testing"

func TestIdentifiesAsCommand(t *testing.T) {
	prefix := "!"

	inputs := []string{
		"fjg jfdgg!",
		"  !",
		"!!!",
		"_!sdfs",
		"t!es",
	}

	for _, str := range inputs {
		b, res := identifiesAsCommand(str, prefix)
		if b {
			t.Errorf("String identified as command, when expected not to. Content{%s}, res{%s}, prefix{%s}", str, res, prefix)
		}
	}

	commands := []string{
		"!test",
		"!!!test",
		"!gfg",
		"!_!_!",
		"!  dffs",
	}

	for _, str := range commands {
		b, res := identifiesAsCommand(str, prefix)
		if !b {
			t.Errorf("String did not identified as command, when expected to. Content{%s}, result{%s}, prefix{%s}", str, res, prefix)
		}
	}

}

func TestRemovePrefix(t *testing.T) {
	if removePrefix("! ", "!") != "" {
		t.Errorf("Issue removing prefix. Got `%s`, wants `%s`", removePrefix("! ", "!"), "")
	}
}
