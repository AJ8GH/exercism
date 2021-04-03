def is_pangram(sentence):
    ALPHABET = 'abcdefghijklmnopqrstuvwxyz'

    characters = ''.join(set(sentence)).lower()

    import re
    re.sub(r'[^a-zA-Z]', '', characters)


    if len(characters) == 26:
        return True

    return False
