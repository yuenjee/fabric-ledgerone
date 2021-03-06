/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package committer

import (
	"sync/atomic"
	"testing"

	"github.com/ledgerone/fabric-ledgerone/common/configtx/test"
	"github.com/ledgerone/fabric-ledgerone/common/ledger/testutil"
	"github.com/ledgerone/fabric-ledgerone/common/tools/configtxgen/encoder"
	"github.com/ledgerone/fabric-ledgerone/common/tools/configtxgen/localconfig"
	"github.com/ledgerone/fabric-ledgerone/common/util"
	ledger2 "github.com/ledgerone/fabric-ledgerone/core/ledger"
	"github.com/ledgerone/fabric-ledgerone/core/ledger/ledgermgmt"
	"github.com/ledgerone/fabric-ledgerone/protos/common"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestKVLedgerBlockStorage(t *testing.T) {
	viper.Set("peer.fileSystemPath", "/tmp/fabric-ledgerone/committertest")
	ledgermgmt.InitializeTestEnv()
	defer ledgermgmt.CleanupTestEnv()
	gb, _ := test.MakeGenesisBlock("TestLedger")
	gbHash := gb.Header.Hash()
	ledger, err := ledgermgmt.CreateLedger(gb)
	assert.NoError(t, err, "Error while creating ledger: %s", err)
	defer ledger.Close()

	committer := NewLedgerCommitter(ledger)
	height, err := committer.LedgerHeight()
	assert.Equal(t, uint64(1), height)
	assert.NoError(t, err)

	bcInfo, _ := ledger.GetBlockchainInfo()
	testutil.AssertEquals(t, bcInfo, &common.BlockchainInfo{
		Height: 1, CurrentBlockHash: gbHash, PreviousBlockHash: nil})

	txid := util.GenerateUUID()
	simulator, _ := ledger.NewTxSimulator(txid)
	simulator.SetState("ns1", "key1", []byte("value1"))
	simulator.SetState("ns1", "key2", []byte("value2"))
	simulator.SetState("ns1", "key3", []byte("value3"))
	simulator.Done()

	simRes, _ := simulator.GetTxSimulationResults()
	simResBytes, _ := simRes.GetPubSimulationBytes()
	block1 := testutil.ConstructBlock(t, 1, gbHash, [][]byte{simResBytes}, true)

	err = committer.CommitWithPvtData(&ledger2.BlockAndPvtData{
		Block: block1,
	})
	assert.NoError(t, err)

	height, err = committer.LedgerHeight()
	assert.Equal(t, uint64(2), height)
	assert.NoError(t, err)

	blocks := committer.GetBlocks([]uint64{0})
	assert.Equal(t, 1, len(blocks))
	assert.NoError(t, err)

	bcInfo, _ = ledger.GetBlockchainInfo()
	block1Hash := block1.Header.Hash()
	testutil.AssertEquals(t, bcInfo, &common.BlockchainInfo{
		Height: 2, CurrentBlockHash: block1Hash, PreviousBlockHash: gbHash})
}

func TestNewLedgerCommitterReactive(t *testing.T) {
	viper.Set("peer.fileSystemPath", "/tmp/fabric-ledgerone/committertest")
	chainID := "TestLedger"

	ledgermgmt.InitializeTestEnv()
	defer ledgermgmt.CleanupTestEnv()
	gb, _ := test.MakeGenesisBlock(chainID)

	ledger, err := ledgermgmt.CreateLedger(gb)
	assert.NoError(t, err, "Error while creating ledger: %s", err)
	defer ledger.Close()

	var configArrived int32
	committer := NewLedgerCommitterReactive(ledger, func(_ *common.Block) error {
		atomic.AddInt32(&configArrived, 1)
		return nil
	})

	height, err := committer.LedgerHeight()
	assert.Equal(t, uint64(1), height)
	assert.NoError(t, err)

	profile := localconfig.Load(localconfig.SampleSingleMSPSoloProfile)
	block := encoder.New(profile).GenesisBlockForChannel(chainID)

	committer.CommitWithPvtData(&ledger2.BlockAndPvtData{
		Block: block,
	})
	assert.Equal(t, int32(1), atomic.LoadInt32(&configArrived))
}
