from rest_framework import serializers
from .sudoku import SudokuInstance

class SudokuSerializer(serializers.Serializer):
    instance = serializers.CharField(max_length=81, min_length=81, allow_blank=False, trim_whitespace=True)

    solution = serializers.CharField(max_length=81, min_length=81, allow_blank=False, trim_whitespace=True)
