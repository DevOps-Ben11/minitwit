import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './pages/App.tsx'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Login from './pages/Login.tsx'
import Register from './pages/Register.tsx'
import Layout from './pages/Layout.tsx'
import { Public } from './pages/Public.tsx'
import './style.css'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <BrowserRouter>
      <Layout>
        <Routes>
          <Route path='/login' element={<Login />} />
          <Route path='/register' element={<Register />} />
          <Route path='/public' element={<Public />} />
          <Route path='*' element={<App />} />
        </Routes>
      </Layout>
    </BrowserRouter>
  </React.StrictMode>
)
