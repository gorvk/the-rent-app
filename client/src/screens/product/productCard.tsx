import Card from "@mui/material/Card";
import CardHeader from "@mui/material/CardHeader";
import CardMedia from "@mui/material/CardMedia";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import { IProductCard } from "../../interfaces/models";
import Divider from "@mui/material/Divider";
import { useNavigate } from "react-router-dom";

const ProductCard = (props: { product: IProductCard }) => {
  const { product } = props;
  const subHeader = `${product.productType} - ${product.shopName}`;
  const navigate = useNavigate();

  const cardCssProp: React.CSSProperties = {
    width: 300,
    height: 380,
    backgroundColor: "white",
    cursor: "pointer",
    boxShadow: "0px 0px 0px 1px rgb(140 140 140/.2)",
  };

  const navigateToProductDetails = () => {
    const url = `/product?i=${product.productId}`;
    navigate(url);
  };

  const ProductCardTypography = (props: { children: string }) => {
    return (
      <Typography
        variant="body2"
        overflow={"hidden"}
        whiteSpace={"nowrap"}
        textOverflow={"ellipsis"}
        width="90%"
      >
        {props.children}
      </Typography>
    );
  };

  return (
    <Card onClick={navigateToProductDetails} style={cardCssProp}>
      <CardMedia
        component="img"
        height="194"
        width="194"
        image="https://m.media-amazon.com/images/I/71msFUl565L._SL1500_.jpg"
        alt="Paella dish"
      />
      <Divider />
      <CardHeader
        title={product.productName}
        subheader={<ProductCardTypography>{subHeader}</ProductCardTypography>}
      />
      <CardContent>
        <ProductCardTypography>{product.city}</ProductCardTypography>
        <ProductCardTypography>
          {product.productDescription}
        </ProductCardTypography>
        <Typography variant="body2" color="#18b300" fontWeight="bold">
          {product.price}
        </Typography>
      </CardContent>
    </Card>
  );
};

export default ProductCard;
