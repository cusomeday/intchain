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
	"enode://fe6f20248eeafb54d15a6a2e54bcd7fe1cfed72e15fa57bc315196d14c59337393bd3b8e65916c1e2c966ba960ca71f69cc4ca1a7a86a7a6d65275d96ac1afa9@159.138.3.248:8552",  // INT Chain
	"enode://8c3985f82dd41b07d3c31caf5088a538a578e9a911cd65600dc7fe3ba909122e0c99843df2c74dfdadb6241ca16a2197c7646cc678461df7b32f77bdb4837cc3@119.8.53.187:8551",   // Oceanus
	"enode://4dc46b217ab8dbea7d313e7dfa4692f9fdf03f1398e666748ba4978b7fb56227922266eb92374be79cdaeb23e46e42c12cd42a96b999c4901f5b862fd92752b8@94.74.127.193:8551",  // Hyperion
	"enode://9fecf5479d6a66140a2127143265d0d38289b94348beb952ee239e3e4ee45de1bf40b4f5029c155b81d820cdfd401068ec8d9a3bf19d327f86760ba13609cf6a@182.160.9.154:8551",  // Rhea
	"enode://b26fbc2c328580b4bfb54b1369e61dc7e7a71819915883c5344169feca40e1379d94158a8c27a5519d1fc13c1e22b0b6d1359136ae2a682c9bd7beb171969875@182.160.13.153:8551", // Phoebe

	"enode://9260fcf0ff3996607d9e1ea7b222107ad3eaf657b1c8aa760690bf971f7e4151c6271a2c87412b31312c76b3d0fb285947d3832e6cb972db14a5be9a45cb7873@162.55.177.167:8550", // Community
	"enode://d71a17716ee2a41c270e01ab63c28863f09fddd1e1e5abbc0d5cdc623e927529b502118812a902ebc868bd76570c851ecb50620cda19333a4ca92a9c32f348c9@49.12.192.182:8550",  // Community
	"enode://ef1979cc634da020a132f6c0ea53cb112f7c1eabc6d285e977bd75eb79dee89693c30686dc52575510e6b1e664427e25ea2c2d36d4630371ae1d250f8ec110db@78.47.204.74:8550",   // Community
	"enode://78ccae42c2baa9e22ea6ce064efff0f74eee138fab44533b199712e03972f8660d9598912fb4d6aca99b615fe6a25fdaa5fce9aa5ffd0cca70f33c2d2242e300@159.138.47.28:8551",  // Community
	"enode://d0dfe66b564f9e5215f915fff362fd4899464e26aaebbf3c3b55bb5790a7b60a8d2434ebd309a9d10b44fac88ff86e72b1701d45a71185398206a9762376e554@119.8.38.127:8551",   // Community
	"enode://585913f8318ccfc6359b63f2e08e01f0097c4fb84098e27b533bd5ca6d9425700238c11599e8662b41f567155cfa14705d7f8ba68f3d3a809083da1a50baeec7@54.238.212.246:8550", // Community
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// INT Chain test network.
var TestnetBootnodes = []string{
	"enode://a12a394da7887a8fa99cf3347cd372b08502e3dd04f6e464967b5a3cdebf7234f2dd4e5fc9720961cf03224920b8b3c4c25ae23ba5d593345ea41b1b7e575b7d@101.32.74.50:8550", // INT Chain
	//"enode://dd7e1ead96067f73ed13c5a7a0e6d1cf7a92644a225478acda42c48750aff021263740481b5d1d72275fe19e4d61d064bbb3293adba0f0845d013f7ece63c75d@101.32.74.50:8551",   // Titans
	"enode://4dd8d8352900f44cc6bfd9d9c3f2d9c8d80d0b6a64b97019dc9e3ca905a4ca6474b84745e5001fc08c5b64910baeb233fcd54bda4f6f855438d7c762494977c0@129.226.59.148:8551", // Oceanus
	"enode://964f2c8d04bcc48e971bb24870035db8cfc7218bfdb4d97849421b47d9cc207fff1b89876b4ff137156fb52a262254a570396a4c06a756d2c68a8f637eaa5031@129.226.128.55:8551", // Iapetus
	"enode://65e60686f0c091b30a701b3c8e02db1aab75ad67e7d319c679a616cad38ce2dad5abb8829cd0d6041a9541387078eed12b88f5d0eb5168456a182ad2d0183e19@129.226.63.13:8551",  // Mnemosyne
}
