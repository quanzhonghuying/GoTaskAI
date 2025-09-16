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
```

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
- **Response (201)**
```json
{
  "id": 2,
  "message": "Task created"
}
```
### (3) タスク更新
- **PUT /tasks/{id}**
- **Request**
```json
{
  "title": "Go API 開発学習（完了）",
  "status": "done"
}
```
- **Response (200)**
```json
{
  "id": 2,
  "message": "Task updated"
}
```
### (4) タスク削除
- **DELETE /tasks/{id}**
- **Response (200)**
```json
{
  "id": 2,
  "message": "Task deleted"
}
```
### 3.2 ユーザー管理 (Users, 基本機能)
### (1) ユーザー一覧取得
- **GET /users**
- **Response (200)**
```json
[
  {
    "id": 101,
    "name": "Alice",
    "email": "alice@example.com"
  }
]
```
### (2) ユーザー新規作成
- **POST /users**
- **Request**
```json
{
  "name": "Bob",
  "email": "bob@example.com"
}
```
- **Response (201)**
```json
{
  "id": 102,
  "message": "User created"
}    
```
### (4) エラーレスポンス例
```json
{
  "error": "Invalid request payload"
}
```
### 5. 今後の拡張予定

- **ユーザー認証機能 (JWT)**
- **タスクにタグ付与 (tags)**
- **ページング取得 (例: tasks?limit=10&offset=0)**
