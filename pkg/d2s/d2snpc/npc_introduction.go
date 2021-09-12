package d2snpc

import "github.com/gucio321/d2d2s/internal/datautils"

func newIntroduction() *Introduction {
	return &Introduction{
		NPCs: make(map[NPCID]bool),
	}
}

type Introduction struct {
	unknown [3]byte
	NPCs    map[NPCID]bool
}

func (i *Introduction) Decode(data [8]byte) {
	sr := datautils.CreateBitMuncher(data[:], 0)

	for n := NPCUnknown0; n <= NPCUnknown16; n++ {
		i.NPCs[n] = sr.GetBit() == 1
	}
}

type NPCID byte

const (
	NPCUnknown0 NPCID = iota
	NPCGheed
	NPCAkara
	NPCKashya
	NPCUnknown1
	NPCCharsi
	NPCUnknown2
	NPCWarriv
	NPCUnknown3
	NPCDrognan
	NPCFara
	NPCLysander
	NPCGeglash
	// act 2?
	NPCMeshif
	NPCUnknown4
	NPCGreiz
	NPCElzix
	NPCUnknown5
	NPCCain
	NPCUnknown6
	NPCUnknown7
	NPCAsheara
	NPCUnknown8
	NPCAlkor
	NPCOrmus
	NPCUnknown9
	NPCUnknown10
	// probably in 3 act?
	NPCMeshif2 // Meshif
	NPCNatalya
	NPCUnknown11
	NPCAnya
	NPCMalah
	NPCNihlathak
	NPCQualKehk // Qual-Kehk
	// NPCCain2 needs a bit invesitgation. In introductions section, it means
	// that there are 2 acts, when cain has something to talk on start
	NPCCain2 // Cain
	// some of this values may not be present anywhere (end of table)
	NPCUnknown12
	NPCUnknown13
	NPCUnknown14
	NPCUnknown15
	NPCUnknown16
)
