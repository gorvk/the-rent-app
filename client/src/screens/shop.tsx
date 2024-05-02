import { useSelector } from "react-redux";
import { RootState } from "../state/store";

const Shop = () => {
  const currentUser = useSelector((state: RootState) => {
    console.log(state);
    return state.currentUser;
  });
  return (
    <h1>
      Welcome to your shop {currentUser?.firstname} {currentUser?.lastName}
    </h1>
  );
};

export default Shop;
