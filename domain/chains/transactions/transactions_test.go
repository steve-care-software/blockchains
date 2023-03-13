package transactions

import (
	"reflect"
	"testing"
)

func TestTransactions_withMultiple_Success(t *testing.T) {
	list := []Transaction{
		NewTransactionForTests([]byte("this is a first publicKey")),
		NewTransactionForTests([]byte("this is a second publicKey")),
		NewTransactionForTests([]byte("this is a third publicKey")),
	}

	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retList := ins.List()
	if !reflect.DeepEqual(list, retList) {
		t.Errorf("the returned list is invalid")
		return
	}
}

func TestTransactions_withSingle_Success(t *testing.T) {
	list := []Transaction{
		NewTransactionForTests([]byte("this is a first publicKey")),
	}

	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retList := ins.List()
	if !reflect.DeepEqual(list, retList) {
		t.Errorf("the returned list is invalid")
		return
	}
}

func TestTransactions_withEmptyList_ReturnsError(t *testing.T) {
	list := []Transaction{}
	_, err := NewBuilder().Create().WithList(list).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestTransactions_withoutList_ReturnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
