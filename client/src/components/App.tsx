import { useAuth } from '@/lib/hooks/useAuth'
import { MyTimeline } from '@/pages/timelines/MyTimeline'
import { Public } from '@/pages/timelines/PublicTimeline'
import { Timeline } from '@/pages/timelines/UserTimeline'
import { Login } from '@/pages/Login'
import { Register } from '@/pages/Register'
import { Routes, Route, Navigate } from 'react-router-dom'

export const App = () => {
  const { isAuthenticated } = useAuth()

  return (
    <Routes>
      <Route path='/public' element={<Public />} />
      <Route path='/timeline/:username' element={<Timeline />} />

      {isAuthenticated ? (
        <>
          <Route path='/' element={<MyTimeline />} />
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
