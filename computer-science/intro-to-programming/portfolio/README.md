# README

## RPG character creation
things I need :  
* character
* job
* classes
* [optional]weapons
* [optional]magic

## notes

### python imports
python uses dot notation for importing files.  
```python
from enums.job import job
```
this imports from `job.py` in `enums` as `job` variable

## __pycache__

__pycache__ is a directory automatically created by Python when you run a Python script. It stores compiled bytecode versions of your Python modules. This compiled bytecode is a more efficient representation of your Python code, which can speed up the execution of your scripts, especially for larger projects.

thought:  
seems like python created these pycache since I tried to import python classes / files.  
in order to import them and understand them as python files, python created __pycache__