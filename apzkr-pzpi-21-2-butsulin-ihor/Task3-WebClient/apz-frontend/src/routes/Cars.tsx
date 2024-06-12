import {useContext, useEffect, useState} from "react";
import AuthContext from "../utils/auth.ts";
import CarsList from "../components/CarsList/CarsList.tsx";
import CarForm from "../components/CarForm/CarForm.tsx";

export default function Cars(){
  const [cars, setCars] = useState([] as Car[]);
  const jwt = useContext(AuthContext);

  async function fetchCars(){
    const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/admin/car/all', {
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
    const data = await response.json() as { status: boolean, body: Car[] };
    if (!data.status) {
      throw new Error('Failed to login');
    }

    setCars(() => data.body);
  }

  useEffect(() => {
    fetchCars();
  }, []);

  return (
    <>
      <CarsList cars={cars}/>
      <br/>
      <CarForm/>
    </>
  )
}