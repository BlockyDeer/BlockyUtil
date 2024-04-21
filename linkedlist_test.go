package blockyutil_test

import (
	"testing"

	blockyutil "github.com/BlockyDeer/BlockyUtil"
)

func TestList(t *testing.T) {
	list := blockyutil.CreateLinkedList[int32]()

	list.PushBack(1)
	list.PushBack(2111)
	list.PushBack(11230)
	list.PushBack(10231)

	//iterFront := list.Front()
	//t.Logf("data1: %d", *(iterFront.GetData()))

	list.Foreach(func(iterator blockyutil.ListIterator[int32], data *int32) {
		t.Logf("%d\n", *data)
	})
}
