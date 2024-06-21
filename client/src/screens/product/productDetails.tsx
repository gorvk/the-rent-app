import { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { getProduct } from "../../svc/product";
import { IGetProductInput } from "../../interfaces/inputs";
import { IProduct } from "../../interfaces/models";
import { Divider } from "@mui/material";
import OrderConfirmation from "./orderConfirmation";

const ProductDetails = () => {
  const { search } = useLocation();
  const [product, setProduct] = useState<IProduct>({} as IProduct);

  const pageContainerStyleProps: React.CSSProperties = {
    marginTop: "5vh",
    display: "flex",
    justifyContent: "space-evenly",
  };

  useEffect(() => {
    const getProductDetail = async () => {
      const searchParams = new URLSearchParams(search);
      const productId = searchParams.get("i");
      if (productId) {
        const payload: IGetProductInput = { id: parseInt(productId) };
        const response = await getProduct(payload);
        setProduct(response.result);
      }
    };
    getProductDetail();
    // eslint-disable-next-line
  }, []);
  return product.productId ? (
    <div style={pageContainerStyleProps}>
      <ProductThumbnail src="https://m.media-amazon.com/images/I/71msFUl565L._SL1500_.jpg" />
      <DetailsSections product={product} />
    </div>
  ) : (
    <></>
  );
};

const DetailsSections = (props: { product: IProduct }) => {
  const { product } = props;
  const containerCSSProperties: React.CSSProperties = { width: "60%" };

  return (
    <div style={containerCSSProperties}>
      <Fields product={product} />
    </div>
  );
};

const Fields = (props: { product: IProduct }) => {
  const { product } = props;
  return (
    <div>
      <ProductHeader
        productName={product.productName || ""}
        shopName={product.shopName || ""}
      />
      <Divider />
      <ProductDetailFields
        style={{ color: "#18b300", fontSize: "1.8em" }}
        label="Monthly Rent:"
        content={product.price || ""}
      />
      <ProductDetailFields
        label="About Product:"
        content={product.productDescription || ""}
      />
      <ProductDetailFields label="Type:" content={product.productType || ""} />
      <ProductDetailFields
        label="Condition:"
        content={product.productCondition || ""}
      />
      <ProductDetailFields
        label="Seller from:"
        content={(product.city, product.country) || ""}
      />
      <ProductDetailFields
        label="Original Purchased Date:"
        content={new Date(product.originalPurchasedDate || "").toDateString()}
      />
      <ProductDetailFields
        label="Original Purchaising Reciept No:"
        content={product.originalPurchaisingRecieptNo || ""}
      />
      <ProductDetailFields
        label="Available Quantity:"
        content={product.quantity?.toString() || ""}
      />
      <ProductDetailFields
        label="Place your order:"
        content={<OrderConfirmation productQuantity={product.quantity} />}
      />
    </div>
  );
};

const ProductHeader = (props: { productName: string; shopName: string }) => {
  const { productName, shopName } = props;
  return (
    <>
      <div style={{ fontSize: "1.8em" }}>{productName}</div>
      <div style={{ display: "block" }}>
        <div>by {shopName}</div>
      </div>
      <br />
    </>
  );
};

const ProductThumbnail = (props: { src: string }) => {
  const { src } = props;
  return (
    <div style={{ width: "30%" }}>
      <img width="100%" alt="product thumbnail" src={src} />
    </div>
  );
};

const ProductDetailFields = (props: {
  label: string;
  content: JSX.Element | string;
  style?: React.CSSProperties;
}) => {
  const { label, content, style } = props;
  return (
    <>
      <br />
      <div>
        <b>{label}</b> <br />
        <div style={style}>{content}</div>
      </div>
    </>
  );
};

export default ProductDetails;
