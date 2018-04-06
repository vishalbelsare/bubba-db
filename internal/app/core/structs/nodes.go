package structs

import (
	"fmt"
	"graph-db/internal/pkg/utils"
	"graph-db/internal/app/core/globals"
)

type Node struct {
	id       int
	isUsed   bool
	isWritten bool
	relationship *Relationship
	property *Property
	label    *Label
}

func (n Node) toBytes() (bs []byte) {
	//todo
	var (
		rel *Relationship
		prop *Property
		label *Label
		relBs, propBs, labelBs []byte
	)
	isUsed := utils.BoolToByteArray(n.isUsed)
	rel = n.GetRelationship()
	if rel != nil {
		relBs = utils.Int32ToByteArray(int32((*rel).id))
	} else {
		relBs = utils.Int32ToByteArray(-1)
	}

	prop = n.GetProperty()
	if prop != nil {
		propBs = utils.Int32ToByteArray(int32((*prop).id))
	} else {
		propBs = utils.Int32ToByteArray(-1)
	}

	label = n.GetLabel()
	if label != nil {
		labelBs = utils.Int32ToByteArray(int32((*label).id))
	} else {
		labelBs = utils.Int32ToByteArray(-1)
	}
	bs = append(isUsed, relBs...)
	bs = append(bs, propBs...)
	bs = append(bs, labelBs...)
	return bs
}

func (n Node) fromBytes(bs []byte) {
	//todo
	var (
		id int32
		rel Relationship
		prop Property
		label Label
	)
	if len(bs) != globals.NodesSize {
		errorMessage := fmt.Sprintf("converter: wrong byte array length. expected array length is 13, actual length is %d", len(bs))
		panic(errorMessage)
	}
	n.isUsed, err = utils.ByteArrayToBool(bs[0:1])
	utils.CheckError(err)
	id, err = utils.ByteArrayToInt32(bs[1:5])
	utils.CheckError(err)
	rel.id = int(id)
	n.relationship = &rel
	id, err = utils.ByteArrayToInt32(bs[5:9])
	utils.CheckError(err)
	prop.id = int(id)
	n.property = &prop
	id, err = utils.ByteArrayToInt32(bs[9:13])
	utils.CheckError(err)
	label.id = int(id)
	n.label = &label
}

func (n Node) GetRelationship() *Relationship {
	//todo
	if n.relationship != nil {
		return n.relationship
	} else if !n.isWritten {
		return nil
	} else {
		offset := n.id * globals.NodesSize
		bs := make([]byte, globals.NodesSize)
		globals.FileHandler.Read(globals.NodesStore, offset, bs)
		relId, err := utils.ByteArrayToInt32(bs[1:5])
		utils.CheckError(err)
		if relId == -1 {
			return nil
		} else {
			var rel Relationship
			rel.id = int(relId)
			n.relationship = &rel
			return n.relationship
		}
	}
}

func (n Node) GetProperty() *Property {
	//todo
	if n.property != nil {
		return n.property
	} else if !n.isWritten {
		return nil
	} else {
		offset := n.id * globals.NodesSize
		bs := make([]byte, globals.NodesSize)
		globals.FileHandler.Read(globals.NodesStore, offset, bs)
		propId, err := utils.ByteArrayToInt32(bs[5:9])
		utils.CheckError(err)
		if propId == -1 {
			return nil
		} else {
			var prop Property
			prop.id = int(propId)
			n.property = &prop
			return n.property
		}
	}
}

func (n Node) GetLabel() *Label {
	//todo
	if n.label != nil {
		return n.label
	} else if !n.isWritten {
		return nil
	} else {
		offset := n.id * globals.NodesSize
		bs := make([]byte, globals.NodesSize)
		globals.FileHandler.Read(globals.NodesStore, offset, bs)
		labelId, err := utils.ByteArrayToInt32(bs[9:13])
		utils.CheckError(err)
		if labelId == -1 {
			return nil
		} else {
			var label Label
			label.id = int(labelId)
			n.label = &label
			return n.label
		}
	}
}

func (n Node) write()  {
	//todo
	offset := globals.NodesSize * n.id
	bs := n.toBytes()
	err := globals.FileHandler.Write(globals.NodesStore, offset, bs)
	utils.CheckError(err)
	n.isWritten = true
}

func (n Node) read() {
	//todo
}

func (n Node) Create() {
	//todo
	id, err := globals.FileHandler.ReadId(globals.NodesId)
	utils.CheckError(err)
	n.id = id
	n.isUsed = true
	n.isWritten = false
	n.write()
}

type Label struct {
	id int
	isUsed bool
	numberOfLabels int
	labelNames [5]LabelTitle
}

type LabelTitle struct {
	id int
	title string
	counter int
}
