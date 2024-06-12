import LoginForm from "../components/LoginForm/LoginForm.tsx";
import {useNavigate} from "react-router-dom";

export default function Login({setUserToken}: {setUserToken: React.Dispatch<React.SetStateAction<string>>}){
    const navigate = useNavigate();

    async function login(userToken: string) {
        const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/login', {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({"googleJWT": userToken})
        });
        if (!response.ok) {
            throw new Error('Failed to login');
        }
        console.log(response.body)
        const data = await response.json() as { status: boolean, body:{ jwt: string } };
        if (!data.status) {
            throw new Error('Failed to login');
        }
        setUserToken(() => data.body.jwt);

        navigate("/")
    }
    const googleClientId = import.meta.env.VITE_GOOGLE_CLIENT_ID as string;

    return (
        <>
           <LoginForm googleClientId={googleClientId} onSubmit={login}/>
        </>
    )
}