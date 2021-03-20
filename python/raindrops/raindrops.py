def convert(number):
    result = ''

    pling = number % 3 == 0
    plang = number % 5 == 0
    plong = number % 7 == 0

    if pling:
        result += 'Pling'
    if plang:
        result += 'Plang'
    if plong:
        result += 'Plong'
    if not (pling or plang or plong):
        result = f'{number}'

    return result
