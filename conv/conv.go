package conv

import "github.com/bytedance/sonic"

func Trans[T any](from any, to *T) error {
	b, err := sonic.Marshal(from)
	if err != nil {
		return err
	}
	err = sonic.Unmarshal(b, to)
	return err
}

func MustTrans[T any](from any) (to *T) {
	if from == nil {
		return nil
	}

	var t T
	b, err := sonic.Marshal(from)
	if err != nil {
		return nil
	}
	err = sonic.Unmarshal(b, &t)
	if err != nil {
		return nil
	}
	return &t
}
