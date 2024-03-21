package conv

import "github.com/bytedance/sonic"

func Trans[T any](from any) (to T, err error) {
	b, err := sonic.Marshal(from)
	if err != nil {
		return to, err
	}
	return Unmarshal[T](b)

}

func Unmarshal[T any](from []byte) (to T, err error) {
	err = sonic.Unmarshal(from, &to)
	return
}

func MustTrans[T any](from any) (to *T) {
	if from == nil {
		return nil
	}

	b, err := sonic.Marshal(from)
	if err != nil {
		return nil
	}
	return MustUnmarshal[T](b)
}

func MustUnmarshal[T any](from []byte) (to *T) {
	if from == nil {
		return nil
	}

	var (
		t   T
		err error
	)

	err = sonic.Unmarshal(from, &t)
	if err != nil {
		return nil
	}
	return &t
}
