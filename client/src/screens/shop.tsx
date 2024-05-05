import { useSelector } from "react-redux";
import { RootState } from "../state/store";
import { IUser } from "../interfaces/models";
import { ShopForm } from "../components/shop/shopForm";
import { ProductForm } from "../components/shop/productForm";

const Shop = () => {
  const currentUser: IUser | null = useSelector((state: RootState) => {
    return state.currentUser;
  });

  if (currentUser == null) {
    return <h1>please login again</h1>;
  }

  return currentUser.isShopEnabled ? <ProductForm /> : <ShopForm />;
};

export default Shop;