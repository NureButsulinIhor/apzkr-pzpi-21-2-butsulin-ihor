import {Button, ListItem, ListItemText} from "@mui/material";
import {useNavigate} from "react-router-dom";
import AuthContext from "../../utils/auth.ts";
import {useContext} from "react";
import {useTranslation} from "react-i18next";

export default function CarListElement({ car }: { car: Car }) {
    const { t } = useTranslation();
    const jwt = useContext(AuthContext);

    async function deleteCar(id: bigint){
        const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/admin/car', {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
                "Authorization": "Bearer " + jwt
            },
            body: JSON.stringify({"carID": id})
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
          key={car.ID}
          secondaryAction={
            <>
              <Button variant="contained" color={"primary"} onClick={() => navigate("/car/" + car.ID)}>
                  {t("Go")}
              </Button>
              <Button variant="outlined" color={"error"} onClick={() => deleteCar(car.ID)}>
                  {t("Delete")}
              </Button>
            </>
          }
          disablePadding
      >
          <ListItemText
              id={"carListElement-label-" + car.ID}
              primary={t("Car") + " " + car.ID.toString()}
          />
      </ListItem>
  );
}