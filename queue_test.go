package proximityhash

import "testing"

func Test_stringQueueEnqueue_addOneItem_shouldBeInStorageAndSizeIncremented(t *testing.T) {
	queue := newStringQueue()

	queue.enqueue("123")

	if queue.size != 1 {
		t.Fatalf("size should be 1!")
	}
	if queue.storage[0] != "123" {
		t.Fatalf("should contain inserted value!")
	}
}

func Test_stringQueueEnqueue_addThreeItems_shouldBeInStorageAndSizeIncremented(t *testing.T) {
	queue := newStringQueue()

	queue.enqueue("123", "456", "789")

	if queue.size != 3 {
		t.Fatalf("size should be 3!")
	}
	if queue.storage[0] != "123" ||  queue.storage[1] != "456" || queue.storage[2] != "789"{
		t.Fatalf("should contain inserted values!")
	}
}

func Test_stringQueueDequeue_queueIsEmpty_okShouldBeFalse(t *testing.T){
	queue := newStringQueue()

	_, ok := queue.dequeue()

	if ok != false {
		t.FailNow()
	}
}

func Test_stringQueueDequeue_queueHasItems_shouldReturnFirstItemAndDecrementSize(t *testing.T){
	queue := newStringQueue()
	queue.enqueue("123", "345")

	res, ok := queue.dequeue()

	if ok != true || res != "123" || queue.size != 1 {
		t.FailNow()
	}
}

func Test_stringQueueIsEmpty_queueIsEmpty_true(t *testing.T){
	queue := newStringQueue()

	res := queue.isEmpty()

	if res != true {
		t.FailNow()
	}
}

func Test_stringQueueIsEmpty_queueIsNotEmpty_false(t *testing.T){
	queue := newStringQueue()
	queue.enqueue("123")

	res := queue.isEmpty()

	if res != false {
		t.FailNow()
	}
}

func Test_newStringQueue_shouldReturnInitializedEmptyQueue(t *testing.T){
	res := newStringQueue()

	if len(res.storage) != 0 || res.size != 0 {
		t.FailNow()
	}
}