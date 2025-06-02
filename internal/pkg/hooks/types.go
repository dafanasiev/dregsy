package hooks

type SyncResult interface {
	MarshalJSON() ([]byte, error)
}

type Hook interface {
	OnSyncFinished(result SyncResult)
}
