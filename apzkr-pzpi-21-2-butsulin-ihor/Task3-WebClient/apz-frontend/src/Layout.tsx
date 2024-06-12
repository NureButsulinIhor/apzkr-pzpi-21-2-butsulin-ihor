import {Outlet} from "react-router-dom";
import './main.css'

export default function Layout({children}: {children: React.ReactNode}) {
    return (
        <>
            <header>
                {children}
            </header>
            <main>
                {/* 2️⃣ Render the app routes via the Layout Outlet */}
                <Outlet />
            </main>
            <footer>osigram ©️ 2024</footer>
        </>
    );
}