FROM golang:1.16 AS builder
WORKDIR /sudoku/generator/
COPY ./sudoku/generator/generator.go ./ 
RUN go build -o generator.so -buildmode=c-shared generator.go

FROM python:3.8
ENV PYTHONUNBUFFERED 1

WORKDIR /app/backend
COPY --from=builder /sudoku/generator/generator.so ./
COPY --from=builder /sudoku/generator/generator.h ./
COPY requirements.txt .
RUN pip3 install -r requirements.txt
COPY . .

CMD ["python", "./manage.py", "runserver", "0.0.0.0:8000"]
EXPOSE 8000
