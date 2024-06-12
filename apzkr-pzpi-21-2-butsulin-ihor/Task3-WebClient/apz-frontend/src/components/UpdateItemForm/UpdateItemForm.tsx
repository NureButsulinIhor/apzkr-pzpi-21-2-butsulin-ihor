import {Button, FormControl, InputAdornment, Paper, Stack, TextField, Typography} from "@mui/material";
import {useContext, useState} from "react";
import AuthContext from "../../utils/auth.ts";
import {useNavigate} from "react-router-dom";
import {useTranslation} from "react-i18next";

export default function UpdateItemForm({item}: {item: Item}) {
    const { t } = useTranslation();
    const jwt = useContext(AuthContext);
    const navigate = useNavigate();
    const [nameState, setNameState] = useState(item.name);
    const [descriptionState, setDescriptionState] = useState(item.description);
    const [weightState, setWeightState] = useState(item.weight);

    async function updateItem(name: string, description: string, weight: number) {
        const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/admin/item', {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
                "Authorization": "Bearer " + jwt
            },
            body: JSON.stringify({"itemID": item.ID, "name": name, "description": description, "weight": weight})
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
                updateItem(nameState, descriptionState, weightState);
            }}>
                <Paper style={{padding: '1rem'}} elevation={4}>
                    <Stack
                        direction="column"
                        spacing={2}
                    >
                        <Typography align="center" variant="h6">
                            {t("Update Item")}
                        </Typography>
                        <FormControl fullWidth>
                            <TextField
                                type="text"
                                label={t("Name")}
                                variant="outlined"
                                value={nameState}
                                onChange={(e) => {
                                    setNameState(e.target.value)
                                }}
                            />
                        </FormControl>
                        <FormControl fullWidth>
                            <TextField
                                type="text"
                                label={t("Description")}
                                variant="outlined"
                                value={descriptionState}
                                onChange={(e) => {
                                    setDescriptionState(e.target.value)
                                }}
                            />
                        </FormControl>
                        <FormControl fullWidth>
                            <TextField
                                type="number"
                                label={t("Weight")}
                                variant="outlined"
                                value={weightState}
                                onChange={(e) => {
                                    setWeightState(+e.target.value)
                                }}
                                InputProps={{
                                    endAdornment: <InputAdornment position="end">{t("kg")}</InputAdornment>,
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