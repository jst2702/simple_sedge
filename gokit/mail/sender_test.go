package mail

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	kitcfg "simplesedge.com/gokit/config"
)

type EmailConfig struct {
	EmailSenderName     string `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress  string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword string `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

func GetEmailConfig(path string) *EmailConfig {
	config := &EmailConfig{}
	err := kitcfg.LoadConfig(path, "app", "env", config)
	if err != nil {
		log.Fatal("cannot loading config ", err)
	}
	return config
}

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config := GetEmailConfig(".")
	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	<p>this is a test message from <a href="https://github.com/jst2702/simple_sedge">simple-sedge</a></p>
	`
	to := []string{"jst2702@gmail.com"}
	attachFiles := []string{"../../README.md"}
	err := sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
