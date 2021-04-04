import re

CODONS = {
    'Methionine': ['AUG'],
    'Phenylalanine': ['UUU', 'UUC'],
    'Leucine': ['UUA', 'UUG'],
    'Serine': ['UCU', 'UCC', 'UCA', 'UCG'],
    'Tyrosine': ['UAU', 'UAC'],
    'Cysteine': ['UGU', 'UGC'],
    'Tryptophan': ['UGG'],
    'STOP': ['UAA', 'UAG', 'UGA']
}

def proteins(strand):
    sequence = re.findall('...', strand)
    polypeptide = []

    for codon in sequence:
        if codon in CODONS['STOP']:
            break
        for protein in CODONS:
            if codon in CODONS[protein]:
                polypeptide.append(protein)

    return polypeptide
