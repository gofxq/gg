package main

import "github.com/bytedance/sonic"

func Trans[T any](from any, to *T) error {
	b, err := sonic.Marshal(from)
	if err != nil {
		return err
	}
	err = sonic.Unmarshal(b, to)
	return err
}
