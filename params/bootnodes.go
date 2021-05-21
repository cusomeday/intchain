// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the INT Chain main network.
var MainnetBootnodes = []string{

	"enode://7aa041c21b8818c6996a505e4281c8b55c6a953989797ab5528c6e5a2f679a74a785e93120612fcbd457b1ad33de91716118584b6f40f541d62f3db64bedd8ec@127.0.0.1:8550",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// INT Chain test network.
var TestnetBootnodes = []string{
	"enode://517aeaa0d9ed9b0f336200d31a536322b05684b59fa7a35d1b594f9ae229d277040873ed324518222abb1a270426f77237f7ae9f1253f4ae9a86b0cf848fed5b@101.32.74.50:8550",    // Titans
	"enode://9381a60ca2c65cb649c56d84789bb1164cafd766a754ed484e414f5b02bc6a622bc8d360043fb0e776471119229fff74d358f57415730d7f3d62d3c0155666d5@129.226.128.55:8551",  // Oceanus
	"enode://b82aa7354bcf98cfe4eb07da7bb39b22ae27618165ab9136a3a6525c3ab7b87114fc7743b4dff312acc1eb31185a22d2bb097fd0692f6bf3ccaa98b605b140d0@129.226.63.13:8551",   // Iapetus
	"enode://99a74bc83ea34e24bc7b6fa6dcf8e1f533febb296c7000f4b9f1f873365ba6514435d35fe0182014bac31b3c0f4e21398f398cfe089c6b284f4c5fe2a5e5acd3@170.106.160.155:8551", // Mnemosyne
}
