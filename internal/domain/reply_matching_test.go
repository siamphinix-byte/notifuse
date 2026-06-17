package domain

import "testing"

func TestRFCMessageIDValueAndHeader(t *testing.T) {
	cases := []struct {
		name       string
		messageID  string
		from       string
		wantValue  string
		wantHeader string
	}{
		{"normal", "abc-123", "hello@example.com", "abc-123@example.com", "<abc-123@example.com>"},
		{"subdomain", "id1", "news@mg.example.com", "id1@mg.example.com", "<id1@mg.example.com>"},
		{"no at sign falls back", "id2", "not-an-address", "id2@notifuse.local", "<id2@notifuse.local>"},
		{"trailing at falls back", "id3", "x@", "id3@notifuse.local", "<id3@notifuse.local>"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := RFCMessageIDValue(tc.messageID, tc.from); got != tc.wantValue {
				t.Errorf("RFCMessageIDValue = %q, want %q", got, tc.wantValue)
			}
			if got := BuildRFCMessageID(tc.messageID, tc.from); got != tc.wantHeader {
				t.Errorf("BuildRFCMessageID = %q, want %q", got, tc.wantHeader)
			}
		})
	}
}

func TestProviderSetsOwnMessageID(t *testing.T) {
	setOwn := []EmailProviderKind{EmailProviderKindMailgun}
	other := []EmailProviderKind{
		EmailProviderKindSendGrid, EmailProviderKindSparkPost,
		EmailProviderKindMailjet, EmailProviderKindPostmark,
		EmailProviderKindSES, EmailProviderKindSMTP,
	}
	for _, k := range setOwn {
		if !ProviderSetsOwnMessageID(k) {
			t.Errorf("ProviderSetsOwnMessageID(%s) = false, want true", k)
		}
	}
	for _, k := range other {
		if ProviderSetsOwnMessageID(k) {
			t.Errorf("ProviderSetsOwnMessageID(%s) = true, want false (not yet wired)", k)
		}
	}
}
