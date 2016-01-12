// Copyright 2016 Meng Shi.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Hash algorithm implementation for Times33
package times33

func Hash(s string) int64 {
	var hash int64  = 5381

	for _, c := range s {
		hash = ((hash << 5) + hash) + int64(c);
	}

	return hash
}
