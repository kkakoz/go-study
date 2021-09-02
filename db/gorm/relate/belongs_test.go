package relate_test

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"testing"
)

func ExistRecord(err error) (bool, error) {
	if err == nil {
		return true, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return false, err
}
func code1(err error) error {
	if b, err := ExistRecord(err); err != nil {
		return err
	} else if !b {
		fmt.Println("not found")
	} else {
		fmt.Println("exist")
	}
	return nil
}

func HandlerNotFound(err error, f func() error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return f()
	}
	return err
}
func code2(err error) error {
	if err != nil {
		return HandlerNotFound(err, func() error {
			return errors.New("not found")
		})
	}
	fmt.Println("exist")
	return nil
}

func Exist(err error) bool {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
		return false
	}
	return true
}

func TestA(t *testing.T) {
	var err error = gorm.ErrRecordNotFound
	if !Exist(err) {
		fmt.Println("not found")
	}
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println("exist")
	//code1(err)
	//code2(err)
}