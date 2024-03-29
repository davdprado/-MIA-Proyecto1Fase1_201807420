package Estructuras

type Particion struct {
	PartStatus byte
	PartType   [2]byte
	PartFit    [2]byte
	PartStart  int64
	PartSize   int64
	PartName   [16]byte
}
type Disco struct {
	Identificador  int
	Path, Fecha    string
	TamañoD, Asign int64
	Fit            string
	Tp             int64
	IDs            int
	Particiones    []ParticionMontada
}

//Fdisk lleva los datos del comando fdisk

type Mbr struct {
	Mbit         int64
	Mtamano      int64
	Mfecha       [20]byte
	MdiscoA      int64
	Mfit         [2]byte
	MParticiones [4]Particion
}

type ParticionMontada struct {
	Identificador           string
	Name, Tipo, Fit, DiscoR string
	Estado                  byte
	EstadoEscrito           bool
	Porcentaje              float32
	Tamaño                  int64
	BitStar                 int64
}
