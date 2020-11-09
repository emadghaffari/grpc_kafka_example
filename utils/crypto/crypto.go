package crypto

import (
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"github.com/xdg/scram"
)

// SHA256 var
var SHA256 scram.HashGeneratorFcn = func() hash.Hash { return sha256.New() }

// SHA512 var
var SHA512 scram.HashGeneratorFcn = func() hash.Hash { return sha512.New() }

// XDGSCRAMClient struct
type XDGSCRAMClient struct {
	*scram.Client
	*scram.ClientConversation
	scram.HashGeneratorFcn
}

// Begin func
func (x *XDGSCRAMClient) Begin(userName, password, authzID string) (err error) {
	x.Client, err = x.HashGeneratorFcn.NewClient(userName, password, authzID)
	if err != nil {
		return err
	}
	x.ClientConversation = x.Client.NewConversation()
	return nil
}

// Step func
func (x *XDGSCRAMClient) Step(challenge string) (response string, err error) {
	response, err = x.ClientConversation.Step(challenge)
	return
}

// Done func
func (x *XDGSCRAMClient) Done() bool {
	return x.ClientConversation.Done()
}
