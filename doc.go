// Copyright (c) 2013-2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

/*
Package base58 provides base58-check encoding.
The alphabet is modifyiable for

Base58 Usage

 To decode a base58 string:

 rawData,err := base58.Decode("3VNr6P")
 if err != nil {
    fmt.Println(err.Error())
	}
 fmt.println(string(rawData))
 >> abcd

 Similarly, to encode the same data:

 encodedData := base58.Encode([]byte("abcd"))
 fmt.println(encodedData)
 >> 3VNr6P
*/
package base58
