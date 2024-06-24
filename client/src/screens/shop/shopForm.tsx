import { useDispatch, useSelector } from "react-redux";
import { IUser } from "../../interfaces/models";
import { RootState } from "../../state/store";
import { ThemeProvider } from "@emotion/react";

import {
  Container,
  CssBaseline,
  Box,
  Typography,
  TextField,
  Button,
  Avatar,
} from "@mui/material";
import Store from "@mui/icons-material/Store";
import { ICreateShopInput } from "../../interfaces/inputs";
import { createShopApi } from "../../svc/shop";
import userSlice from "../../state/slices/userSlice";


export const ShopForm = () => {
  const currentUser: IUser | null = useSelector((state: RootState) => {
    return state.currentUser;
  });

  const dispatch = useDispatch();

  const handleCreateShop = async (event: React.FormEvent<HTMLFormElement>) => {
    try {
      event.preventDefault();
      const data = new FormData(event.currentTarget);
      const payload: ICreateShopInput = {
        shopName: data.get("shopName")?.toString() || "",
        email: data.get("email")?.toString() || "",
        phoneNumber: data.get("phoneNumber")?.toString() || "",
        mapLocation: data.get("mapLocation")?.toString() || "",
        shopType: data.get("shopType")?.toString() || "",
        shopDescription: data.get("shopDescription")?.toString() || "",
        city: data.get("city")?.toString() || "",
        country: data.get("country")?.toString() || "",
      };
      const response = await createShopApi(payload);
      if (response.isSuccess && currentUser) {
        const user = { ...currentUser };
        user.isShopEnabled = true;
        const setCurrentUser = userSlice.actions.setCurrentUser;
        dispatch(setCurrentUser(user));
      }
    } catch (error) {
      console.error(error);
    }
  };

  return (
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Avatar sx={{ m: 1, bgcolor: "secondary.main" }}>
            <Store />
          </Avatar>
          <Typography component="h1" variant="h5">
            Create and Open your own shop
          </Typography>
          <Box
            component="form"
            onSubmit={handleCreateShop}
            noValidate
            sx={{ mt: 1 }}
          >
            <TextField
              margin="normal"
              required
              fullWidth
              name="shopName"
              label="Shop Name"
              type="text"
              id="shop-name"
            />
            <TextField
              margin="normal"
              required
              fullWidth
              name="email"
              label="Shop Email"
              type="email"
              id="email"
            />
            <TextField
              margin="normal"
              required
              fullWidth
              name="phoneNumber"
              label="Phone Number"
              type="text"
              id="phone-number"
              autoComplete="current-password"
            />
            <TextField
              margin="normal"
              required
              fullWidth
              name="mapLocation"
              label="Map Location"
              type="text"
              id="map-location"
            />
            <TextField
              margin="normal"
              required
              fullWidth
              name="shopType"
              label="Shop Type"
              type="text"
              id="shop-type"
            />
            <TextField
              margin="normal"
              required
              fullWidth
              name="shopDescription"
              label="Shop Description"
              type="text"
              id="shop-description"
            />
            <TextField
              margin="normal"
              required
              fullWidth
              name="city"
              label="City"
              type="text"
              id="city"
            />
            <TextField
              margin="normal"
              required
              fullWidth
              name="country"
              label="Country"
              type="text"
              id="country"
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              create shop
            </Button>
          </Box>
        </Box>
      </Container>
  );
};
