import { useEffect, useState } from "react";
import { ProductsList } from "../product/productsList";
import { ISearchProductsInput } from "../../interfaces/inputs";
import { IProductCard } from "../../interfaces/models";
import { getSearchedProductsApi } from "../../svc/product";
import { useLocation } from "react-router-dom";

export const SearchResult = () => {
  const { search } = useLocation();
  const params = new URLSearchParams(search);
  const searchTerm = params.get("q") || "";
  const [productsList, setProductsList] = useState<IProductCard[]>([]);

  useEffect(() => {
    const searchProduct = async () => {
      const payload: ISearchProductsInput = { searchTerm };
      const response = await getSearchedProductsApi(payload);
      if (response.result.length > 0) {
        setProductsList(response.result);
      }
    };
    searchProduct();
  }, [searchTerm]);

  return (
    <div className="py-8 px-8">
      <p className="m-1 text-2xl font-bold">Search result for {searchTerm}</p>
      <ProductsList productsList={productsList} />
    </div>
  );
};
