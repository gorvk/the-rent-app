import { useEffect, useState } from "react";
import { getCurrentUserOrdersApi } from "../../svc/order";
import { IOrderList } from "../../interfaces/models";
import { Card, CardContent, CardHeader, Typography } from "@mui/material";

const Order = () => {
  const [orders, setOrders] = useState<IOrderList[]>([]);
  useEffect(() => {
    const getCurrentUserOrders = async () => {
      try {
        const response = await getCurrentUserOrdersApi();
        if (response.isSuccess) {
          setOrders(response.result || []);
        }
      } catch (error) {
        console.error(error);
      }
    };
    getCurrentUserOrders();
  }, []);

  return orders.map((order) => {
    return (
      <Card style={{
        margin: "2em"
      }}>
        <CardHeader title={order.productName} />
        <CardContent>
          <Typography variant="body2" color="#18b300" fontWeight="bold">
            {order.price}
          </Typography>
          <Typography
            variant="body2"
            overflow={"hidden"}
            whiteSpace={"nowrap"}
            textOverflow={"ellipsis"}
            width="90%"
          >
            {order.orderStatus}
          </Typography>
        </CardContent>
      </Card>
    );
  });
};

export default Order;
