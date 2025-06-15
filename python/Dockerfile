FROM python:3.13.3-alpine
WORKDIR /app
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt
COPY src/ ./src
WORKDIR /app/src
ENTRYPOINT ["python", "server.py"]
