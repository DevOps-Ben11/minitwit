import { useAuth } from '@/lib/hooks/useAuth'
import { MyTimeline } from '@/pages/timelines/MyTimeline'
import { PublicTimeline } from '@/pages/timelines/PublicTimeline'
import { UserTimeline } from '@/pages/timelines/UserTimeline'
import { Login } from '@/pages/Login'
import { Logout } from '@/pages/Logout'
import { Register } from '@/pages/Register'
import { Routes, Route, Navigate } from 'react-router-dom'

export const App = () => {
  const { isAuthenticated } = useAuth()

  return (
    <Routes>
      {/* Everyone can access these routes */}
      <Route path='/public' element={<PublicTimeline />} />
      <Route path='/timeline/:username' element={<UserTimeline />} />

      {isAuthenticated ? (
        <>
          <Route path='/' element={<MyTimeline />} />
          <Route path='/logout' element={<Logout />} />
          <Route path='*' element={<Navigate to='/' />} />
        </>
      ) : (
        <>
          <Route path='/login' element={<Login />} />
          <Route path='/register' element={<Register />} />
          <Route path='*' element={<Navigate to='/public' />} />
        </>
      )}
    </Routes>
  )
}
