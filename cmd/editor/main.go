package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/AllenDang/giu"

	"github.com/gucio321/d2d2s/pkg/d2s"
	"github.com/gucio321/d2d2s/pkg/d2swidget"
)

func main() {
	d2spath := flag.String("path", "", "D2S file path")
	flag.Parse()

	d2sdata, err := ioutil.ReadFile(*d2spath)
	if err != nil {
		log.Fatalf("error loading data file %v: %v", d2spath, err)
	}

	d2s, err := d2s.Load(d2sdata)
	if err != nil {
		log.Fatalf("error loading d2s file %v: %v", d2spath, err)
	}

	wnd := giu.NewMasterWindow("d2s editor", 640, 480, 0)
	wnd.Run(func() {
		giu.SingleWindow().Layout(
			d2swidget.D2S(d2s),
		)
	})
}
