package kafka

import (
	"io"
	"io/ioutil"
)

func Send(rw io.ReadWriter, req []byte) ([]byte, error) {
	var err error

	_, err = rw.Write(req)
	if err != nil {
		return nil, err
	}

	res, err := ioutil.ReadAll(rw)
	if err != nil {
		return nil, err
	}

	return res, nil
}