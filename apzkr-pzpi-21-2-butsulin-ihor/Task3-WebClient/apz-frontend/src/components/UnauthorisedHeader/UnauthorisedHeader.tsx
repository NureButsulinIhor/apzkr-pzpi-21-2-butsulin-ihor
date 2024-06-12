import {AppBar, Toolbar, Typography} from "@mui/material";
import HeaderItem from "../HeaderItem/HeaderItem.tsx";
import {useTranslation} from "react-i18next";

export default function UnauthorisedHeader(){
    const { t, i18n } = useTranslation();

    return (
        <AppBar position="static">
            <Toolbar variant="dense">
                <HeaderItem variant={"h5"} href={"/"} name={t("Warehouse")}/>
                <Typography style={{marginRight: "1rem"}} variant={"h6"} color="inherit" onClick={() => i18n.changeLanguage("ua")}>
                    {"UA"}
                </Typography>
                <Typography style={{marginRight: "1rem"}} variant={"h6"} color="inherit" onClick={() => i18n.changeLanguage("en")}>
                    {"EN"}
                </Typography>
            </Toolbar>
        </AppBar>
    );
}