import { useEffect, useState } from "react";
import { IGetAllProductsOutput } from "../../interfaces/outputs";
import { getAllProductsApi } from "../../svc/product";
import { ProductsList } from "../product/productsList";
import { IProductCard } from "../../interfaces/models";

const Home = () => {
  const [productsList, setProductsList] = useState<IProductCard[]>([]);

  useEffect(() => {
    const getAllProducts = async (): Promise<void> => {
      try {
        const products: IGetAllProductsOutput = await getAllProductsApi();
        setProductsList(products.result);
      } catch (error) {
        setProductsList([]);
        console.error(error);
      }
    };
    getAllProducts();
  }, []);

  return (
    <div style={{ padding: "40px" }}>
      <ProductsList productsList={productsList} />
    </div>
  );
};

export default Home;
