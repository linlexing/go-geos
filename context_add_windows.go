package geos
// #include <stdlib.h>
// #include "go-geos.h"
import "C"

func (c *Context) NewGeomFromWKB(wkb []byte) (*Geom, error) {
	c.Lock()
	defer c.Unlock()
	c.err = nil
	if c.wkbReader == nil {
		c.wkbReader = C.GEOSWKBReader_create_r(c.handle)
	}
	wkbCBuf := C.CBytes(wkb)
	defer C.free(wkbCBuf)
	return c.newGeom(C.GEOSWKBReader_read_r(c.handle, c.wkbReader, (*C.uchar)(wkbCBuf), C.ulonglong(C.ulong(len(wkb)))), nil), c.err
}
