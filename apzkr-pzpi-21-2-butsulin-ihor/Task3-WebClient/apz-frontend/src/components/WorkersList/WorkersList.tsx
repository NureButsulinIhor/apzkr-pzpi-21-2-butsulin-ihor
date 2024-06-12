import {Divider, List} from "@mui/material";
import WorkerListElement from "../WorkerListElement/WorkerListElement.tsx";

export default function WorkersList({ workers }: { workers: UserWorker[] }) {
    return (
        <List dense sx={{ width: '100%', bgcolor: 'background.paper' }}>
            {workers && workers.map((value) => (
                <>
                    <WorkerListElement key={"workerElement" + value.ID} worker={value} />
                    <br/>
                    <Divider key={"WorkerListDivider" + value.ID} variant="inset" component="li" />
                </>
            ))}
        </List>
    );
}