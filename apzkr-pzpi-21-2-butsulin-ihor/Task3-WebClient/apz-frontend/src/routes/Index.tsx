import {useContext, useEffect, useState} from "react";
import AuthContext from "../utils/auth.ts";
import WarehousesList from "../components/WarehouseList/WarehousesList.tsx";
import WarehouseForm from "../components/WarehouseForm/WarehouseForm.tsx";

export default function Index(){
  const [warehouses, setWarehouses] = useState([] as Warehouse[]);
  const jwt = useContext(AuthContext);

  async function fetchWarehouses(){
    const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/admin/warehouse/all', {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + jwt
      },
    });
    if (!response.ok) {
      throw new Error('Failed!');
    }
    console.log(response.body)
    const data = await response.json() as { status: boolean, body: Warehouse[] };
    if (!data.status) {
      throw new Error('Failed to login');
    }

    setWarehouses(() => data.body);
  }

  useEffect(() => {
    fetchWarehouses();
  }, []);

  return (
    <>
      <WarehousesList warehouses={warehouses}/>
      <br/>
      <WarehouseForm/>
    </>
  )
}