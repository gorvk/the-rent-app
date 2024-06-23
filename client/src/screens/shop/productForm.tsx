import { IProductCard } from "../../interfaces/models";
import { ThemeProvider } from "@emotion/react";
import {
  Container,
  CssBaseline,
  Box,
  Avatar,
  Typography,
  TextField,
  Button,
  createTheme,
  Card,
  CardContent,
  Grid,
  Dialog,
  DialogContent,
} from "@mui/material";
import AddBusinessIcon from "@mui/icons-material/AddBusiness";
import React, { useEffect, useState } from "react";
import AddCircleIcon from "@mui/icons-material/AddCircle";
import { ICreateProductInput } from "../../interfaces/inputs";
import { GetCurrentShopProductApi, createProductApi } from "../../svc/product";
import ProductCard from "../product/productCard";
const defaultTheme = createTheme();

export const ProductForm = () => {
  const [productsList, setProductsList] = useState<IProductCard[]>([]);

  const GetCurrentShopProduct = async () => {
    try {
      const response = await GetCurrentShopProductApi();
      if (response.result) {
        setProductsList(response.result);
        handleClose();
      }
    } catch (error) {
      console.error(error);
    }
  };

  useEffect(() => {
    GetCurrentShopProduct();
    // eslint-disable-next-line
  }, []);

  const cardCssProp: React.CSSProperties = {
    width: 300,
    height: 380,
    backgroundColor: "white",
    cursor: "pointer",
    border:
      "rgba(9, 30, 66, 0.25) 0px 1px 1px, rgba(9, 30, 66, 0.13) 0px 0px 1px 1px",
  };

  const [open, setOpen] = React.useState(false);

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <>
      <Grid item>
        <Grid
          container
          justifyContent={{ xs: "center", md: "center", lg: "left" }}
          rowSpacing={9}
          columnSpacing={9}
        >
          <Grid onClick={handleClickOpen} item>
            <Card style={cardCssProp}>
              <CardContent
                style={{
                  display: "flex",
                  justifyContent: "center",
                  alignItems: "center",
                  height: "85%",
                }}
              >
                <AddCircleIcon />
                <div>Add new Product</div>
              </CardContent>
            </Card>
          </Grid>
          {productsList.map((product, index) => (
            <Grid key={index} item>
              <ProductCard product={product} />
            </Grid>
          ))}
        </Grid>
      </Grid>
      <Form
        open={open}
        handleClose={handleClose}
        refreshProductsList={GetCurrentShopProduct}
      />
    </>
  );
};

const Form = (props: {
  open: boolean;
  handleClose: () => void;
  refreshProductsList: () => Promise<void>;
}) => {
  const { open, handleClose, refreshProductsList } = props;

  const handleCreateProduct = async (
    event: React.FormEvent<HTMLFormElement>
  ) => {
    try {
      event.preventDefault();
      const data = new FormData(event.currentTarget);
      const payload: ICreateProductInput = {
        productName: data.get("productName")?.toString() || "",
        productType: data.get("productType")?.toString() || "",
        productCondition: data.get("productCondition")?.toString() || "",
        price: data.get("price")?.toString() || "",
        originalPurchasedDate:
          data.get("originalPurchasedDate")?.toString() || "",
        originalPurchaisingRecieptNo:
          data.get("originalPurchaisingRecieptNo")?.toString() || "",
        productDescription: data.get("productDescription")?.toString() || "",
        quantity: parseInt(data.get("quantity")?.toString() || "1")
      };
      const response = await createProductApi(payload);
      if (response.isSuccess) {
        refreshProductsList();
      }
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <Dialog fullWidth open={open} onClose={handleClose}>
      <DialogContent>
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
                <AddBusinessIcon />
              </Avatar>
              <Typography component="h1" variant="h5">
                Add product to your shop
              </Typography>
              <Box
                component="form"
                onSubmit={handleCreateProduct}
                noValidate
                sx={{ mt: 1 }}
              >
                <TextField
                  margin="normal"
                  required
                  fullWidth
                  name="productName"
                  label="Product Name"
                  type="text"
                  id="product-name"
                />
                <TextField
                  margin="normal"
                  required
                  fullWidth
                  name="productType"
                  label="Product Type"
                  type="text"
                  id="product-type"
                />
                <TextField
                  margin="normal"
                  required
                  fullWidth
                  name="productCondition"
                  label="Product Condition"
                  type="text"
                  id="product-condition"
                />
                <TextField
                  margin="normal"
                  required
                  fullWidth
                  name="price"
                  label="Price"
                  type="text"
                  id="price"
                />
                <TextField
                  margin="normal"
                  required
                  fullWidth
                  InputLabelProps={{ shrink: true }}
                  name="originalPurchasedDate"
                  label="Original Purchased Date"
                  type="date"
                  id="original-purchased-date"
                />
                <TextField
                  margin="normal"
                  required
                  fullWidth
                  name="originalPurchaisingRecieptNo"
                  label="Original Purchaising Reciept No"
                  type="text"
                  id="original-purchaising-reciept-no"
                />
                <TextField
                  margin="normal"
                  required
                  fullWidth
                  name="productDescription"
                  label="Product Description"
                  type="text"
                  id="product-description"
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
                  create shop
                </Button>
              </Box>
            </Box>
          </Container>
        </ThemeProvider>
      </DialogContent>
    </Dialog>
  );
};
