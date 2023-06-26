package structs

import (
	"container/list"

	"YJS-GO/utils"
)

type IContent interface {
	Copy() IContent
	Splice(offset int) IContent
	MergeWith(right IContent)
	GetContent() list.List
	GetLength() int
	Countable() bool
}

type IContentExt interface {
	IContent
	Write(encoder utils.IUpdateEncoder, offset int)
	Gc(store utils.StructStore)
	Delete(transaction utils.Transaction)
	Integrate(transaction utils.Transaction, item Item)
	GetRef() int
}
