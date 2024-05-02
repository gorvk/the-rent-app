import Card from "@mui/material/Card";
import CardHeader from "@mui/material/CardHeader";
import CardMedia from "@mui/material/CardMedia";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import { IProduct } from "../../interfaces/models";

const Product = (props: { product: IProduct }) => {
  const { product } = props;
  return (
    <Card sx={{ width: 300, height: 350, backgroundColor: "white" }}>
      <CardMedia
        component="img"
        height="194"
        width="194"
        image="https://5.imimg.com/data5/SELLER/Default/2021/2/LZ/TT/JX/122336/primo-plastic-air-cooler-1000x1000.jpg"
        alt="Paella dish"
      />
      <CardHeader title={product.productName} subheader={product.productType} />
      <CardContent>
        <Typography variant="body2" height={20} overflow={"hidden"} whiteSpace={"nowrap"} textOverflow={"ellipsis"}>
          {product.productDescription}
        </Typography>
        <Typography variant="body2" color="#18b300" fontWeight="bold">
          {product.price}
        </Typography>
      </CardContent>
    </Card>
  );
};

export default Product;
