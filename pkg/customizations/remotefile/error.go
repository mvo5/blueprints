package remotefile

// XXX: copied from "osbuild-composer/internal/worker/clienterrors"
// but we also have a copy of this in "images/internal/worker/clienterrors"
type ClientErrorCode int

const (
	ErrorRemoteFileResolution ClientErrorCode = 36
)

type Error struct {
	ID      ClientErrorCode `json:"id"`
	Reason  string          `json:"reason"`
	Details interface{}     `json:"details,omitempty"`
}
