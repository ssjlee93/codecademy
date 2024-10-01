from enum import Enum

class Element(Enum):
    FIRE = "Fire"
    WATER = "Water"
    EARTH = "Earth"
    WIND = "Wind"
    LIGHT = "Light"
    DARK = "Dark"
    NONE = "None"

    def __str__(self):
        return self.value

    def __repr__(self):
        return self.value