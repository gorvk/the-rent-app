import { IProductCard } from "../../interfaces/models";
import { useNavigate } from "react-router-dom";

const ProductCard = (props: { product: IProductCard }) => {
  const { product } = props;
  const subHeader = `${product.productType} - ${product.shopName}`;
  const navigate = useNavigate();

  const navigateToProductDetails = () => {
    const url = `/product?i=${product.productId}`;
    navigate(url);
  };

  return (
    <div
      onClick={navigateToProductDetails}
      className="bg-white shadow-[0px_0px_0px_1px_#8c8c8c33] m-1 cursor-pointer hover:shadow-[0px_0px_0px_1px_#8c8c8c80]"
    >
      <img
        className="w-64 mx-auto"
        src="https://m.media-amazon.com/images/I/71msFUl565L._SL1500_.jpg"
      />
      <hr />
      <div className="bg-white px-3 py-5">
        <p className="truncate text-2xl">{product.productName}</p>
        <p className="truncate text-gray-500">{subHeader}</p>
        <br />
        <p className="truncate">{product.city}</p>
        <p className="truncate">{product.productDescription}</p>
        <p className="text-green-500 font-bold">{product.price}</p>
      </div>
    </div>
  );
};

export default ProductCard;
