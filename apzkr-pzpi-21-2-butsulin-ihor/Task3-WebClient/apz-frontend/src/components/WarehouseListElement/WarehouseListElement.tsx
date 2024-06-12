import {Button, ListItem, ListItemText} from "@mui/material";
import {useNavigate} from "react-router-dom";
import AuthContext from "../../utils/auth.ts";
import {useContext} from "react";
import {useTranslation} from "react-i18next";

export default function WarehouseListElement({ warehouse }: { warehouse: Warehouse }) {
    const { t } = useTranslation();
    const jwt = useContext(AuthContext);

    async function deleteWarehouse(id: bigint){
        const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/admin/warehouse', {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
                "Authorization": "Bearer " + jwt
            },
            body: JSON.stringify({"warehouseID": id})
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
          key={warehouse.ID}
          secondaryAction={
            <>
              <Button variant="contained" color={"primary"} onClick={() => navigate("/warehouse/" + warehouse.ID)}>
                  {t("Go")}
              </Button>
              <Button variant="outlined" color={"error"} onClick={() => deleteWarehouse(warehouse.ID)}>
                  {t("Delete")}
              </Button>
            </>
          }
          disablePadding
      >
          <ListItemText
              id={"warehouseListElement-label-" + warehouse.ID}
              primary={t("Warehouse")+" " + warehouse.ID.toString()}
          />
      </ListItem>
  );
}