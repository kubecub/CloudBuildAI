// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package utils

// IntPtr converts an integer to an *int as a convenience
func IntPtr(i int) *int {
	return &i
}

// Float32Ptr converts a float32 to a *float32 as a convenience
func Float32Ptr(f float32) *float32 {
	return &f
}
