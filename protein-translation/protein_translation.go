package protein

import "errors"

var (
	ErrStop        = errors.New("invalid stop")
	ErrInvalidBase = errors.New("invalid base")
)

func FromRNA(rna string) (lst []string, err error) {
	for i := 0; i < len(rna); i += 3 {
		protein, err1 := FromCodon(rna[i : i+3])
		if protein == "STOP" {
			return
		}
		if err1 != nil {
			err = err1
			return
		}
		lst = append(lst, protein)
	}

	return
}

func FromCodon(codon string) (protein string, err error) {
	switch codon {
	case "AUG":
		protein = "Methionine"
	case "UUU", "UUC":
		protein = "Phenylalanine"
	case "UUA", "UUG":
		protein = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		protein = "Serine"
	case "UAU", "UAC":
		protein = "Tyrosine"
	case "UGU", "UGC":
		protein = "Cysteine"
	case "UGG":
		protein = "Tryptophan"

	case "UAA", "UAG", "UGA":
		protein = "STOP"
		err = ErrStop
	default:
		err = ErrInvalidBase
	}

	return
}
