import {Divider, List} from "@mui/material";
import WarehouseListElement from "../WarehouseListElement/WarehouseListElement.tsx";

export default function WarehousesList({ warehouses }: { warehouses: Warehouse[] }) {
    return (
        <List dense sx={{ width: '100%', bgcolor: 'background.paper' }}>
            {warehouses && warehouses.map((value) => (
                <>
                    <WarehouseListElement key={"warehouseElement" + value.ID} warehouse={value} />
                    <br/>
                    <Divider key={"WarehouseListDivider" + value.ID} variant="inset" component="li" />
                </>
            ))}
        </List>
    );
}