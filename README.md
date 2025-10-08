Currently Implementing 

# New Community Microservices Architecture

## Gateway / BFF

| Domain       | Microservice   | Responsibility                                         | Language / Rationale                     | Gateway Exposure |
|-------------|----------------|-------------------------------------------------------|-----------------------------------------|----------------|
| Gateway / BFF | API Gateway   | Single frontend entry point, request routing, auth, aggregation, rate limiting | **Node.js** – JSON-first, high concurrency | **Exposed** |

## Catalog

| Domain  | Microservice          | Responsibility                        | Language / Rationale              | Gateway Exposure |
|--------|----------------------|--------------------------------------|----------------------------------|----------------|
| Catalog | Product-Write Service | Create/update/delete products        | **Java** – transactional          | Internal       |
|        | Product-Read Service  | Serve product data (cache + DB)      | **Java** – fast, consistent reads | Exposed        |
|        | Category Service      | Manage categories & hierarchies      | **Java** – hierarchical rules support | Exposed   |
|        | Brand Service         | Manage brands, manufacturers         | **Java** – simple CRUD            | Exposed        |

## Pricing & Promotions

| Domain                 | Microservice           | Responsibility                     | Language / Rationale         | Gateway Exposure |
|------------------------|----------------------|-----------------------------------|-----------------------------|----------------|
| Pricing & Promotions   | List-Price Service    | Base list price management         | **Java** – pricing rules, audit | Internal    |
|                        | Discount-Engine Service | Compute promotions, coupons       | **Java** – rules engine      | Internal       |
|                        | Tax Calculation Service | Apply tax rules by region         | **Java** – deterministic     | Internal       |

## Inventory

| Domain    | Microservice           | Responsibility                        | Language / Rationale           | Gateway Exposure |
|-----------|-----------------------|--------------------------------------|-------------------------------|----------------|
| Inventory | Inventory-Read Service | Real-time stock lookups (cache backed) | **Java** – thread-safe reads  | Exposed        |
|           | Inventory-Write Service | Reserve/release stock (atomic ops)  | **Java** – concurrency-safe updates | Internal |
|           | Warehouse Sync Service | Sync stock with ERP                   | **Python** – orchestration + API integration | Internal |

## Cart & Checkout

| Domain        | Microservice            | Responsibility                        | Language / Rationale          | Gateway Exposure |
|---------------|------------------------|--------------------------------------|-------------------------------|----------------|
| Cart & Checkout | Cart-CRUD Service     | CRUD cart items                        | **Node.js** – JSON, session-heavy | Exposed    |
|               | Cart-Pricing Service    | Apply pricing + promotions            | **Java** – reuse pricing rules | Internal      |
|               | Cart-Snapshot Service   | Persist cart snapshot                  | **Java** – fast persistence    | Internal      |
|               | Checkout-Orchestrator Service | Coordinate cart → payment → inventory events | **Python (FastAPI)** – orchestration | Internal |

## Orders & Payments

| Domain           | Microservice          | Responsibility                        | Language / Rationale           | Gateway Exposure |
|-----------------|---------------------|--------------------------------------|-------------------------------|----------------|
| Orders & Payments | Order-Create Service | Validate & create order record        | **Java** – transactional integrity | Internal |
|                  | Order-Read Service   | Fetch order status/history            | **Java** – fast queries       | Exposed       |
|                  | Payment-Gateway Adapter | External gateway integration        | **Java** – reliable retries   | Internal      |
|                  | Wallet-Service       | Internal wallet debits/credits        | **Java** – transactional      | Internal      |
|                  | Refund Service       | Process refunds, reversal workflows   | **Python** – workflow orchestration | Internal |
|                  | Invoice Service      | Generate & store invoice PDFs          | **Python** – async + flexible PDF libs | Internal |

## User & Auth

| Domain      | Microservice   | Responsibility                | Language / Rationale           | Gateway Exposure |
|------------|----------------|------------------------------|-------------------------------|----------------|
| User & Auth | User-Service  | Profile CRUD, addresses      | **Java** – domain modeling    | Exposed        |
|            | Auth-Service   | Login, JWT mint/verify, MFA  | **Java** – security critical  | Exposed (auth flow) |

## Search & Recommendation

| Domain                 | Microservice           | Responsibility                   | Language / Rationale     | Gateway Exposure |
|------------------------|----------------------|---------------------------------|-------------------------|----------------|
| Search & Recommendation | Search-Index Service | Build/update search index        | **Java** – indexing batch | Internal      |
|                        | Search-Query Service | Query index via typo-tolerant API | **Java** – low-latency  | Exposed       |
|                        | Rec-Model Service    | Serve ML models for recommendations | **Python** – ML ecosystem | Exposed    |

## Notifications

| Domain        | Microservice    | Responsibility          | Language / Rationale      | Gateway Exposure |
|---------------|----------------|------------------------|---------------------------|----------------|
| Notifications | Email-Service  | Send transactional emails | **Python** – async APIs  | Internal       |
|               | SMS-Service    | Send SMS               | **Python** – async       | Internal       |
|               | Push-Service   | Mobile & web push      | **Node.js** – websocket/event-driven | Exposed |

## Observability & Ops

| Domain                | Microservice       | Responsibility                   | Language / Rationale          | Gateway Exposure |
|-----------------------|------------------|---------------------------------|-------------------------------|----------------|
| Observability & Ops   | Metrics-Exporter  | Expose Prometheus metrics        | **Java** – uniform metrics    | Internal       |
|                       | Tracing-Collector | Collect & forward OpenTelemetry spans | **Java** – centralized   | Internal       |
|                       | Log-Forwarder    | Ship logs to DataDog             | **Java** – standardized       | Internal       |
|                       | Config-Service   | Distribute dynamic config via gRPC/HTTP | **Java** – reliable config distribution | Internal |

## AI Layer

| Domain    | Microservice       | Responsibility                   | Language / Rationale           | Gateway Exposure |
|-----------|------------------|---------------------------------|-------------------------------|----------------|
| AI Layer  | Intent-Parser     | NLU → structured intents        | **Python (FastAPI)** – NLP    | Exposed (AI features) |
|           | Plan-Generator    | Turn intent into JSON workflow   | **Python** – orchestration     | Internal       |
|           | Plan-Executor    | Execute workflows (calls & Kafka events) | **Python** – async + workflow management | Internal |
|           | Audit-Service    | Persist AI decisions & allow human review | **Java** – audit compliance | Internal       |
