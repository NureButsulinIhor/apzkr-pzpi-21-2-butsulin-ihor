import {useContext, useEffect, useState} from "react";
import AuthContext from "../utils/auth.ts";
import {useParams} from "react-router-dom";
import SlotsList from "../components/SlotsList/SlotsList.tsx";
import CreateSlotForm from "../components/CreateSlotForm/CreateSlotForm.tsx";
import {Typography} from "@mui/material";
import {useTranslation} from "react-i18next";

export default function Car() {
    const { t } = useTranslation();
    const userToken = useContext(AuthContext);
    const {carID} = useParams();
    const [car, setCar] = useState({} as Car);

    async function fetchCar() {
        const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/admin/car/' + carID, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${userToken}`
            },
        });
        if (!response.ok) {
            throw new Error('Failed!');
        }
        const data = await response.json() as { status: boolean, body: Car };
        if (!data.status) {
            throw new Error('Failed to login');
        }

        setCar(() => data.body)
    }

    useEffect(() => {
        fetchCar();
    }, []);

    return (
      <>
          <h1>{t("Car")} {carID}</h1>
          {car.storage && <>
              <Typography variant="h6">{t("Manager")}: {car.owner.email}</Typography>
              <br/>
              <SlotsList slots={car.storage.slots}/>
              <CreateSlotForm storageID={car.storageID}/>
          </>}
      </>
    )
}