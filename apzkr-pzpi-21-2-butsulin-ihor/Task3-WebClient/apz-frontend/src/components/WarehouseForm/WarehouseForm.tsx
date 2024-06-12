import {Button, FormControl, InputLabel, MenuItem, Paper, Select, Stack, Typography} from "@mui/material";
import {useContext, useEffect, useState} from "react";
import AuthContext from "../../utils/auth.ts";
import {useNavigate} from "react-router-dom";
import {useTranslation} from "react-i18next";

export default function WarehouseForm() {
    const { t } = useTranslation();
    const jwt = useContext(AuthContext);
    const navigate = useNavigate();
    const [managers, setManagers] = useState([] as User[]);
    const [managerID, setManagerID] = useState(BigInt(0));

    async function createWarehouse(managerUserID: bigint) {
        const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/admin/warehouse', {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": "Bearer " + jwt
            },
            body: JSON.stringify({"managerUserID": managerUserID})
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

    async function fetchManagers() {
        const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/register/manager/all', {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": "Bearer " + jwt
            },
        });
        if (!response.ok) {
            throw new Error('Failed');
        }
        const data = await response.json() as { status: boolean, body: { managers: Manager[], unsetManagers: User[] } };
        if (!data.status) {
            throw new Error('Failed ');
        }

        setManagers(() => data.body.unsetManagers);
    }

    useEffect(() => {
        fetchManagers();
    }, []);

    return (
        <Paper style={{padding: '1rem'}} elevation={4}>
            <form onSubmit={(e) => {
                e.preventDefault();
                createWarehouse(managerID);
            }}>
                <Paper style={{padding: '1rem'}} elevation={4}>
                    <Stack
                        direction="column"
                        spacing={2}
                    >
                        <Typography align="center" variant="h6">
                            {t("Create Warehouse")}
                        </Typography>
                        <FormControl fullWidth>
                            <InputLabel id="CreateWarehouse-select-label">Email</InputLabel>
                            <Select
                                labelId="CreateWarehouse-select-label"
                                id="CreateWarehouse-select"
                                value={managerID}
                                label={t("Email")}
                                onChange={(e) => setManagerID(e.target.value as bigint)}
                            >
                                {managers && managers.map((manager) => (
                                    <MenuItem key={"MenuItemCreateWarehouseManager-"+manager.ID.toString()} value={manager.ID}>{manager.email}</MenuItem>
                                ))}
                            </Select>
                        </FormControl>

                        <Button type="submit">{t("Submit")}</Button>
                    </Stack>
                </Paper>

            </form>

        </Paper>
    );
}