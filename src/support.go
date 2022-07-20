/*
 * @Author: Bin
 * @Date: 2022-07-20
 * @FilePath: /goMobileDemoApp/src/support.go
 */
package support

import "fmt"

type Support struct {
}

func NewSupport() *Support {
	return &Support{}
}

func (s *Support) Echo() *Support {
	fmt.Println("Init Support Library")
	return s
}
