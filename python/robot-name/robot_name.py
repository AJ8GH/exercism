import random
import string

class Robot:
    names = []

    def __init__(self):
        self.generate_name()

    def generate_name(self):
        random.seed()
        name = ''
        while len(name) < 2:
            name += random.choice(string.ascii_uppercase)
        while len(name) < 5:
            name += str(random.randint(0, 9))
        self.validate_uniqueness(name)

    def validate_uniqueness(self, name):
        if name not in self.names:
            self.names.append(name)
            self.name = name
        else:
            self.generate_name()

    def reset(self):
        self.generate_name()
