import ProductCard from "./productCard";
import { IProductCard } from "../../interfaces/models";
export const ProductsList = (props: { productsList: IProductCard[] }) => {
  const productsList = props.productsList || [];
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4">
      {productsList.map((product, index) => (
        <ProductCard key={index} product={product} />
      ))}
    </div>
  );
};
