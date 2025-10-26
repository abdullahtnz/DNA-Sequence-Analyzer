# DNA-Sequence-Analyzer
Short Summary: The code analyze DNA sequence, which is obtained via reading .FASTA file or given plain DNA sequence. It generates an output which shows the result of analyze including the complementary DNA strand, The Amino Acid names in codons using The Universal Genetic Code, GC content, and The percentage of all nucleotide bases (A,T,C,G)

# Detailed Describtion
Running this code, a person provided with 2 options - .FASTA file or plain DNA sequence. After providing the programme with one of those, this Go code start analyzing and evaluating that DNA Sequence. Because all the DNA structure contains 4 Nucleotide base elements - A (Adenine), T(Thymine), C (Cytosine), G(Guanine) - bioinformatics engineers must find all the bases with their percentage in the DNA. Most importantly, to have information about which type of living organism the DNA belong to, it is a must to determine GC content of DNA. This code does all these important steps for you in just some milliseconds.

After these, in order to find the Amino Acids in the DNA, this software seperates the DNA sequence into 3-bases codons (ATG ~ start codon, UGA ~ end codon etc.), it evaluates each of these codons using The Universal Genetic Code Chart. The output includes the name of Amino Acids in DNA codons, their number, and percentage.

# Important Features
- The most important feature of this code is the runtime. Due to the Golang's fast runtime, analyzing and evaluating the DNA sequence and printing all the outputs takes very little time.
- The flexibility. The program provide the user with the flexibility of choosing how to give ( input ) the DNA Sequence, whether it will be a .FASTA file ( which is mostly used for advanced bioinformatic concepts ) or the plain DNA sequence text ( ATGCAT..., for basic studies ).

