//
// Copyright (c) 2017. Marcus Brummer.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gopixl

/*
#cgo LDFLAGS: -lpixl
#include <pixl/pixl.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

// Image todo
type Image struct {
	Width  uint32
	Height uint32
	handle *C.CPixlImage
}

// LoadImage loads the image at the specified path
func LoadImage(path string) *Image {
	cpath := C.CString(path)
	img := C.pixl_load_image(cpath)
	C.free(unsafe.Pointer(cpath))

	return &Image{
		Width:  uint32(img.width),
		Height: uint32(img.height),
		handle: img,
	}
}

// Save saves image
func (img *Image) Save(path string, quality int) {
	cpath := C.CString(path)
	C.pixl_save_image(img.handle, cpath, C.int(quality))
	C.free(unsafe.Pointer(cpath))
}

// Free frees resources
func (img *Image) Free() {
	C.pixl_destroy_image(img.handle)
}
