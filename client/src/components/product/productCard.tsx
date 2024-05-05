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
    border: "rgba(9, 30, 66, 0.25) 0px 1px 1px, rgba(9, 30, 66, 0.13) 0px 0px 1px 1px"
  }

  return (
    <Card
      onClick={() => {
        navigate("/product");
      }}
      style={cardCssProp}
    >
      <CardMedia
        component="img"
        height="194"
        width="194"
        image="https://5.imimg.com/data5/SELLER/Default/2021/2/LZ/TT/JX/122336/primo-plastic-air-cooler-1000x1000.jpg"
        alt="Paella dish"
      />
      <Divider />
      <CardHeader
        title={product.productName}
        subheader={
          <Typography
            variant="body2"
            overflow={"hidden"}
            whiteSpace={"nowrap"}
            textOverflow={"ellipsis"}
            width="90%"
          >
            {subHeader}
          </Typography>
        }
      />
      <CardContent>
        <Typography
          variant="body2"
          overflow={"hidden"}
          whiteSpace={"nowrap"}
          textOverflow={"ellipsis"}
        >
          {product.city}, {product.country}
        </Typography>
        <Typography
          variant="body2"
          overflow={"hidden"}
          whiteSpace={"nowrap"}
          textOverflow={"ellipsis"}
        >
          {product.productDescription}
        </Typography>
        <Typography variant="body2" color="#18b300" fontWeight="bold">
          {product.price}
        </Typography>
      </CardContent>
    </Card>
  );
};

export default ProductCard;
