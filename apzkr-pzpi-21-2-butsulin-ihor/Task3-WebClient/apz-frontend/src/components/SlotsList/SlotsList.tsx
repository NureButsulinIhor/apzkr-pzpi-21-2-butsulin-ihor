import {Divider, List} from "@mui/material";
import SlotListElement from "../SlotListElement/SlotListElement.tsx";

export default function SlotsList({ slots }: { slots: Slot[] }) {
    return (
        <List dense sx={{ width: '100%', bgcolor: 'background.paper' }}>
            {slots && slots.map((value) => (
                <>
                    <SlotListElement key={"slotElement" + value.ID} slot={value} />
                    <br/>
                    <Divider key={"SlotListDivider" + value.ID} variant="inset" component="li" />
                </>
            ))}
        </List>
    );
}