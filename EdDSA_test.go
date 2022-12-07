package naive_ed25519

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestSignifyVerify(t *testing.T) {
	tests := []struct {
		secret    string
		public    string
		msg       string
		signature string
		status    bool
	}{
		{
			"4ccd089b28ff96da9db6c346ec114e0f5b8a319f35aba624da8cf6ed4fb8a6fb",
			"3d4017c3e843895a92b70aa74d1b7ebc9c982ccf2ec4968cc0cd55f12af4660c",
			"72",
			"92a009a9f0d4cab8720e820b5f642540a2b27b5416503f8fb3762223ebdb69da085ac1e43e15996e458f3613d0f11d8c387b2eaeb4302aeeb00d291612bb0c00",
			true,
		},
		// TODO: add more from sign.input
	}

	for i, _ := range tests {
		secret, _ := hex.DecodeString(tests[i].secret)
		public, _ := hex.DecodeString(tests[i].public)
		msg, _ := hex.DecodeString(tests[i].msg)
		signature, _ := hex.DecodeString(tests[i].signature)
		status := tests[i].status
		if !reflect.DeepEqual(Signify(secret, msg), signature) {
			t.Errorf("Invalid signature %d", i)
		}
		if Verify(public, msg, signature) != status {
			t.Errorf("Invalid verification %d", i)
		}
	}
}
