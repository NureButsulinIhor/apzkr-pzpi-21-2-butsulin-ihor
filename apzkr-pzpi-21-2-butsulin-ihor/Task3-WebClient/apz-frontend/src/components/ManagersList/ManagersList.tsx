import {Divider, List} from "@mui/material";
import ManagerListElement from "../ManagerListElement/ManagerListElement.tsx";

export default function ManagersList({ managers }: { managers: User[] | null }) {
    return (
        <List dense sx={{ width: '100%', bgcolor: 'background.paper' }}>
            {managers != null && managers.map((value) => (
                <>
                    <ManagerListElement key={"managerElement" + value.ID} manager={value} />
                    <br/>
                    <Divider key={"ManagerListDivider" + value.ID} variant="inset" component="li" />
                </>
            ))}
        </List>
    );
}