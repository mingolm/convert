package convert

import "context"

type Csv struct {
	*Config
}

func (cv *Csv) Run(ctx context.Context) error {
	return nil
}
