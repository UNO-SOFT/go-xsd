package main

import (
	"encoding/xml"

	"github.com/UNO-SOFT/go-xsd/xsd-makepkg/tests"

	"github.com/UNO-SOFT/go-xsd/util/dev/go"

	kml "github.com/UNO-SOFT/go-xsd-pkg/schemas.opengis.net/kml/2.2.0/ogckml22.xsd_go"
)

type KmlDoc struct {
	XMLName xml.Name `xml:"http://www.opengis.net/kml/2.2 kml"`
	kml.TKmlType
}

func main() {
	var (
		dirBasePath  = udevgo.GopathSrcGithub("UNO-SOFT", "go-xsd", "xsd-makepkg", "tests", "xsd-test-kml")
		makeEmptyDoc = func() interface{} { return &KmlDoc{} }
	)
	tests.TestViaRemarshal(dirBasePath, makeEmptyDoc)
}
