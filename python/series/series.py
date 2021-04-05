def slices(series, length):
    validate(series, length)

    number_of_slices = len(series) - length + 1
    slices = []
    start = 0

    while start < number_of_slices:
        end = start + length
        slices.append(series[slice(start, end)])
        start += 1

    return slices

def validate(series, length):
    if len(series) <= 0 or length <= 0 or length > len(series):
        raise ValueError('.+')
