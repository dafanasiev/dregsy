package hooks

import "encoding/json"

type SyncResult interface {
	EncodeJSON(js *json.Encoder) error
}

type Hook interface {
	OnSyncFinished(result SyncResult)
}
