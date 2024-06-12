import {Button, Stack, Typography} from "@mui/material";
import {useContext, useState} from "react";
import AuthContext from "../../utils/auth.ts";
import {useNavigate} from "react-router-dom";
import {useTranslation} from "react-i18next";

export default function DeviceData({device, slotID}: {device: Device | null, slotID: bigint}) {
    const { t } = useTranslation();
    const jwt = useContext(AuthContext);
    const navigate = useNavigate();
    const [deviceJWTState, setDeviceJWTState] = useState("");

    async function createDevice() {
        const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/register/device', {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": "Bearer " + jwt
            },
            body: JSON.stringify({"slotID": slotID})
        });
        if (!response.ok) {
            throw new Error('Failed');
        }
        const data = await response.json() as { status: boolean, body: { jwt: string } };
        if (!data.status) {
            throw new Error('Failed ');
        }

        setDeviceJWTState(() => data.body.jwt)
    }

    async function deleteDevice(ID: string) {
        const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/register/device', {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
                "Authorization": "Bearer " + jwt
            },
            body: JSON.stringify({"deviceID": ID})
        });
        if (!response.ok) {
            throw new Error('Failed!');
        }
        console.log(response.body)
        const data = await response.json() as { status: boolean };
        if (!data.status) {
            throw new Error('Failed to login');
        }

        navigate("/")
    }

    return (
        <>
            {device && device.ID ?
                <>
                    <Typography variant="h5">{t("Device Connected")}</Typography>
                    <Button variant="outlined" color={"error"} onClick={() => deleteDevice(device?.ID)}>
                        {t("Delete")}
                    </Button>
                </>
                : <>
                    <Typography variant="h5">{t("No Device")}</Typography>
                    {deviceJWTState === "" ?
                        <Button variant="contained" color={"primary"} onClick={createDevice}>
                            {t("Create Device")}
                        </Button>
                        : <Stack
                            direction="column"
                            spacing={2}
                        ><Typography variant="h6">{t("Device JWT")}: {deviceJWTState}</Typography></Stack>
                    }
                </>
            }
        </>
    )
}