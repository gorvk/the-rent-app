import MenuIcon from "@mui/icons-material/Menu";
import StoreIcon from "@mui/icons-material/Store";
import LocalShippingIcon from "@mui/icons-material/LocalShipping";
import AccountCircleIcon from "@mui/icons-material/AccountCircle";

import {
  Box,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Drawer,
  IconButton,
} from "@mui/material";
import { Fragment, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Home, Login } from "@mui/icons-material";
import { useSelector } from "react-redux";
import { RootState } from "../state/store";

interface DrawerItem {
  label: string;
  icon: JSX.Element;
  link: string;
}

const DrawerItems = (props: { items: DrawerItem[] }) => {
  const navigate = useNavigate();
  const items = props.items;
  return (
    <>
      {items.map((item, index) => (
        <ListItem key={index} disablePadding>
          <ListItemButton onClick={() => navigate(item.link)}>
            <ListItemIcon>{item.icon}</ListItemIcon>
            <ListItemText primary={item.label} />
          </ListItemButton>
        </ListItem>
      ))}
    </>
  );
};

const SideDrawer = () => {
  const currentUser = useSelector((state: RootState) => state.currentUser);
  const [open, setOpen] = useState(false);

  const drawerItems: DrawerItem[] = [
    { label: "Home", icon: <Home />, link: "/" },
  ];
  const protectedDrawerItems: DrawerItem[] = [
    { label: "My Shop", icon: <StoreIcon />, link: "/shop" },
    { label: "Track Orders", icon: <LocalShippingIcon />, link: "/orders" },
  ];

  const protectedAuthDrawerItems: DrawerItem[] = [
    { label: "Login", icon: <Login />, link: "/login" },
    { label: "Register", icon: <AccountCircleIcon />, link: "/register" },
  ];

  const toggleDrawer = (newOpen: boolean) => () => {
    setOpen(newOpen);
  };

  return (
    <Fragment>
      <IconButton
        size="large"
        edge="start"
        color="inherit"
        aria-label="open drawer"
        sx={{ mr: 2 }}
        onClick={toggleDrawer(true)}
      >
        <MenuIcon />
      </IconButton>
      <Drawer open={open} onClose={toggleDrawer(false)}>
        <Box
          sx={{ width: 250 }}
          role="presentation"
          onClick={toggleDrawer(false)}
        >
          <List>
            <DrawerItems items={drawerItems} />
            {currentUser === null || currentUser.id === undefined ? (
              <DrawerItems items={protectedAuthDrawerItems} />
            ) : (
              <DrawerItems items={protectedDrawerItems} />
            )}
          </List>
        </Box>
      </Drawer>
    </Fragment>
  );
};

export default SideDrawer;
