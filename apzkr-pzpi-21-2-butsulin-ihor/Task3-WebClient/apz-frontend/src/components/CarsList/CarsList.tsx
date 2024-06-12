import {Divider, List} from "@mui/material";
import CarListElement from "../CarListElement/CarListElement.tsx";

export default function CarsList({ cars }: { cars: Car[] }) {
    return (
        <List dense sx={{ width: '100%', bgcolor: 'background.paper' }}>
            {cars && cars.map((value) => (
                <>
                    <CarListElement key={"carElement" + value.ID} car={value} />
                    <br/>
                    <Divider key={"CarListDivider" + value.ID} variant="inset" component="li" />
                </>
            ))}
        </List>
    );
}