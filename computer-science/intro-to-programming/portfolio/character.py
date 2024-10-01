class Character:
    
    def __init__(self, name, element, job):
        self.name = name
        self.element = element
        self.job = job
    
    def __str__(self):
        return f"{self.name} is a {self.job} who uses {self.element}."
    
    def __repr__(self) -> str:
        return (f"{self.name} is a {self.job} who specializes in {self.element}.")
        