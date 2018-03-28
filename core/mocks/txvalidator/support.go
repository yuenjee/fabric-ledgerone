/*
Copyright IBM Corp. 2017 All Rights Reserved.

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

package support

import (
	"github.com/ledgerone/fabric-ledgerone/common/channelconfig"
	mockpolicies "github.com/ledgerone/fabric-ledgerone/common/mocks/policies"
	"github.com/ledgerone/fabric-ledgerone/common/policies"
	"github.com/ledgerone/fabric-ledgerone/common/resourcesconfig"
	"github.com/ledgerone/fabric-ledgerone/core/ledger"
	"github.com/ledgerone/fabric-ledgerone/msp"
	"github.com/ledgerone/fabric-ledgerone/protos/common"
)

type Support struct {
	LedgerVal     ledger.PeerLedger
	MSPManagerVal msp.MSPManager
	ApplyVal      error
	ACVal         channelconfig.ApplicationCapabilities
}

func (ms *Support) ChaincodeByName(chainname, ccname string) (resourcesconfig.ChaincodeDefinition, bool) {
	return nil, false
}

func (ms *Support) Capabilities() channelconfig.ApplicationCapabilities {
	return ms.ACVal
}

// Ledger returns LedgerVal
func (ms *Support) Ledger() ledger.PeerLedger {
	return ms.LedgerVal
}

// MSPManager returns MSPManagerVal
func (ms *Support) MSPManager() msp.MSPManager {
	return ms.MSPManagerVal
}

// Apply returns ApplyVal
func (ms *Support) Apply(configtx *common.ConfigEnvelope) error {
	return ms.ApplyVal
}

func (ms *Support) PolicyManager() policies.Manager {
	return &mockpolicies.Manager{}
}

func (cs *Support) GetMSPIDs(cid string) []string {
	return []string{"DEFAULT"}
}
