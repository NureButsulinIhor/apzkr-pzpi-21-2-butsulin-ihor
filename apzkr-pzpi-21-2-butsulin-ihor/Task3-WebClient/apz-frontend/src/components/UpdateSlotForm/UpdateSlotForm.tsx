import {Button, FormControl, InputAdornment, Paper, Stack, TextField, Typography} from "@mui/material";
import {useContext, useState} from "react";
import AuthContext from "../../utils/auth.ts";
import {useNavigate} from "react-router-dom";
import {useTranslation} from "react-i18next";

export default function UpdateSlotForm({slot}: {slot: Slot}) {
    const { t } = useTranslation();
    const jwt = useContext(AuthContext);
    const navigate = useNavigate();
    const [maxWeightState, setMaxWeightState] = useState(slot.maxWeight);

    async function updateSlot(maxWeight: number) {
        const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/admin/slot', {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
                "Authorization": "Bearer " + jwt
            },
            body: JSON.stringify({"slotID": slot.ID, "maxWeight": maxWeight})
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
                updateSlot(maxWeightState);
            }}>
                <Paper style={{padding: '1rem'}} elevation={4}>
                    <Stack
                        direction="column"
                        spacing={2}
                    >
                        <Typography align="center" variant="h6">
                            {t("Update Slot")}
                        </Typography>
                        <FormControl fullWidth>
                            <TextField
                                type="number"
                                label={t("Max weight")}
                                variant="outlined"
                                value={maxWeightState}
                                onChange={(e) => {
                                    setMaxWeightState(+e.target.value)
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