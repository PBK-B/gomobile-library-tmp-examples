/*
 * @Author: Bin
 * @Date: 2022-07-20
 * @FilePath: /goMobileDemoApp/src/support.go
 */
package support

import "fmt"

type SupportObserve interface {
	OnCallback(code int, msg string)
}

type Support struct {
	observe SupportObserve
}

func NewSupport() *Support {
	return &Support{}
}

func (s *Support) Echo() *Support {
	fmt.Println("Init Support Library")
	return s
}

func (s *Support) Call(name string) *Support {
	// TODO: Implement function
	if name == "" {
		s.observe.OnCallback(404, "")
		return s
	}
	s.observe.OnCallback(200, "Hello "+name)
	return s
}

func (s *Support) SetObserve(o SupportObserve) *Support {
	s.observe = o
	return s
}
