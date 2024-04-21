/*
This file is part of BlockyUtil project. It is licensed under The MIT License.

The MIT License (MIT)
Copyright (c) 2024 BlockyDeer <blockydeer@outlook.com>
Permission is hereby granted, free of charge, to any person obtaining a copy of this
software and associated documentation files (the “Software”), to deal in the Software
without restriction, including without limitation the rights to use, copy, modify,
merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to the following
conditions:
The above copyright notice and this permission notice shall be included in all copies
or substantial portions of the Software.
The Software is provided “as is”, without warranty of any kind, express or implied,
including but not limited to the warranties of merchantability, fitness for a
particular purpose and noninfringement. In no event shall the authors or copyright
holders be liable for any claim, damages or other liability, whether in an action of
contract, tort or otherwise, arising from, out of or in connection with the software
or the use or other dealings in the Software.
*/
package blockyutil

// ensure that the expression will not less than 0
func ZeroFloor(value float64) float64 {
	if value > 0 {
		return value
	} else {
		return 0
	}
}

// ensure that the expression will not less than a specific value
func Rooftop(value float64, top float64) float64 {
	if value < top {
		return top
	} else {
		return value
	}
}

// check if the value is between the minimum and the maximum or it really is the minimum or the maximum
func Between(value float64, minimum float64, maximum float64) bool {
	return value <= maximum && value >= minimum
}
