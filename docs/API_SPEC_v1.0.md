# API仕様書 v1.0

## 1. 概要
本書は **GoTaskAI** システムにおける RESTful API の仕様を定義する。  
対象範囲：タスク管理 (tasks) ＋ ユーザー管理 (users, 基本機能)。  

---

## 2. 共通仕様

- Base URL: `/api/v1`
- データ形式: JSON
- 認証方式: （未実装、将来的に JWT を導入予定）
- エラーレスポンス:
  - `400 Bad Request` → リクエストパラメータ不正
  - `404 Not Found` → リソース未存在
  - `500 Internal Server Error` → サーバ内部エラー

---

## 3. API 定義

### 3.1 タスク管理 (Tasks)

#### (1) タスク一覧取得
- **GET** `/tasks`
- **Response (200)**
```json
[
  {
    "id": 1,
    "title": "設計ドキュメント作成",
    "status": "open",
    "user_id": 101,
    "created_at": "2025-09-16T10:00:00Z",
    "updated_at": "2025-09-16T11:00:00Z"
  }
]
## 3.1 タスク管理 (Tasks)

### (2) タスク新規作成
- **POST** `/tasks`
- **Request**
```json
{
  "title": "Go API 開発学習",
  "status": "open",
  "user_id": 101
}
```