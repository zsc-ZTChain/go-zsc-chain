package db

import (
	"math/big"
	"math/rand"
	"sync"
	"testing"

	"github.com/simplechain-org/go-simplechain/common"
	"github.com/simplechain-org/go-simplechain/core/types"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/stretchr/testify/assert"
)

func setupIndexDB(t *testing.T) *storm.DB {
	rootDB, err := OpenStormDB(nil, "testing-cross-db")
	if err != nil {
		t.Fatal(err)
	}
	return rootDB
}

func generateCtx(n int) []*types.CrossTransactionWithSignatures {
	ctxList := make([]*types.CrossTransactionWithSignatures, n)
	for i := 0; i < n; i++ {
		bigI := big.NewInt(int64(i + 1))
		ctxList[i] = &types.CrossTransactionWithSignatures{
			Data: types.CtxDatas{
				CTxId:            common.BigToHash(bigI),
				TxHash:           common.BigToHash(bigI),
				Value:            big.NewInt(rand.Int63n(1e18)),
				From:             common.BigToAddress(bigI),
				BlockHash:        common.Hash{},
				DestinationId:    bigI,
				DestinationValue: big.NewInt(rand.Int63n(1e18)),
				Input:            bigI.Bytes(),
				V:                nil,
				R:                nil,
				S:                nil,
			},
		}
	}
	return ctxList
}

func TestIndexDB_ReadWrite(t *testing.T) {
	root := setupIndexDB(t)

	ctxList := generateCtx(2)

	testFunction1 := func(t *testing.T, db *indexDB) {
		assert.NoError(t, db.Write(ctxList[0]))
		assert.Error(t, db.Write(ctxList[0]))
		assert.Equal(t, db.total, 1)

		assert.NoError(t, db.Write(ctxList[1]))
		ctx, err := db.Read(ctxList[1].ID())
		assert.NoError(t, err, "")
		assert.Equal(t, ctxList[1], ctx, "")
	}
	// Write without cache
	{
		db := NewIndexDB(big.NewInt(1), root, 0)
		db.db.Drop(&CrossTransactionIndexed{})
		testFunction1(t, db)
	}

	// Write with cache
	{
		db := NewIndexDB(big.NewInt(2), root, 10)
		db.db.Drop(&CrossTransactionIndexed{})
		testFunction1(t, db)
	}

	// Write in restart db
	{
		root.Close()
		root = setupIndexDB(t)
		db := NewIndexDB(big.NewInt(2), root, 10)

		assert.NoError(t, db.Load(), "load occurs an error")
		assert.Equal(t, 2, db.Size())
		db.db.Drop(&CrossTransactionIndexed{})
	}

	// Concurrent Write
	{
		ctxList := generateCtx(40)
		db := NewIndexDB(big.NewInt(3), root, 10)
		db.db.Drop(&CrossTransactionIndexed{})
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			for i := 0; i < 20; i++ {
				assert.NoError(t, db.Write(ctxList[i]))
			}
		}()
		go func() {
			defer wg.Done()
			for i := 20; i < 40; i++ {
				assert.NoError(t, db.Write(ctxList[i]))
			}
		}()
		wg.Wait()
		assert.Equal(t, 40, db.Size())
		count, err := db.db.Count(&CrossTransactionIndexed{})
		assert.NoError(t, err)
		assert.Equal(t, 40, count)
	}

	root.Close()
}

func TestIndexDB_Query(t *testing.T) {
	ctxList := generateCtx(20)
	rootDB := setupIndexDB(t)
	defer rootDB.Close()

	db := NewIndexDB(big.NewInt(1), rootDB, 20)
	db.db.Drop(&CrossTransactionIndexed{})
	for _, ctx := range ctxList {
		assert.NoError(t, db.Write(ctx), "")
	}

	// query without filter
	{
		list := db.QueryByPrice(10, 0)
		assert.Equal(t, 10, len(list))
		for i := 1; i < 10; i++ {
			price1, _ := list[i-1].Price().Float64()
			price2, _ := list[i].Price().Float64()
			assert.LessOrEqual(t, price1, price2)
		}
	}

	// query last 5
	{
		list := db.QueryByPrice(5, 3)
		assert.Equal(t, 5, len(list))
		list = db.QueryByPrice(5, 4)
		assert.Equal(t, 0, len(list))
	}

	// query with status filter after delete
	{
		assert.NoError(t, db.Delete(ctxList[0].ID()))
		list := db.QueryByPrice(100, 0, q.Eq(StatusField, types.CtxStatusFinished))
		assert.Equal(t, 1, len(list))
		list = db.QueryByPrice(100, 0, q.Not(q.Eq(StatusField, types.CtxStatusFinished)))
		assert.Equal(t, 19, len(list))
	}
}
