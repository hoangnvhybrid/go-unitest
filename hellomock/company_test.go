package hellomock

import (
	//"github.com/golang/mock/gomock"
	//gomocket "github.com/selvatico/go-mocket"
	"testing"
)

func TestCompany_Meeting(t *testing.T) {
	person := NewPerson("Nguyen Hoang")
	company := NewCompany(person)
	t.Log(company.Meeting("hoa"))
}
//func TestCompany_Meeting2(t *testing.T) {
//	//clt := gomock.NewController(t)
//}
