import React, { useState } from "react";

function Login() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [message, setMessage] = useState("");

    const handleLogin = async () => {
        try {
            const response = await fetch("http://localhost:8080/login", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ email, password }),
            });

            if (!response.ok) throw new Error("ログイン失敗：サーバーエラー");

            const data = await response.json();
            localStorage.setItem("token", data.token);
            setMessage("✅ ログイン成功！トークンを保存しました");
        } catch (err: any) {
            setMessage("❌ 通信エラー：" + err.message);
        }
    };

    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100">
            {/* カード */}
            <div className="bg-white shadow-lg rounded-lg p-8 w-full max-w-md">
                <h1 className="text-2xl font-bold mb-6 text-center text-gray-800">
                    ユーザーログイン
                </h1>

                {/* メール */}
                <label className="block mb-2 text-sm font-medium text-gray-700">
                    メールアドレス
                </label>
                <input
                    type="email"
                    placeholder="メールアドレスを入力"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    className="w-full mb-4 px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                />

                {/* パスワード */}
                <label className="block mb-2 text-sm font-medium text-gray-700">
                    パスワード
                </label>
                <input
                    type="password"
                    placeholder="パスワードを入力"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    className="w-full mb-6 px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                />

                {/* ボタン */}
                <button
                    onClick={handleLogin}
                    className="w-full bg-blue-600 text-white py-2 rounded-lg hover:bg-blue-700 transition"
                >
                    ログイン
                </button>

                {/* メッセージ */}
                {message && (
                    <p className="mt-4 text-center text-sm text-gray-600">{message}</p>
                )}
            </div>
        </div>
    );
}

export default Login;
