import { useAuth } from '@/lib/hooks/useAuth'
import axios from 'axios'
import { Link } from 'react-router-dom'
import { toast } from 'react-toastify'

type Props = {
  children: React.ReactNode
}

const Layout = ({ children }: Props) => {
  const { isAuthenticated, username } = useAuth()

  const handleLogout = async () => {
    try {
      await axios.post('/api/logout')
      toast.success('You were logged out')
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
            <Link data-testid='my-timeline-link' to='/'>
              my timeline
            </Link>{' '}
            |
            <Link data-testid='public-timeline-link' to='/public'>
              public timeline
            </Link>{' '}
            |
            <button data-testid='sign-out' onClick={handleLogout}>
              sign out [{username}]
            </button>
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
