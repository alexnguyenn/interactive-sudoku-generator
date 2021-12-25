from django.shortcuts import render
from rest_framework.views import APIView
from .serializers import SudokuSerializer
from rest_framework.response import Response
from .sudoku import SudokuInstance

# Create your views here.
class SudokuAPI(APIView):

    def get(self, request, formate=None, **kwargs): 
        generated_instance = SudokuInstance()
        serializer = SudokuSerializer(generated_instance)
        return Response(serializer.data)
