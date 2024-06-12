import {Button, ListItem, ListItemText} from "@mui/material";
import {useNavigate} from "react-router-dom";
import AuthContext from "../../utils/auth.ts";
import {useContext} from "react";
import {useTranslation} from "react-i18next";

export default function WorkerListElement({ worker }: { worker: UserWorker }) {
    const { t } = useTranslation();
    const jwt = useContext(AuthContext);

    async function deleteWorker(id: bigint){
        const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/register/worker', {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
                "Authorization": "Bearer " + jwt
            },
            body: JSON.stringify({"workerID": id})
        });
        if (!response.ok) {
            throw new Error('Failed!');
        }
        console.log(response.body)
        const data = await response.json() as { status: boolean };
        if (!data.status) {
            throw new Error('Failed');
        }

        navigate("/")
    }

    const navigate = useNavigate();

    return (
      <ListItem
          key={"worker+"+worker.ID}
          secondaryAction={
            <>
              <Button variant="outlined" color={"error"} onClick={() => deleteWorker(worker.ID)}>
                  {t("Delete")}
              </Button>
            </>
          }
          disablePadding
      >
          <ListItemText
              id={"workerListElement-label-" + worker.ID}
              primary={worker.user.email}
          />
      </ListItem>
  );
}