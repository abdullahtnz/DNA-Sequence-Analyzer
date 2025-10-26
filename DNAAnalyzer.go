package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FASTAReader(filepath string) (map[string]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fmt.Println(" ")

	sequences := make(map[string]string)
	var currentName string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, ">") {
			currentName = line[1:]
			sequences[currentName] = ""
		} else {
			sequences[currentName] += line
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sequences, nil
}

func dnaAnalyzer(structure string) {
	var complementary string = ""

	for i := 0; i < len(structure); i++ {
		switch structure[i] {

		case 'T':
			complementary += "A"

		case 'A':
			complementary += "T"

		case 'C':
			complementary += "G"

		case 'G':
			complementary += "C"
		}
	}

	var GCcounter int

	GCcounter = strings.Count(structure, "C") + strings.Count(structure, "G")
	GCpercentage := (float64(GCcounter) / float64(len(structure))) * 100

	A := strings.Count(structure, "A")
	T := strings.Count(structure, "T")
	C := strings.Count(structure, "C")
	G := strings.Count(structure, "G")

	fmt.Printf("The number of Adenine base (A): %d\n", A)
	fmt.Printf("The number of Thymine base (T): %d\n", T)
	fmt.Printf("The number of Cytosine base (C): %d\n", C)
	fmt.Printf("The number of Guanine base (G): %d\n", G)
	fmt.Printf("GC content of DNA: %.2f%%\n", GCpercentage)

	var codonTable = map[string]string{
		"ATG": "Met",
		"TGG": "Trp",
		"TAA": "*",
		"TAG": "*",
		"TGA": "*",
		"TTT": "Phe",
		"TTC": "Phe",
		"TTA": "Leu",
		"TTG": "Leu",
		"TCT": "Ser",
		"TCC": "Ser",
		"TCA": "Ser",
		"TCG": "Ser",
		"TAT": "Tyr",
		"TAC": "Tyr",
		"TGT": "Cys",
		"TGC": "Cys",
		"CCT": "Pro",
		"CCC": "Pro",
		"CCA": "Pro",
		"CCG": "Pro",
		"CAT": "His",
		"CAC": "His",
		"CAA": "Gln",
		"CAG": "Gln",
		"CGT": "Arg",
		"CGC": "Arg",
		"CGA": "Arg",
		"CGG": "Arg",
		"ATT": "Ile",
		"ATC": "Ile",
		"ATA": "Ile",
		"ACT": "Thr",
		"ACC": "Thr",
		"ACA": "Thr",
		"ACG": "Thr",
		"AAT": "Asn",
		"AAC": "Asn",
		"AAA": "Lys",
		"AAG": "Lys",
		"AGT": "Ser",
		"AGC": "Ser",
		"AGA": "Arg",
		"AGG": "Arg",
		"GTT": "Val",
		"GTC": "Val",
		"GTA": "Val",
		"GTG": "Val",
		"GCT": "Ala",
		"GCC": "Ala",
		"GCA": "Ala",
		"GCG": "Ala",
		"GAT": "Asp",
		"GAC": "Asp",
		"GAA": "Glu",
		"GAG": "Glu",
		"GGT": "Gly",
		"GGC": "Gly",
		"GGA": "Gly",
		"GGG": "Gly",
	}

	aaCount := make(map[string]int)

	for i := 0; i+2 < len(structure); i += 3 {
		codon := structure[i : i+3]

		if aa, exists := codonTable[codon]; exists {
			aaCount[aa]++
		} else {
			aaCount["X"]++
		}
	}
	fmt.Println(" ")
	fmt.Println("=== The number of Amino Acids and their contents ===")
	fmt.Println(" ")
	total := 0
	for _, count := range aaCount {
		total += count
	}

	for aa, count := range aaCount {
		percentage := float64(count) / float64(total) * 100
		fmt.Printf("%s: %d (%.2f%%)\n", aa, count, percentage)
	}

	fmt.Println(" ")
	fmt.Println("______________________________________________________________________________ ")
	fmt.Println(" ")

}

func main() {

	var filepath string
	var seq string
	var plainornot bool

	fmt.Print("Would you like to give .FASTA file path ? ( true or false )")
	fmt.Scan(&plainornot)

	if plainornot {
		fmt.Print("Enter FASTA File path: ")
		fmt.Scan(&filepath)

		seqMap, err := FASTAReader(filepath)
		if err != nil {
			fmt.Println("Error reading FASTA:", err)
			return
		}

		for name, sequence := range seqMap {
			fmt.Printf("\n=== Sequence: %s ===\n", name)
			dnaAnalyzer(sequence)
		}
	} else {
		fmt.Print("Enter plain DNA sequence: ")
		fmt.Scan(&seq)
		dnaAnalyzer(seq)
	}

}
