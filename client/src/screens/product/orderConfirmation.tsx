import { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import { createOrderApi } from "../../svc/order";
import { IPlaceOrderInput } from "../../interfaces/inputs";
import { ThemeProvider } from "@emotion/react";
import {
  Container,
  CssBaseline,
  Box,
  TextField,
  Button,
  createTheme,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
} from "@mui/material";
import { RootState } from "../../state/store";
import { useSelector } from "react-redux";

const OrderConfirmation = (props: { productQuantity: number }) => {
  const { productQuantity } = props;
  const { search } = useLocation();
  const [productId, setProductId] = useState<number>(0);
  const [quantities, setQuantities] = useState<number[]>([]);

  useEffect(() => {
    const searchParams = new URLSearchParams(search);
    const productId = searchParams.get("i");
    if (productId) {
      const filledQuantities = fillQuanities(productQuantity);
      console.log(productQuantity);
      setQuantities(filledQuantities);
      setProductId(parseInt(productId));
    }
    // eslint-disable-next-line
  }, []);

  return <Form productId={productId} quantities={quantities} />;
};

const fillQuanities = (quantity: number) => {
  const quantities: number[] = [];
  for (let i = 0; i < quantity; i++) {
    quantities.push(i + 1);
  }
  return quantities;
};

const Form = (props: { productId: number; quantities: number[] }) => {
  const currentUser = useSelector((state: RootState) => state.currentUser);
  const defaultTheme = createTheme();
  const { productId, quantities } = props;
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
      <Container
        component="main"
        maxWidth="xs"
        style={{ paddingLeft: 0, marginLeft: 0 }}
      >
        <CssBaseline />
        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
            alignItems: "left",
          }}
        >
          <Box component="form" onSubmit={placeOrder} noValidate sx={{ mt: 1 }}>
            <TextField
              disabled={quantities.length === 0}
              margin="normal"
              required
              fullWidth
              name="address"
              label="Delivery Address"
              type="text"
              id="delivery-address"
            />
            <FormControl fullWidth>
              <InputLabel id="demo-simple-select-label">Quantity</InputLabel>
              <Select
                disabled={quantities.length === 0}
                labelId="demo-simple-select-label"
                id="demo-simple-select"
                label="Quantity"
                name="quantity"
              >
                {quantities.map((quantity) => (
                  <MenuItem key={quantity} value={quantity}>
                    {quantity}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>
            <Button
              disabled={quantities.length === 0}
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              {quantities.length === 0 ? "Out of stock" : "Go to payment"}
            </Button>
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
  );
};

export default OrderConfirmation;
