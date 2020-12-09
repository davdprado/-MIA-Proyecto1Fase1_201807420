package Estructuras

type Particion struct {
	PartStatus bool
	PartType   byte
	PartFit    byte
	PartStart  int64
	PartSize   int64
	PartName   [16]byte
}

type ParticionLogica struct {
	PartStatus bool
	PartType   byte
	PartFit    byte
	PartStart  int64
	PartSize   int64
	PartName   [16]byte
}

type Mbr struct {
	Mtamaño      int64
	Mfecha       [15]byte
	Mdisco       int64
	MParticiones [4]Particion
	Mlibre       int64
}
type EBR struct {
	PartStatus      bool
	PartFit         byte
	PartStart       int64
	PartSize        int64
	PartNext        int64
	PartName        [16]byte
	ParticionLogica ParticionLogica
}

type Mkdisk struct {
	Size       int64
	Path, Name string
	Unit       byte
}

type ParticionMontada struct {
	Identificador string
	Path, Name    string
	Estado        bool
	EstadoEscrito bool
}
