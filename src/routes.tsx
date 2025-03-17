import { Route, Routes } from "react-router";
import Rootlayout from "./root-layout";

export default function MainRoutes() {
  return (
    <Routes>
      <Route element={<Rootlayout />}>
        <Route index element={<>Home</>} />
      </Route>
    </Routes>
  );
}
