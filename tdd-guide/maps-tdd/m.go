package m

import "errors"

type Dictonary map[string]string

var ErrorNotFound = errors.New("error not found")
var ErrorRepeatDefine = errors.New("error repeat define")
var ErrorUpdateNotFound = errors.New("error update not found")
var ErrorRemoveNotFound = errors.New("error remove not found")

func (dictonary Dictonary) Search(key string) (string, error) {

	value, ok := dictonary[key]

	if !ok {
		return "", ErrorNotFound
	}
	return value, nil
}

func (d Dictonary) Add(key, val string) error {
	_, ok := d[key]
	if !ok {
		d[key] = val
		return nil
	}
	return ErrorRepeatDefine
}

func (d Dictonary) Update(key, val string) error {
	_, ok := d[key]
	if !ok {
		return ErrorUpdateNotFound
	}
	d[key] = val
	return nil
}

func (d Dictonary) Remove(key string) error {
	_, ok := d[key]
	if !ok {
		return ErrorRemoveNotFound
	}
	delete(d, key)
	return nil
}
