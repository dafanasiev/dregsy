package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xelalexv/dregsy/internal/pkg/hooks"
)

type Webhook struct {
	url string
}

func (w *Webhook) OnSyncFinished(result hooks.SyncResult) {
	//TODO implement me
	b := bytes.NewBuffer(nil)
	js := json.NewEncoder(b)
	js.SetIndent("", "  ")
	js.SetEscapeHTML(false)
	err := result.EncodeJSON(js)
	if err != nil {
		panic("ERR")
	}
	fmt.Println(b.String())
	panic("implement me")
}

func NewWebhook(url string) (*Webhook, error) {
	return &Webhook{
		url: url,
	}, nil
}
