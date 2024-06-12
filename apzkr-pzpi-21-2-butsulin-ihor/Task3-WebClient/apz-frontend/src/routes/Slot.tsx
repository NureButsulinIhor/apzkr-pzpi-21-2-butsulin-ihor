import {useContext, useEffect, useState} from "react";
import AuthContext from "../utils/auth.ts";
import {useParams} from "react-router-dom";
import {InputAdornment, TextField, Typography} from "@mui/material";
import UpdateSlotForm from "../components/UpdateSlotForm/UpdateSlotForm.tsx";
import UpdateItemForm from "../components/UpdateItemForm/UpdateItemForm.tsx";
import DeviceData from "../components/DeviceData/DeviceData.tsx";
import CreateItemForm from "../components/CreateItemForm/CreateItemForm.tsx";
import {useTranslation} from "react-i18next";

export default function Slot() {
    const { t } = useTranslation();
    const userToken = useContext(AuthContext);
    const {slotID} = useParams();
    const [slot, setSlot] = useState(null as Slot|null);

    async function fetchSlot() {
        const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/admin/slot/' + slotID, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${userToken}`
            },
        });
        if (!response.ok) {
            throw new Error('Failed!');
        }
        const data = await response.json() as { status: boolean, body: Slot };
        if (!data.status) {
            throw new Error('Failed to login');
        }

        setSlot(() => data.body)
    }

    useEffect(() => {
        fetchSlot();
    }, []);

    return (
      <>
          <h1>{t("Slot")} {slotID}</h1>
          {slot && <>
              <TextField
                  type="number"
                  disabled
                  label={t("Max weight")}
                  variant="outlined"
                  value={slot.maxWeight}
                  InputProps={{
                      endAdornment: <InputAdornment position="end">{t("kg")}</InputAdornment>,
                  }}
              />
              <TextField
                  type="number"
                  disabled
                  label={t("Last weight")}
                  variant="outlined"
                  value={slot.weighingResults.reduceRight((acc, curr) => acc.ID > curr.ID ? acc : curr, 0).weight}
                  InputProps={{
                      endAdornment: <InputAdornment position="end">kg</InputAdornment>,
                  }}
              />
              <br/>
              {slot.item && slot.item.ID ?
                  <>
                      <Typography variant="h5">{t("Item")}: {slot.item.name}</Typography>
                      <Typography variant="h6">{t("Description")}: {slot.item.description}</Typography>
                      <Typography variant="h6">{t("Weight")}: {slot.item.weight}</Typography>
                      <UpdateItemForm item={slot.item}/>
                  </>
                  : <>
                      <Typography variant="h5">{t("Empty")}</Typography>
                      <CreateItemForm slotID={slot.ID}/>
                  </>
              }
              <br/>
              <DeviceData device={slot.device} slotID={slot.ID}/>
              <br/>
              <UpdateSlotForm slot={slot}/>
          </>}
      </>
    )
}