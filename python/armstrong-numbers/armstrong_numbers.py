def is_armstrong_number(number):
    digits = [int(digit) for digit in str(number)]
    squares = list(map(lambda num: num ** len(digits), digits))
    return sum(squares) == number
