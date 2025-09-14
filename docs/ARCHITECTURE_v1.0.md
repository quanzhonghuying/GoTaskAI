# GoTaskAI アーキテクチャ設計書 v1.0

---

## 1. システム構成概要
GoTaskAI は以下のコンポーネントで構成される Web アプリケーションである。

- **フロントエンド (React)**  
  ユーザーインターフェース。ログイン、タスク管理、AI 提案の表示を担当。

- **バックエンド (Go + Gin)**  
  REST API サーバー。認証、タスク CRUD、AI サービスとの連携を提供。

- **データベース (PostgreSQL)**  
  ユーザー情報とタスクデータを永続化。

- **AI サービス (OpenAI API)**  
  タスク内容の要約、リマインダー提案などを自動生成。

- **AWS インフラ**  
  デプロイ環境 (EC2, RDS, S3 など) を提供し、可用性・拡張性を確保。

---

## 2. アーキテクチャ構成図
```mermaid
graph TD
    FE[フロントエンド (React)]
    BE[バックエンド (Go + Gin)]
    DB[(PostgreSQL)]
    AI[OpenAI API]
    AWS[(AWS Infrastructure)]

    FE --> BE
    BE --> DB
    BE --> AI
    BE --> AWS
## 3. モジュール構成

### 3.1 フロントエンド (React)
- ログイン / 登録画面  
- タスク一覧画面  
- タスク編集 / 作成画面  
- AI 提案の表示画面  

### 3.2 バックエンド (Go + Gin)
- 認証モジュール（JWT）  
- ユーザー管理モジュール  
- タスク管理モジュール（CRUD）  
- AI 連携モジュール（OpenAI API 呼び出し）  
- AWS 連携モジュール（S3 アップロード等）  

### 3.3 データベース (PostgreSQL)
- USERS テーブル  
- TASKS テーブル  
- インデックス設計による検索性能最適化  

### 3.4 インフラ (AWS)
- **ECS (Fargate)** または **Lambda**: API サーバー実行環境  
- **RDS (PostgreSQL)**: データベース  
- **S3**: 静的ファイル・添付データ保存  
- **CloudFront**: フロントエンド配信  
- **CloudWatch**: ログ収集・モニタリング  

---

## 4. データフロー
```mermaid
sequenceDiagram
    participant User as ユーザー
    participant FE as フロントエンド(React)
    participant BE as バックエンド(Go API)
    participant DB as データベース(PostgreSQL)
    participant AI as OpenAI API

    User->>FE: ログイン情報入力
    FE->>BE: /api/v1/login (POST)
    BE->>DB: 認証確認
    DB-->>BE: 認証結果
    BE-->>FE: JWT トークン返却

    User->>FE: タスク作成要求
    FE->>BE: /api/v1/tasks (POST)
    BE->>DB: タスク保存
    DB-->>BE: 保存成功
    BE->>AI: タスク内容送信
    AI-->>BE: 要約・提案返却
    BE-->>FE: 結果表示

## 5. API 設計（概要）

| エンドポイント | メソッド | 概要 |
|----------------|---------|------|
| `/api/v1/register` | POST | ユーザー登録 |
| `/api/v1/login` | POST | ログイン（JWT 発行） |
| `/api/v1/tasks` | GET | タスク一覧取得 |
| `/api/v1/tasks` | POST | タスク作成 |
| `/api/v1/tasks/{id}` | PUT | タスク更新 |
| `/api/v1/tasks/{id}` | DELETE | タスク削除 |
| `/api/v1/tasks/{id}/ai` | GET | AI 提案取得 |

---

## 6. データベース設計（概要）

### users テーブル
| カラム名 | 型 | 制約 | 説明 |
|----------|----|------|------|
| id | SERIAL | PK | ユーザーID |
| email | VARCHAR(255) | UNIQUE, NOT NULL | メールアドレス |
| password_hash | VARCHAR(255) | NOT NULL | パスワードハッシュ |
| created_at | TIMESTAMP | DEFAULT now() | 作成日時 |

### tasks テーブル
| カラム名 | 型 | 制約 | 説明 |
|----------|----|------|------|
| id | SERIAL | PK | タスクID |
| user_id | INT | FK → users.id | 所属ユーザー |
| title | VARCHAR(255) | NOT NULL | タスクタイトル |
| description | TEXT | - | タスク詳細 |
| completed | BOOLEAN | DEFAULT false | 完了フラグ |
| created_at | TIMESTAMP | DEFAULT now() | 作成日時 |

---

## 7. セキュリティ設計
- JWT による認証  
- bcrypt によるパスワードハッシュ化  
- HTTPS (TLS) 通信必須  
- SQL インジェクション対策（ORM 利用）  
- CORS 制御（許可ドメイン制御）  

---

## 8. 非機能要件対応
- **性能**: 同時接続 500 ユーザーを想定  
- **可用性**: AWS Auto Scaling により冗長化  
- **拡張性**: モジュール単位で機能追加可能  
- **監視**: CloudWatch によるモニタリング + アラート設定  

---

## 9. 今後の拡張
- 多言語対応（日本語 / 英語）  
- 通知機能（メール / Slack 連携）  
- チームコラボレーション機能  
- 外部 API 連携（Google Calendar, Notion など）  