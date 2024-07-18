FROM python:3.12-alpine

RUN apk add --update curl

WORKDIR /app

COPY . .

RUN pip install --no-cache-dir -r requirements.txt

CMD ["python", "./main.py"]