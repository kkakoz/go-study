package gomock_study_test

import (
	. "github.com/golang/mock/gomock"
	"learn-go/study_test/gomock_study"
	"testing"
)

func TestRepo(t *testing.T) {
	ctrl := NewController(t)
	defer ctrl.Finish()
	mockRepo := gomock_study.NewMockRepository(ctrl)
	mockRepo.EXPECT().Create("a", "b").Return(nil)
	//mockRepo.EXPECT().Create("c", "d").Return(errors.New("test err"))
	err := mockRepo.Create("a", "b")
	if err != nil {
		t.Fatal(err)
	}
	//err = mockRepo.Create("c", "d")
	//if err != nil {
	//	t.Fatal(err)
	//}
	mockRepo.EXPECT().GetInt(Any()).Return(100)
	i := mockRepo.GetInt("a")
	t.Log("i = ", i)

}
