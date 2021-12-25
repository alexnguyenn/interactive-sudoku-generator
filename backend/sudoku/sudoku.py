from ctypes import *

class SudokuInstance(object):
    
    def __init__(self):
        lib = cdll.LoadLibrary("../backend/sudoku/generator/generator.so")

        class GoSlice(Structure):
            _fields_ = [("data", POINTER(c_void_p)), ("len", c_longlong), ("cap", c_longlong)]

        lib.generatePuzzleToPython.argtypes = [GoSlice, GoSlice] 
        lib.generatePuzzleToPython.restypes = None
        instance = GoSlice((c_void_p * 81)(), 81, 81) 
        solution = GoSlice((c_void_p * 81)(), 81, 81) 
        lib.generatePuzzleToPython(instance, solution)

        instance_arr = [0 if not instance.data[i] else instance.data[i] for i in range(instance.len)]
        solution_arr = [0 if not solution.data[i] else solution.data[i] for i in range(solution.len)]
        
        self.instance = ''.join(map(str, instance_arr))
        self.solution = ''.join(map(str, solution_arr))

    
    # For debugging 
    def __str__(self):      
        return '(' + self.instance + ', ' + self.solution + ')'
