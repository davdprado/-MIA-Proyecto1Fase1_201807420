package Estructuras

type Particion struct {
	PartStatus bool
	PartType   byte
	PartFit    byte
	PartStart  int64
	PartSize   int64
	PartName   [16]byte
}
type Disco struct {
	Identificador byte
	Letra         byte
	Path          string
	Particiones   [100]ParticionMontada
}

//Fdisk lleva los datos del comando fdisk
type Fdisk struct {
	Size   int64
	Unit   byte
	Path   string
	Type   byte
	Fit    string
	Delete string
	Name   string
	Add    int64
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
	Mtamano      int64
	Mfecha       [20]byte
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
