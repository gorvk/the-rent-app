import { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import { createOrderApi } from "../../svc/order";
import { ICreateProductInput, IPlaceOrderInput } from "../../interfaces/inputs";
import { ThemeProvider } from "@emotion/react";
import {
  Dialog,
  DialogContent,
  Container,
  CssBaseline,
  Box,
  Avatar,
  Typography,
  TextField,
  Button,
  createTheme,
} from "@mui/material";
import { createProductApi } from "../../svc/product";
import LocalShippingIcon from "@mui/icons-material/LocalShipping";
import { preProcessFile } from "typescript";
import { RootState } from "../../state/store";
import { useSelector } from "react-redux";

const OrderConfirmation = () => {
  const { search } = useLocation();
  const [productId, setProductId] = useState<number>(0);

  useEffect(() => {
    const searchParams = new URLSearchParams(search);
    const productId = searchParams.get("i");
    if (productId) {
      setProductId(parseInt(productId));
    }
  }, []);

  return <Form productId={productId} />;
};

const Form = (props: { productId: number }) => {
  const currentUser = useSelector((state: RootState) => state.currentUser);
  const defaultTheme = createTheme();
  const { productId } = props;
  const placeOrder = async (event: React.FormEvent<HTMLFormElement>) => {
    try {
      event.preventDefault();
      const data = new FormData(event.currentTarget);
      const payload: IPlaceOrderInput = {
        productId: productId,
        quantity: parseInt(data.get("quantity")?.toString() || "1"),
        toMapLocation:
          data.get("address")?.toString() || currentUser?.userAddress || "",
      };
      const response = await createOrderApi(payload);
      if (response.isSuccess) {
        // TODO: redirect to payment gateway
        console.log(response);
      }
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <ThemeProvider theme={defaultTheme}>
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            marginBlock: 3,
          }}
        >
          <Avatar sx={{ m: 1, bgcolor: "secondary.main" }}>
            <LocalShippingIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            Fill order details
          </Typography>
          <Box component="form" onSubmit={placeOrder} noValidate sx={{ mt: 1 }}>
            <TextField
              margin="normal"
              required
              fullWidth
              name="address"
              label="Delivery Address"
              type="text"
              id="delivery-address"
            />
            <TextField
              margin="normal"
              required
              fullWidth
              name="quantity"
              label="Quantity"
              type="number"
              id="quantity"
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              Go to payment
            </Button>
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
  );
};

export default OrderConfirmation;
