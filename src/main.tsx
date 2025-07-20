import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import {
  BrowserRouter,
  Routes,
  Route
} from "react-router"
import Layout from "@/Layout.tsx";

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route element={<Layout />}>
          <Route index element={<>HOME</>} />
        </Route>
      </Routes>
    </BrowserRouter>
  </StrictMode>,
)
