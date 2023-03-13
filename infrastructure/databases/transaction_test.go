package databases

import (
	"os"
	"testing"

	"github.com/steve-care-software/blockchains/applications"
	"github.com/steve-care-software/blockchains/domain/chains/transactions"
	"github.com/steve-care-software/databases/infrastructure/files"
	"github.com/steve-care-software/libs/cryptography/hash"
)

func TestTransaction_repositoryAndService_Success(t *testing.T) {
	dirPath := "./test_files"
	dstExtension := "destination"
	bckExtension := "backup"
	readChunkSize := uint(1000000)
	defer func() {
		os.RemoveAll(dirPath)
	}()

	fileName := "my_file.db"
	dbApp := files.NewApplication(dirPath, dstExtension, bckExtension, readChunkSize)
	err := dbApp.New(fileName)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pContext, err := dbApp.Open(fileName)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	defer dbApp.Close(*pContext)

	// create repository:
	repositoryBuilder := NewTransactionRepositoryBuilder(dbApp, transactions.NewTransactionBuilder(func(pubKey []byte, signature []byte, hash hash.Hash) bool {
		return true
	}))

	repository, err := repositoryBuilder.Create().WithContext(*pContext).WithKind(applications.KindTransaction).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// create service:
	service, err := NewTransactionServiceBuilder(dbApp, repositoryBuilder).Create().WithContext(*pContext).WithKind(applications.KindTransaction).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// insert transactions:
	firstTrx := transactions.NewTransactionForTests([]byte("this is a firstpubKey"))
	secondTrx := transactions.NewTransactionForTests([]byte("this is a second pubKey"))
	err = service.InsertList([]transactions.Transaction{
		firstTrx,
		secondTrx,
	})

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// commit:
	err = dbApp.Commit(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// retrieve the list of transactions:
	retTrxHashesList := repository.List()
	if len(retTrxHashesList) != 2 {
		t.Errorf("%d hashes were expected, %d returned", 2, len(retTrxHashesList))
		return
	}

	// retrieve the first transaction:
	retTrx, err := repository.Retrieve(firstTrx.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !firstTrx.Hash().Compare(retTrx.Hash()) {
		t.Errorf("the returned transaction is invalid")
		return
	}

	// retrieve the list of transactions:
	retTransactions, err := repository.RetrieveList(retTrxHashesList)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retTrxList := retTransactions.List()
	if len(retTrxList) != 2 {
		t.Errorf("%d hashes were expected, %d returned", 2, len(retTrxList))
		return
	}

	// erase a transaction:
	err = service.Erase(firstTrx.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// commit:
	err = dbApp.Commit(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retListAfterErase := repository.List()
	if len(retListAfterErase) != 1 {
		t.Errorf("%d hashes were expected, %d returned", 1, len(retListAfterErase))
		return
	}

	// erase all:
	err = service.EraseAll([]hash.Hash{
		secondTrx.Hash(),
	})

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// commit:
	err = dbApp.Commit(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retListAfterSecondErase := repository.List()
	if len(retListAfterSecondErase) != 0 {
		t.Errorf("%d hashes were expected, %d returned", 0, len(retListAfterSecondErase))
		return
	}
}

func TestTransaction_repositoryAndService_insertsDuplicate_ReturnsError(t *testing.T) {
	dirPath := "./test_files"
	dstExtension := "destination"
	bckExtension := "backup"
	readChunkSize := uint(1000000)
	defer func() {
		os.RemoveAll(dirPath)
	}()

	fileName := "my_file.db"
	dbApp := files.NewApplication(dirPath, dstExtension, bckExtension, readChunkSize)
	err := dbApp.New(fileName)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pContext, err := dbApp.Open(fileName)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	defer dbApp.Close(*pContext)

	// create repository:
	repositoryBuilder := NewTransactionRepositoryBuilder(dbApp, transactions.NewTransactionBuilder(func(pubKey []byte, signature []byte, hash hash.Hash) bool {
		return true
	}))

	// create service:
	service, err := NewTransactionServiceBuilder(dbApp, repositoryBuilder).Create().WithContext(*pContext).WithKind(applications.KindTransaction).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// insert transactions:
	firstTrx := transactions.NewTransactionForTests([]byte("this is a firstpubKey"))
	err = service.Insert(firstTrx)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// commit:
	err = dbApp.Commit(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = service.Insert(firstTrx)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
