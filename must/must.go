/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package must

func Get[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
