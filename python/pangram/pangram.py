def is_pangram(sentence):
    ALPHABET = 'abcdefghijklmnopqrstuvwxyz'

    sentence = sentence.lower()
    results = []

    for char in sentence:
        if char not in ALPHABET or char in results:
            continue
        results.append(char)

    return len(results) == 26
