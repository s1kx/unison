package state

import "testing"

func TestGetInstance(t *testing.T) {
	instance, err1 := GetInstance()
	instance2, err2 := GetInstance()
	if err1 != nil {
		t.Error(err1.Error())
	} else if err2 != nil {
		t.Error(err2.Error())
	}

	if instance.Info().Data != instance2.Info().Data {
		t.Error("First instance was not similar to second")
	}
}

func TestSetGuildState(t *testing.T) {
	err := SetGuildState("andersfylling", Normal)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetGuildState(t *testing.T) {
	SetGuildState("andersfylling", Normal)
	tt, err := GetGuildState("andersfylling")
	if err != nil {
		t.Error(err.Error())
	}

	if tt != Normal {
		t.Errorf("Guild state was incorrect. Have %d, wants %d", tt, Normal)
	}
}

func TestSetGuildValue(t *testing.T) {
	err := SetGuildValue("andersfylling", "test", []byte("testing"))

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetGuildValue(t *testing.T) {
	expected := "testing"
	SetGuildValue("andersfylling", "test", []byte(expected))
	val, err := GetGuildValue("andersfylling", "test")
	if err != nil {
		t.Error(err.Error())
	}

	if string(val) != expected {
		t.Errorf("Incorrect value. Have %s, wants %s", string(val), expected)
	}
}
