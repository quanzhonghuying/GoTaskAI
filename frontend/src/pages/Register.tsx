import { useState } from "react";

export default function Register() {
    // メールアドレスの state
    const [email, setEmail] = useState("");
    // パスワードの state
    const [password, setPassword] = useState("");
    // 成功 / 失敗メッセージの state
    const [message, setMessage] = useState("");

    // 登録処理
    const handleRegister = async () => {
        try {
            const response = await fetch("http://localhost:8080/register", {
                method: "POST", // HTTP メソッド
                headers: { "Content-Type": "application/json" }, // JSON 形式で送信
                body: JSON.stringify({ email, password }), // 入力値を JSON に変換
            });

            if (!response.ok) {
                // 200 系以外のステータスコードの場合はエラー
                throw new Error("登録失敗: " + response.status);
            }

            const data = await response.json();
            // サーバーから返却されたレスポンスを JSON に変換
            setMessage(data.message || "ユーザー登録成功");
        } catch (err: any) {
            // エラーメッセージをセット
            setMessage("通信エラー: " + err.message);
        }
    };

    return (
        // 背景と中央寄せのレイアウト
        <div className="min-h-screen flex items-center justify-center bg-gray-100">
            {/* 登録カード */}
            <div className="bg-white p-8 rounded-2xl shadow-lg w-full max-w-md">
                {/* タイトル */}
                <h1 className="text-2xl font-bold text-center mb-6 text-gray-800">
                    ユーザー登録
                </h1>

                {/* メールアドレス入力 */}
                <input
                    type="email"
                    placeholder="メールアドレス"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    className="w-full mb-4 px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                />

                {/* パスワード入力 */}
                <input
                    type="password"
                    placeholder="パスワード"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    className="w-full mb-6 px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                />

                {/* 登録ボタン */}
                <button
                    onClick={handleRegister}
                    className="w-full bg-green-600 text-white py-2 rounded-lg hover:bg-green-700 transition-colors"
                >
                    登録
                </button>

                {/* 結果メッセージ */}
                {message && (
                    <p
                        className={`mt-4 text-center text-sm ${message.includes("成功")
                                ? "text-green-600 bg-green-100 p-2 rounded-lg"
                                : "text-red-600 bg-red-100 p-2 rounded-lg"
                            }`}
                    >
                        {message}
                    </p>
                )}
            </div>
        </div>
    );
}
