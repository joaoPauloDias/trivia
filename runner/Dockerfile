FROM python:3.8-slim

WORKDIR /app

COPY runner.py .

RUN pip install requests

ENTRYPOINT ["python3", "./runner.py"]
CMD ["args"]
