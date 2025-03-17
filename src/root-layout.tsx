import { Outlet } from "react-router";

export default function Rootlayout() {
  return (
    <>
      <Outlet />
    </>
  );
}

function Desktop() {
  return <>Desktop</>;
}

function Tablet() {
  return <>Tablet</>;
}

function Mobile() {
  return <>Mobile</>;
}
