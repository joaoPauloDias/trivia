FROM python:3.8-slim

WORKDIR /app

COPY runner.py .

RUN pip install requests

CMD ["python3", "./runner.py"]
