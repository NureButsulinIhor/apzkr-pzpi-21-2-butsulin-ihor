import {useContext, useEffect, useState} from "react";
import AuthContext from "../utils/auth.ts";
import {Typography} from "@mui/material";
import ManagersList from "../components/ManagersList/ManagersList.tsx";
import CreateManagerForm from "../components/CreateManagerForm/CreateManagerForm.tsx";
import {useTranslation} from "react-i18next";

export default function Managers() {
    const { t } = useTranslation();
    const userToken = useContext(AuthContext);
    const [managers, setManagers] = useState([] as User[]);

    async function fetchManagers() {
        const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/register/manager/all', {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": "Bearer " + userToken
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
      <>
          {managers && <>
                {managers.length !=0 && <Typography variant="h5">{t("Free managers")}:</Typography>}
                <br/>
                <ManagersList managers={managers}/>
          </>
          }
          <CreateManagerForm/>
      </>
    )
}