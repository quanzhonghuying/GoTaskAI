# GoTaskAI 基本設計書 v1.0

## 1. システム構成概要
GoTaskAI は **フロントエンド（React）**、**バックエンド（Go API）**、**データベース（PostgreSQL）**、  
および **AI サービス（OpenAI API）** を組み合わせた Web アプリケーションである。  
AWS 上にデプロイし、可用性と拡張性を確保する。

---

## 2. アーキテクチャ構成図
```mermaid
graph TD
    FE[Frontend (React)]
    BE[Backend (Go + Gin)]
    DB[(PostgreSQL)]
    AI[OpenAI API]
    AWS[(AWS Infrastructure)]

    FE --> BE
    BE --> DB
    BE --> AI
    BE --> AWS
```
## 3. モジュール構成

### 3.1 フロントエンド (React)
- ログイン / 登録画面  
- タスク一覧画面  
- タスク編集 / 作成画面  
- AI 提案の表示  

### 3.2 バックエンド (Go + Gin)
- 認証モジュール（JWT）  
- ユーザー管理モジュール  
- タスク管理モジュール（CRUD）  
- AI 連携モジュール（OpenAI API 呼び出し）  

### 3.3 データベース (PostgreSQL)
- USERS テーブル  
- TASKS テーブル  

### 3.4 インフラ (AWS)
- API Gateway / ECS (Fargate) または Lambda  
- RDS（PostgreSQL）  
- S3（静的ファイル保存）  
- CloudFront（フロントエンド配信）  

---

## 4. データフロー
1. ユーザーは React フロントエンドから操作する  
2. フロントエンドは Go API にリクエストを送信  
3. Go API は DB でデータを管理し、必要に応じて AI API を呼び出す  
4. 処理結果をフロントエンドに返却する  

---

## 5. API 設計（概要）
- **POST /api/v1/register** → ユーザー登録  
- **POST /api/v1/login** → ログイン（JWT 返却）  
- **GET /api/v1/tasks** → タスク一覧  
- **POST /api/v1/tasks** → タスク作成  
- **PUT /api/v1/tasks/{id}** → タスク更新  
- **DELETE /api/v1/tasks/{id}** → タスク削除  
- **POST /api/v1/tasks/{id}/summarize** → AI 要約  

詳細は `API_SPEC_v1.0.md` を参照する。  

---

## 6. データベース設計（概要）

### USERS テーブル
- id (PK)  
- username  
- email  
- password_hash  
- created_at  

### TASKS テーブル
- id (PK)  
- user_id (FK → USERS.id)  
- title  
- description  
- completed  
- created_at  
- updated_at  

詳細は `DB_SCHEMA_v1.0.md` を参照する。  

---

## 7. セキュリティ設計
- 認証方式：JWT（JSON Web Token）  
- 通信方式：HTTPS  
- DB アクセス：ユーザーごとにタスクを分離  
- パスワード保存：ハッシュ化（bcrypt）  

---

## 8. 非機能要件対応
- **パフォーマンス**：キャッシュ導入を検討（Redis など）  
- **可用性**：AWS SLA に基づき 99.9% 稼働  
- **拡張性**：マイクロサービス化を前提に API 設計  
- **運用性**：GitHub Actions による CI/CD、自動デプロイ  

---

## 9. 今後の拡張
- チームコラボレーション機能  
- 通知機能（メール / Slack）  
- 多言語対応（日本語・英語・中国語など）  
