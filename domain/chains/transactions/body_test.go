package transactions

import (
	"fmt"
	"testing"

	"github.com/steve-care-software/libs/cryptography/hash"
)

func TestBody_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pAddressHash, err := hashAdapter.FromBytes([]byte("this is an address"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pReferenceHash, err := hashAdapter.FromBytes([]byte("this is a reference"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	fees := uint(34)
	body, err := NewBodyBuilder().Create().WithAddress(*pAddressHash).WithReference(*pReferenceHash).WithFees(fees).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !body.Address().Compare(*pAddressHash) {
		t.Errorf("the returned address is invalid")
		return
	}

	if !body.Reference().Compare(*pReferenceHash) {
		t.Errorf("the returned reference is invalid")
		return
	}

	if body.Fees() != fees {
		t.Errorf("the fees was expected to be %d, %d returned", fees, body.Fees())
		return
	}

	pHash, err := hashAdapter.FromMultiBytes([][]byte{
		pAddressHash.Bytes(),
		pReferenceHash.Bytes(),
		hash.Hash(fmt.Sprintf("%d", fees)),
	})

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !body.Hash().Compare(*pHash) {
		t.Errorf("the returned hash is invalid")
		return
	}
}

func TestBody_withoutAddress_ReturnsError(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pReferenceHash, err := hashAdapter.FromBytes([]byte("this is a reference"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	fees := uint(34)
	_, err = NewBodyBuilder().Create().WithReference(*pReferenceHash).WithFees(fees).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestBody_withoutReference_ReturnsError(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pAddressHash, err := hashAdapter.FromBytes([]byte("this is an address"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	fees := uint(34)
	_, err = NewBodyBuilder().Create().WithAddress(*pAddressHash).WithFees(fees).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestBody_withoutFees_ReturnsError(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pAddressHash, err := hashAdapter.FromBytes([]byte("this is an address"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pReferenceHash, err := hashAdapter.FromBytes([]byte("this is a reference"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	_, err = NewBodyBuilder().Create().WithAddress(*pAddressHash).WithReference(*pReferenceHash).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
