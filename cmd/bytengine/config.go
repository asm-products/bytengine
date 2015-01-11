package main

import (
	_ "github.com/asm-products/bytengine/auth/mongo"
	_ "github.com/asm-products/bytengine/bytestore/diskv"
	_ "github.com/asm-products/bytengine/bytestore/mongo"
	_ "github.com/asm-products/bytengine/cmdhandler/base"
	_ "github.com/asm-products/bytengine/datafilter/builtin"
	_ "github.com/asm-products/bytengine/filesystem/mongo"
	_ "github.com/asm-products/bytengine/parser/base"
	_ "github.com/asm-products/bytengine/statestore/redis"
)
