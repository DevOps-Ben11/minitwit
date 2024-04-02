import { useAuth } from '@/lib/hooks/useAuth'
import axios from 'axios'
import { Link, useNavigate } from 'react-router-dom'

type Props = {
  children: React.ReactNode
}

const Layout = ({ children }: Props) => {
  const { isAuthenticated, username } = useAuth()
  const navigate = useNavigate()

  const handleLogout = async () => {
    try {
      await axios.post('/api/logout')

      navigate('/public', {
        state: { flashMessage: 'You were logged out' },
      })
    } catch (error) {
      console.error(error)
    }
  }

  return (
    <div className='page'>
      <h1>MiniTwit</h1>

      <div className='navigation'>
        {isAuthenticated ? (
          <>
            <Link to='/'>my timeline</Link> |
            <Link to='/public'>public timeline</Link> |
            <button onClick={handleLogout}>sign out [{username}]</button>
          </>
        ) : (
          <>
            <Link to='/public'>public timeline</Link> |
            <Link to='/register'>sign up</Link> |
            <Link to='/login'>sign in</Link>
          </>
        )}
      </div>

      {children}

      <div className='footer'>MiniTwit &mdash; A Go - React Application</div>
    </div>
  )
}

export default Layout
