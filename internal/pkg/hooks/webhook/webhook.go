package webhook

import (
	"fmt"
	"github.com/xelalexv/dregsy/internal/pkg/hooks"
)

type Webhook struct {
	url string
}

func (w *Webhook) OnSyncFinished(result hooks.SyncResult) {
	b, err := result.MarshalJSON()
	if err != nil {
		panic("ERR")
	}
	fmt.Println(string(b))
	panic("implement me")
}

func NewWebhook(url string) (*Webhook, error) {
	return &Webhook{
		url: url,
	}, nil
}
