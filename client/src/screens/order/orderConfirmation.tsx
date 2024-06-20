import { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";

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
  return <h1>Order Confirmation {productId}</h1>;
};

export default OrderConfirmation;
