import React from 'react'
import ReactDOM from 'react-dom/client'
// import App from './pages/App.tsx'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Login from './pages/Login.tsx'
import Register from './pages/Register.tsx'
import Layout from './pages/Layout.tsx'
import { Public } from './pages/Public.tsx'
import { MyTimeline } from './pages/MyTimeline.tsx'
import { useAuth } from './lib/hooks/useAuth.tsx'
import { Timeline } from './pages/Timeline.tsx'
import './style.css'

const App = () => {
  const { isAuthenticated } = useAuth()

  return isAuthenticated ? <AuthenticatedApp /> : <UnAuthenticatedApp />
}

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <BrowserRouter>
      <Layout>
        <App />
      </Layout>
    </BrowserRouter>
  </React.StrictMode>
)

const AuthenticatedApp = () => {
  return (
    <Routes>
      <Route path='/public' element={<Public />} />
      <Route path='/' element={<MyTimeline />} />
      <Route path='/:username' element={<Timeline />} />
      <Route path='*' element={<MyTimeline />} />
    </Routes>
  )
}

const UnAuthenticatedApp = () => {
  return (
    <Routes>
      <Route path='/login' element={<Login />} />
      <Route path='/register' element={<Register />} />
      <Route path='/' element={<Public />} />
      <Route path='/public' element={<Public />} />
      <Route path='/:username' element={<Timeline />} />
      <Route path='*' element={<Public />} />
    </Routes>
  )
}
