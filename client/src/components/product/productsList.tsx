import { Grid } from "@mui/material";
import ProductCard from "./productCard";
import { IProductCard } from "../../interfaces/models";
export const ProductsList = (props: { productsList: IProductCard[] }) => {
  const productsList = props.productsList || [];
  return (
    <Grid item>
      <Grid
        container
        justifyContent={{ xs: "center", md: "center", lg: "left" }}
        rowSpacing={9}
        columnSpacing={9}
      >
        {productsList.map((product, index) => (
          <Grid key={index} item>
            <ProductCard product={product} />
          </Grid>
        ))}
      </Grid>
    </Grid>
  );
};
