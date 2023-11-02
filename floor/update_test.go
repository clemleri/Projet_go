package floor 

import (
	"testing"
	"reflect"
)

func Test_size(t *testing.T) {
	content := updateFromFileFloor(camXPos, camYPos) 
	if len(content)!= Global.NumTileY {
		t.Fail()
	}
	for i:=0;i<len(content);i++ {
		if content[i]!= Global.NumTileX {
			t.Fail()
		}
	}
}

func Test_out(t *testing.T) {
	content := updateFromFileFloor(camXPos, camYPos) 
	for i:=0; i<len()
}

