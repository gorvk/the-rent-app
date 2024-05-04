import { useEffect } from "react";
import { IGetAllProductsOutput } from "../interfaces/outputs";
import { getAllProductsApi } from "../svc/product";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../state/store";
import productsSlice from "../state/slices/productsSlice";
import { ProductsList } from "../common/product/productsList";

const Home = (props: { isRedirectedFromSearch?: boolean }) => {
  const productsList = useSelector((state: RootState) => state.productsList);
  const dipatch = useDispatch();

  useEffect(() => {
    if (!props.isRedirectedFromSearch) {
      getAllProducts();
    }
  }, []);

  const getAllProducts = async (): Promise<void> => {
    try {
      const products: IGetAllProductsOutput = await getAllProductsApi();
      dipatch(productsSlice.actions.setProductsList(products.result));
    } catch (error) {
      dipatch(productsSlice.actions.setProductsList([]));
      console.error(error);
    }
  };

  return (
    <div style={{padding:"40px"}}>
      <ProductsList productsList={productsList}/>
    </div>
  );
};

export default Home;
