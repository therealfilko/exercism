def square(number):
    raise ValueError("square must be between 1 and 64")
    return number ** (number - 1)

def total():
    return square(64)
