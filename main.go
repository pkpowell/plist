package plist

import (
	"fmt"
	"strings"

	plist "github.com/micromdm/plist"
	hplist "howett.net/plist"
)

const data = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>typekey</key>
	<string>A</string>
	<key>typeAkey</key>
	<string>VALUE-A</string>
</dict>
</plist>`

type TypeDecider struct {
	ActualType interface{} `plist:"-"`
}

type TypeA struct {
	TypeAKey string `plist:"typeAkey"`
}

type TypeB struct {
	TypeBKey string `plist:"typeBkey"`
}

func (t *TypeDecider) UnmarshalPlist(f func(interface{}) error) error {
	// stub struct for decoding a single key to tell which
	// specific type we should umarshal into
	typeKey := &struct {
		TypeKey string `plist:"typekey"`
	}{}

	err := f(typeKey)
	if err != nil {
		return err
	}

	// switch using the decoded value to determine the correct type
	switch typeKey.TypeKey {
	case "A":
		t.ActualType = new(TypeA)
	case "B":
		t.ActualType = new(TypeB)
	case "":
		return fmt.Errorf("empty typekey (or wrong input data)")
	default:
		return fmt.Errorf("unknown typekey: %s", typeKey.TypeKey)
	}

	// decode into the actual type
	return f(t.ActualType)
}

// ExampleUnmarshaler demonstrates using structs that use the Unmarshaler interface.
func MUnmarshaler() {
	decider := new(TypeDecider)
	err := plist.Unmarshal([]byte(data), decider)
	if err != nil {
		fmt.Println(err)
		return
	}

	typeA, ok := decider.ActualType.(*TypeA)
	if !ok {
		fmt.Println("actual type is not TypeA")
		return
	}

	fmt.Println(typeA.TypeAKey)
	// Output: VALUE-A
}

func HUnmarshaler() {
	var d TypeDecider
	err := hplist.NewDecoder(strings.NewReader(data)).Decode(&d)
	if err != nil {
		fmt.Println("getDSCL Unmarshal error", err.Error())
		// return nil
	}

}
