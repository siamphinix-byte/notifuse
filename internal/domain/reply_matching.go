package domain

import "strings"

// RFCMessageIDValue returns the canonical message-id value WITHOUT angle brackets
// ("messageID@domain", domain taken from the From address). This is what we store
// as message_history.smtp_message_id and what we match against a reply's parsed
// In-Reply-To (also bracket-stripped), so both sides use the identical form.
func RFCMessageIDValue(messageID, fromAddress string) string {
	domain := "notifuse.local"
	if at := strings.LastIndex(fromAddress, "@"); at >= 0 && at+1 < len(fromAddress) {
		domain = fromAddress[at+1:]
	}
	return messageID + "@" + domain
}

// BuildRFCMessageID returns the RFC 5322 Message-ID header value ("<value>") for
// providers that let us set it (set_own). The recipient's reply echoes this in
// In-Reply-To; after bracket-stripping it equals RFCMessageIDValue (the stored
// smtp_message_id), enabling the match.
func BuildRFCMessageID(messageID, fromAddress string) string {
	return "<" + RFCMessageIDValue(messageID, fromAddress) + ">"
}

// ProviderSetsOwnMessageID reports whether Notifuse sets the recipient-visible RFC
// Message-ID itself for this provider (the "set_own" strategy). For these, the
// stored smtp_message_id is BuildRFCMessageID(...). Capture providers (e.g. Mailjet)
// and sender-match providers (SendGrid, SparkPost) are handled separately.
//
// Only providers whose send path actually sets the header are listed here; others
// are added as their send services are updated.
func ProviderSetsOwnMessageID(kind EmailProviderKind) bool {
	switch kind {
	case EmailProviderKindMailgun:
		return true
	default:
		return false
	}
}
