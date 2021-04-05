def classify(number):
    assert_positive_int(number)

    factors = sum_of_factors(number)

    if factors == number:
        return 'perfect'
    elif factors > number:
        return 'abundant'
    return 'deficient'

def sum_of_factors(number):
    factors = 0
    for num in range(1, number):
        if number % num == 0:
            factors += num
    return factors

def assert_positive_int(number):
    if number <= 0 or isinstance(number, int) == False:
        raise ValueError('.+')
