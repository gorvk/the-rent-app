import { useLocation } from "react-router-dom";

const ProductDetails = () => {
  const { search } = useLocation();
  const a = new URLSearchParams(search);
  const i = a.get("i");
  return <h1>Product Details for {i}</h1>;
};

export default ProductDetails;
