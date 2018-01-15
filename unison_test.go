package unison_test

import (
	"github.com/s1kx/unison"
	"github.com/sirupsen/logrus"
)

func ExampleRun() {
	settings := &unison.Config{
		Commands:   []*unison.Command{},   // No commands added
		EventHooks: []*unison.EventHook{}, // No hooks added
		Services:   []*unison.Service{},   // No services added
	}

	// Start the bot
	err := unison.Run(settings)
	if err != nil {
		logrus.Error(err)
	}
}
