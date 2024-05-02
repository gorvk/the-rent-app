import { Grid } from "@mui/material";
import { useState, useEffect } from "react";
import Product from "../common/product/product";
import { IProduct } from "../interfaces/models";
import { IGetAllProductsOutput } from "../interfaces/outputs";
import { getAllProductsApi } from "../svc/product";

const AllProducts = () => {
  const [products, setProducts] = useState<IProduct[]>([{} as IProduct]);
  useEffect(() => {
    getAllProducts();
  }, []);

  const getAllProducts = async (): Promise<void> => {
    try {
      const products: IGetAllProductsOutput = await getAllProductsApi();
      setProducts(products.result);
    } catch (error) {
      setProducts([] as IProduct[]);
      console.error(error);
    }
  };
  
  return (
    <Grid item xs={12}>
      <Grid
        container
        justifyContent={{ xs: "center", md: "center", lg: "left" }}
        padding={5}
        rowSpacing={8}
        columnSpacing={10}
      >
        {products.map((product, index) => (
          <Grid key={index} item>
            <Product product={product} key={index} />
          </Grid>
        ))}
      </Grid>
    </Grid>
  );
};

export default AllProducts;
