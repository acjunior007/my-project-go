package main

import "errors"

func Process() error {
	err1 := errors.New("database connection failed")
	err2 := errors.New("dache service not responding")

	return errors.Join(err1, err2)
}
