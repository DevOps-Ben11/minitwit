import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './pages/App.tsx'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Login from './pages/Login.tsx'
import Register from './pages/Register.tsx'
import Layout from './pages/Layout.tsx'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <BrowserRouter>
      <Layout user={null}>
        <Routes>
          {' '}
          {/* This is like a switch only rendering one component (Route) at a time. */}
          <Route path='/login' element={<Login />} />
          <Route path='/register' element={<Register />} />
          <Route path='*' element={<App />} />{' '}
          {/* This mean every path should lead to main page if not created. */}
        </Routes>
      </Layout>
    </BrowserRouter>
  </React.StrictMode>
)
