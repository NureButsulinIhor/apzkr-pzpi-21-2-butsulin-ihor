import {useContext, useEffect, useState} from "react";
import AuthContext from "../utils/auth.ts";
import {useParams} from "react-router-dom";
import SlotsList from "../components/SlotsList/SlotsList.tsx";
import CreateSlotForm from "../components/CreateSlotForm/CreateSlotForm.tsx";
import {Typography} from "@mui/material";
import WorkersList from "../components/WorkersList/WorkersList.tsx";
import CreateWorkerForm from "../components/CreateWorkerForm/CreateWorkerForm.tsx";
import {useTranslation} from "react-i18next";

export default function Warehouse() {
    const { t } = useTranslation();
    const userToken = useContext(AuthContext);
    const {warehouseID} = useParams();
    const [warehouse, setWarehouse] = useState({} as Warehouse);

    async function fetchWarehouse() {
        const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/admin/warehouse/' + warehouseID, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${userToken}`
            },
        });
        if (!response.ok) {
            throw new Error('Failed!');
        }
        const data = await response.json() as { status: boolean, body: Warehouse };
        if (!data.status) {
            throw new Error('Failed to login');
        }

        setWarehouse(() => data.body)
    }

    useEffect(() => {
        fetchWarehouse();
    }, []);

    return (
      <>
          <h1>{t("Warehouse")} {warehouseID}</h1>
          {warehouse.storage && <>
              <Typography variant="h6">{t("Manager")}: {warehouse.manager.user.email}</Typography>
              <br/>
              <SlotsList slots={warehouse.storage.slots}/>
              <CreateSlotForm storageID={warehouse.storageID}/>
          </>}
          {warehouse.workers && <>
                {warehouse.workers.length !=0 && <Typography variant="h5">{t("Workers")}:</Typography>}
                <br/>
                <WorkersList workers={warehouse.workers}/>
          </>
          }
          <CreateWorkerForm warehouseID={warehouse.ID}/>
      </>
    )
}