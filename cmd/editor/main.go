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
			giu.Button("Save").OnClick(func() {
				d2sdata, err := d2s.Encode()
				if err != nil {
					log.Fatalf("error encoding d2s file: %v", err)
				}

				if err := ioutil.WriteFile(*d2spath, d2sdata, 0o644); err != nil {
					log.Fatalf("error writting d2s file: %v", err)
				}
			}),
		)
	})
}
