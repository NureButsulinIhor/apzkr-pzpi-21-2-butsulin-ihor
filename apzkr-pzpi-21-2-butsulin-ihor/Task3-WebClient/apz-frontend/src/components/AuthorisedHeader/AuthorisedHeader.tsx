import {AppBar, Toolbar, Typography} from "@mui/material";
import HeaderItem from "../HeaderItem/HeaderItem.tsx";
import {useTranslation} from "react-i18next";

export default function AuthorisedHeader(){
    const { t, i18n } = useTranslation();
    return (
        <AppBar position="static">
            <Toolbar variant="dense">
                <HeaderItem variant={"h5"} href={"/"} name={t("Warehouse")}/>

                <HeaderItem variant={"h6"} href={"/cars"} name={t("Cars")}/>
                <HeaderItem variant={"h6"} href={"/managers"} name={t("Managers")}/>
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