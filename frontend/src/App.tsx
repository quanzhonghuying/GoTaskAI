import { BrowserRouter, Routes, Route, Link } from "react-router-dom";
import Login from "./pages/Login";
import Register from "./pages/Register";

function App() {
    return (
        <BrowserRouter>
            <div>
                {/* ナビゲーションバー */}
                <nav className="bg-gray-800 p-4">
                    <ul className="flex space-x-6 text-white font-medium">
                        <li>
                            <Link to="/login" className="hover:text-blue-400">ログイン</Link>
                        </li>
                        <li>
                            <Link to="/register" className="hover:text-blue-400">ユーザー登録</Link>
                        </li>
                    </ul>
                </nav>


                {/* ページ */}
                <Routes>
                    <Route path="/login" element={<Login />} />
                    <Route path="/register" element={<Register />} />
                </Routes>
            </div>
        </BrowserRouter>
    );
}

export default App;
