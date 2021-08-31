package main

import (
	"io/ioutil"
	"log"

	d2d2s "github.com/gucio321/d2d2s/pkg/d2s"
	"github.com/gucio321/d2d2s/pkg/d2s/d2senums"
	"github.com/gucio321/d2d2s/pkg/d2s/d2sitems"
	"github.com/gucio321/d2d2s/pkg/d2s/d2sitems/itemdata"
)

func main() {
	d2s := d2d2s.NewCharacter(
		"Example_Char",
		d2senums.CharacterClassBarbarian,
	).SetLevel(27).PushItems(
		d2sitems.NewItem(itemdata.BerRune).
			SetLocation(
				d2senums.LocationStored,
				d2senums.EquippedAnywhere,
				1, 1,
				d2senums.StorageInventory,
			),
	)

	data, err := d2s.Encode()
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("Example_Char.d2s", data, 0o644); err != nil {
		log.Fatal(err)
	}
}
