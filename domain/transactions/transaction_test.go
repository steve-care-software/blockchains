package transactions

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/libs/cryptography/hash"
)

func TestTransaction_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pubKey := []byte("this is an address")
	pAddressHash, err := hashAdapter.FromBytes(pubKey)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	body := NewBodyForTests(*pAddressHash)
	signature := []byte("this is a signature")
	trx, err := NewTransactionBuilder(func(pubKey []byte, signature []byte, hash hash.Hash) bool {
		return true
	}).Create().WithBody(body).WithSignature(signature).WithPublicKey(pubKey).Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !body.Hash().Compare(trx.Body().Hash()) {
		t.Errorf("the returned body is invalid")
		return
	}

	if bytes.Compare(pubKey, trx.PublicKey()) != 0 {
		t.Errorf("the returned publicKey is invalid")
		return
	}

	if bytes.Compare(signature, trx.Signature()) != 0 {
		t.Errorf("the returned signature is invalid")
		return
	}

	pHash, err := hashAdapter.FromMultiBytes([][]byte{
		body.Hash().Bytes(),
		signature,
		pubKey,
	})

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !trx.Hash().Compare(*pHash) {
		t.Errorf("the returned hash is invalid")
		return
	}
}

func TestTransaction_signatureCannotBeVerified_ReturnsError(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pubKey := []byte("this is an address")
	pAddressHash, err := hashAdapter.FromBytes(pubKey)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	body := NewBodyForTests(*pAddressHash)
	signature := []byte("this is a signature")
	_, err = NewTransactionBuilder(func(pubKey []byte, signature []byte, hash hash.Hash) bool {
		return false
	}).Create().WithBody(body).WithSignature(signature).WithPublicKey(pubKey).Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestTransaction_hashedPublicKeyDoNotMatchAddress_ReturnsError(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pAddressHash, err := hashAdapter.FromBytes([]byte("this is an address"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	body := NewBodyForTests(*pAddressHash)
	pubKey := []byte("this pubKey do not match the address")
	signature := []byte("this is a signature")
	_, err = NewTransactionBuilder(func(pubKey []byte, signature []byte, hash hash.Hash) bool {
		return true
	}).Create().WithBody(body).WithSignature(signature).WithPublicKey(pubKey).Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestTransaction_withEmptyPublicKey_ReturnsError(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pAddressHash, err := hashAdapter.FromBytes([]byte("this is an address"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pubKey := []byte{}
	body := NewBodyForTests(*pAddressHash)
	signature := []byte("this is a signature")
	_, err = NewTransactionBuilder(func(pubKey []byte, signature []byte, hash hash.Hash) bool {
		return true
	}).Create().WithBody(body).WithSignature(signature).WithPublicKey(pubKey).Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestTransaction_without_ReturnsError(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pAddressHash, err := hashAdapter.FromBytes([]byte("this is an address"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	body := NewBodyForTests(*pAddressHash)
	signature := []byte("this is a signature")
	_, err = NewTransactionBuilder(func(pubKey []byte, signature []byte, hash hash.Hash) bool {
		return true
	}).Create().WithBody(body).WithSignature(signature).Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestTransaction_withEmptySignature_ReturnsError(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pubKey := []byte("this is an address")
	pAddressHash, err := hashAdapter.FromBytes([]byte(pubKey))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	signature := []byte{}
	body := NewBodyForTests(*pAddressHash)
	_, err = NewTransactionBuilder(func(pubKey []byte, signature []byte, hash hash.Hash) bool {
		return true
	}).Create().WithBody(body).WithSignature(signature).WithPublicKey(pubKey).Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestTransaction_withoutSignature_ReturnsError(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pubKey := []byte("this is an address")
	pAddressHash, err := hashAdapter.FromBytes([]byte(pubKey))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	body := NewBodyForTests(*pAddressHash)
	_, err = NewTransactionBuilder(func(pubKey []byte, signature []byte, hash hash.Hash) bool {
		return true
	}).Create().WithBody(body).WithPublicKey(pubKey).Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestTransaction_withoutBody_ReturnsError(t *testing.T) {
	pubKey := []byte("this is an address")
	signature := []byte("this is a signature")
	_, err := NewTransactionBuilder(func(pubKey []byte, signature []byte, hash hash.Hash) bool {
		return true
	}).Create().WithSignature(signature).WithPublicKey(pubKey).Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
