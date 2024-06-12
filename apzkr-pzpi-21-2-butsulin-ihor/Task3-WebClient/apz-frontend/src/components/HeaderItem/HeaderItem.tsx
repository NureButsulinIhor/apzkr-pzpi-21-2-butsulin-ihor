import {Link} from "react-router-dom";
import {Typography, TypographyPropsVariantOverrides} from "@mui/material";
import {OverridableStringUnion} from "@mui/types";
import {Variant} from "@mui/material/styles/createTypography";

export default function HeaderItem({variant, href, name}: {variant: OverridableStringUnion<Variant | "inherit", TypographyPropsVariantOverrides> | undefined, href: string, name: string}) {
    return (
        <Typography style={{marginRight: "1rem"}} variant={variant} color="inherit" component={({children, ...params}: {children: React.ReactNode}) => (
            <Link to={href} {...params}>
                {children}
            </Link>
        )}>
            {name}
        </Typography>
    );
}