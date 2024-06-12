import {ListItem, ListItemText} from "@mui/material";

export default function ManagerListElement({ manager }: { manager: User }) {
    return (
      <ListItem
          key={"manager+"+manager.ID}
          disablePadding
      >
          <ListItemText
              id={"managerListElement-label-" + manager.ID}
              primary={manager.email}
          />
      </ListItem>
  );
}