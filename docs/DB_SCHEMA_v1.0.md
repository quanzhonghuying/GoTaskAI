# GoTaskAI データベース設計書 v1.0

---

## 1. 概要
本ドキュメントは、GoTaskAI システムで利用するデータベース (PostgreSQL) の詳細設計を記載する。  
基本設計書に記載した概要に基づき、テーブル定義・ER 図・インデックス設計を明確化する。

---

## 2. ER 図
```mermaid
erDiagram
    USERS ||--o{ TASKS : has
    USERS {
        int id PK
        varchar email
        varchar password_hash
        timestamp created_at
    }
    TASKS {
        int id PK
        int user_id FK
        varchar title
        text description
        boolean completed
        timestamp created_at
        timestamp updated_at
    }

