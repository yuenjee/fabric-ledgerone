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

package bddtests

import (
	"fmt"

	"github.com/ledgerone/fabric-ledgerone/core/comm"
	"google.golang.org/grpc"
)

// NewGrpcClient return a new GRPC client connection for the specified peer address
func NewGrpcClient(peerAddress string) (*grpc.ClientConn, error) {
	var tmpConn *grpc.ClientConn
	var err error
	tmpConn, err = comm.NewClientConnectionWithAddress(peerAddress, true, false, nil)
	if err != nil {
		fmt.Printf("error connection to server at host:port = %s\n", peerAddress)
	}
	return tmpConn, err

}
