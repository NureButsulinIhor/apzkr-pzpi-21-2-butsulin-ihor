import {Paper, Stack, Typography} from "@mui/material";
import {useEffect} from "react";
import {useTranslation} from "react-i18next";

type CredentialResponse = {
    credential: string;
}

export default function LoginForm({ googleClientId, onSubmit }: { googleClientId: string, onSubmit: (googleToken: string) => void }) {
    const { t } = useTranslation();
    google.accounts.id.initialize({
        client_id: googleClientId,
        callback: (response: CredentialResponse) => onSubmit(response.credential)
    });

    useEffect(() => {
        google.accounts.id.renderButton(document.getElementById("loginButton"), {
            theme: "outline",
            size: "large",
            text: "continue_with",
            locale: "en"
        });
    }, []);

    return (
        <Paper style={{padding: '1rem'}} elevation={4}>
            <Stack
                direction="column"
                spacing={2}
            >
                <Typography align="center" variant="h6">
                    {t("Login")}
                </Typography>

                <div id={"loginButton"}></div>
            </Stack>
        </Paper>
    );
}