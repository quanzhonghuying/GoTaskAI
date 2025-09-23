import React, { useState } from "react";

function Login(): JSX.Element {
    const [email, setEmail] = useState<string>("");
    const [password, setPassword] = useState<string>("");
    const [message, setMessage] = useState<string>("");

    const handleLogin = async (): Promise<void> => {
        try {
            const response = await fetch("http://localhost:8080/login", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ email, password }),
            });

            if (!response.ok) {
                throw new Error("ログイン失敗：サーバーエラー");
            }

            const data: { token: string } = await response.json();
            localStorage.setItem("token", data.token);
            setMessage("ログイン成功！トークンを保存しました");
        } catch (err) {
            if (err instanceof Error) {
                setMessage("通信エラー：" + err.message);
            } else {
                setMessage("不明なエラーが発生しました");
            }
        }
    };

    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-100">
            {/* 中央寄せのカード */}
            <div className="bg-white p-8 rounded-2xl shadow-lg w-full max-w-md">
                {/* タイトル */}
                <h1 className="text-2xl font-bold text-center mb-6 text-gray-800">
                    ユーザーログイン
                </h1>

                {/* メールアドレス */}
                <input
                    type="email"
                    placeholder="メールアドレス"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    className="w-full mb-4 px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                />

                {/* パスワード */}
                <input
                    type="password"
                    placeholder="パスワード"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    className="w-full mb-6 px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                />

                {/* ログインボタン */}
                <button
                    onClick={handleLogin}
                    className="w-full bg-blue-600 text-white py-2 rounded-lg hover:bg-blue-700 transition"
                >
                    ログイン
                </button>

                {/* 成功／失敗メッセージ */}
                {message && (
                    <div
                        className={`mt-4 p-3 rounded-lg text-sm font-medium ${message.includes("成功")
                                ? "bg-green-100 text-green-800"
                                : "bg-red-100 text-red-800"
                            }`}
                    >
                        {message}
                    </div>
                )}
            </div>
        </div>
    );
}

export default Login;
