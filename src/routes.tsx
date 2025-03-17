import { Route, Routes } from "react-router";

export default function MainRoutes() {
  return (
    <Routes>
      <Route element={<>Layout</>}>
        <Route index element={<>Home</>} />
      </Route>
    </Routes>
  );
}
