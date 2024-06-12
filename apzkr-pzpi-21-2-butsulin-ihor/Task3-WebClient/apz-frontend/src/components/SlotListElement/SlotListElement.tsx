import {Button, ListItem, ListItemText} from "@mui/material";
import {useNavigate} from "react-router-dom";
import AuthContext from "../../utils/auth.ts";
import {useContext} from "react";
import {useTranslation} from "react-i18next";

export default function SlotListElement({ slot }: { slot: Slot }) {
    const jwt = useContext(AuthContext);
    const { t } = useTranslation();

    async function deleteSlot(id: bigint){
        const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/admin/slot', {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
                "Authorization": "Bearer " + jwt
            },
            body: JSON.stringify({"slotID": id})
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

    const navigate = useNavigate();

    return (
      <ListItem
          key={slot.ID}
          secondaryAction={
            <>
              <Button variant="contained" color={"primary"} onClick={() => navigate("/slot/" + slot.ID)}>
                  {t("Go")}
              </Button>
              <Button variant="outlined" disabled={slot.item != null && slot.item.ID != null} color={"error"} onClick={() => deleteSlot(slot.ID)}>
                  {t("Delete")}
              </Button>
            </>
          }
          disablePadding
      >
          <ListItemText
              id={"slotListElement-label-" + slot.ID}
              primary={t("Slot") + " " + slot.ID.toString()}
              secondary={
                slot.item
                    ? t("Item") + ": " + slot.item.name
                    : t("Empty")
              }
          />
      </ListItem>
  );
}