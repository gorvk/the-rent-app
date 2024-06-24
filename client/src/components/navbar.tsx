import { MouseEvent, useState } from "react";
import IconButton from "@mui/material/IconButton";
import { Link } from "react-router-dom";
import { useSelector } from "react-redux";
import { RootState } from "../state/store";
import { SearchBar } from "./searchBar";
import SideDrawer from "./drawerList";
import { AccountMenuIcon } from "./icons";
import AccountMenu from "./accountMenu";

export function Navbar() {
  const currentUser = useSelector((state: RootState) => state.currentUser);
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);

  const handleProfileMenuOpen = (event: MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const menuId = "primary-search-account-menu";

  return (
    <>
      <div className="static bg-purple-500 shadow-[0px_0px_0px_1px_#8c8c8c33] text-black flex items-center justify-between py-1 px-3">
        <div className="w-[10%] border-red-500 border-solid border-2 flex items-center justify-between">
          <SideDrawer />
          <Link className="font-bold text-lg" to={"/"}>
            THE RENT APP
          </Link>
        </div>
        <div className="w-[23%] border-green-500 border-solid border-2 flex items-center justify-between">
          {currentUser?.firstname && `Hi, ${currentUser?.firstname}`}
          <SearchBar />
          {currentUser?.id !== undefined && (
            <IconButton
              size="large"
              edge="end"
              aria-label="account of current user"
              aria-controls={menuId}
              aria-haspopup="true"
              onClick={handleProfileMenuOpen}
              color="inherit"
            >
              <AccountMenuIcon />
            </IconButton>
          )}
        </div>
      </div>
      <AccountMenu
        anchorEl={anchorEl}
        menuId={menuId}
        setAnchorEl={setAnchorEl}
      ></AccountMenu>
    </>
  );
}
