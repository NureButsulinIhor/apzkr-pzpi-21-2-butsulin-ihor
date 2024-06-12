import {Route, Routes} from "react-router-dom";
import {useState} from "react";
import Layout from "./Layout.tsx";
import UnauthorisedHeader from "./components/UnauthorisedHeader/UnauthorisedHeader.tsx";
import Login from "./routes/Login.tsx";
import AuthorisedHeader from "./components/AuthorisedHeader/AuthorisedHeader.tsx";
import Index from "./routes/Index.tsx";
import AuthContext from "./utils/auth.ts";
import Warehouse from "./routes/Warehouse.tsx";
import Slot from "./routes/Slot.tsx";
import Cars from "./routes/Cars.tsx";
import Car from "./routes/Car.tsx";
import Managers from "./routes/Managers.tsx";

export default function App() {
    const [userToken, setUserToken] = useState("");

    function isUserAuthorised() {
        return userToken !== "" && userToken !== null && userToken !== undefined;
    }

    return (
        <Routes>
            {isUserAuthorised()
                ? <Route element={
                    <Layout>
                        <AuthorisedHeader/>
                    </Layout>
                }>
                    <Route path="/" element={<AuthContext.Provider value={userToken}><Index/></AuthContext.Provider>} />
                    <Route path="/warehouse/:warehouseID" element={<AuthContext.Provider value={userToken}><Warehouse /> </AuthContext.Provider>} />
                    <Route path="/cars" element={<AuthContext.Provider value={userToken}><Cars/></AuthContext.Provider>} />
                    <Route path="/managers" element={<AuthContext.Provider value={userToken}><Managers/></AuthContext.Provider>} />
                    <Route path="/warehouse/:warehouseID" element={<AuthContext.Provider value={userToken}><Warehouse /> </AuthContext.Provider>} />
                    <Route path="/car/:carID" element={<AuthContext.Provider value={userToken}><Car /> </AuthContext.Provider>} />
                    <Route path="/slot/:slotID" element={<AuthContext.Provider value={userToken}><Slot /> </AuthContext.Provider>} />
                </Route>
                : <Route element={
                    <Layout>
                        <UnauthorisedHeader/>
                    </Layout>
                }>
                    <Route path="/" element={<Login  setUserToken={setUserToken}/>} />
                </Route>
            }
        </Routes>
    );
}