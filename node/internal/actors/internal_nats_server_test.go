package actors

import (
	"log/slog"
	"slices"
	"testing"

	"github.com/nats-io/nkeys"
	"github.com/synadia-io/nex/models"
)

func TestGenerateConfig(t *testing.T) {
	hostUser, err := nkeys.CreateUser()
	if err != nil {
		t.Fatal(err)
	}

	ns := &InternalNatsServer{
		hostUser: hostUser,
		logger:   slog.Default(),
		nodeOptions: models.NodeOptions{
			AgentOptions: []models.AgentOptions{
				{
					Name: "native",
				},
				{
					Name: "javascript",
				},
			},
		},
	}

	creds, err := ns.buildAgentCredentials()
	if err != nil {
		t.Fatal(err)
	}
	ns.creds = creds

	if len(creds) != 2 {
		t.Errorf("should have built creds of length 2; got %d", len(creds))
	}

	opts, err := ns.generateConfig()
	if err != nil {
		t.Errorf("failed to generate server config; %s", err)
	}
	if opts == nil {
		t.Error("failed to generate server config")
		return
	}
	if opts.Accounts == nil {
		t.Error("failed to generate server config")
		return
	}

	accounts := make([]string, 0)
	for _, acct := range opts.Accounts {
		accounts = append(accounts, acct.Name)
	}

	if !slices.Contains(accounts, "nexhost") {
		t.Errorf("server accounts did not contain nexthost; %s", accounts)
	}

	if !slices.Contains(accounts, "native") {
		t.Errorf("server accounts did not contain native; %s", accounts)
	}

	if !slices.Contains(accounts, "javascript") {
		t.Errorf("server accounts did not contain javascript; %s", accounts)
	}
}
