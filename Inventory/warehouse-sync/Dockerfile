FROM python:3.11-slim AS builder

WORKDIR /app

COPY pyproject.toml requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

# Stage 2: Production runtime
FROM python:3.11-slim AS runtime

WORKDIR /app
ENV PYTHONPATH=/app/src

COPY . .
RUN pip install --no-cache-dir -r requirements.txt

# Export 
EXPOSE 8000

CMD ["python","src/cmd/run_sync.py"]