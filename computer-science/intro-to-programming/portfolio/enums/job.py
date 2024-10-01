from enum import Enum

class Job(Enum):
    MAGE = "Mage"
    WARRIOR = "Warrior"
    ROGUE = "Rogue"
    PRIEST = "Priest"
    HUNTER = "Hunter"
    WARLOCK = "Warlock"
    PALADIN = "Paladin"
    SHAMAN = "Shaman"
    DRUID = "Druid"
    MONK = "Monk"

    def __str__(self):
        return self.value

    def __repr__(self):
        return self.value
