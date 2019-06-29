package in

import "errors"

type GetPing struct {
	Text string `form:"text"`
}

func (p *GetPing) Validate() error {
	if p.Text == "" {
		return errors.New("text is require")
	}

	return nil
}
