import {Button, FormControl, Paper, Stack, TextField, Typography} from "@mui/material";
import {useContext, useState} from "react";
import AuthContext from "../../utils/auth.ts";
import {useNavigate} from "react-router-dom";
import {useTranslation} from "react-i18next";

export default function CreateManagerForm() {
    const { t } = useTranslation();
    const jwt = useContext(AuthContext);
    const navigate = useNavigate();
    const [emailState, setEmailState] = useState("");

    async function createManager(email: string) {
        const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/register/manager', {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": "Bearer " + jwt
            },
            body: JSON.stringify({"email": email})
        });
        if (!response.ok) {
            throw new Error('Failed');
        }
        const data = await response.json() as { status: boolean };
        if (!data.status) {
            throw new Error('Failed ');
        }

        navigate("/")
    }

    return (
        <Paper style={{padding: '1rem'}} elevation={4}>
            <form onSubmit={(e) => {
                e.preventDefault();
                createManager(emailState);
            }}>
                <Paper style={{padding: '1rem'}} elevation={4}>
                    <Stack
                        direction="column"
                        spacing={2}
                    >
                        <Typography align="center" variant="h6">
                            {t("Create Manager")}
                        </Typography>
                        <FormControl fullWidth>
                            <TextField
                                type="email"
                                label={t("Email")}
                                variant="outlined"
                                value={emailState}
                                onChange={(e) => {
                                    setEmailState(e.target.value)
                                }}
                            />
                        </FormControl>

                        <Button type="submit">{t("Submit")}</Button>
                    </Stack>
                </Paper>

            </form>

        </Paper>
    );
}